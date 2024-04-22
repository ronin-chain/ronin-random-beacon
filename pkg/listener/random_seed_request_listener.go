package listener

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ronin-chain/ronin-random-beacon/pkg/contract"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/chaininfo"
	"github.com/ronin-chain/ronin-random-beacon/pkg/event"
	"github.com/ronin-chain/ronin-random-beacon/pkg/utils"
)

type RandomRequestListener struct {
	ChainInfoId   string
	ChainListener *event.EventLooper

	contract       *contract.RandomBeaconCoordinator
	database       *db.Client
	startingBlock  uint64
	blockSize      uint64
	processedBlock uint64
}

// Implements EventProcessor interface.
func (l *RandomRequestListener) Started(ctx context.Context) {
	log.Info("[RandomRequestListener][Started]")

	defer l.syncUp()
	timer := time.NewTicker(time.Second)
	timeout := time.NewTimer(time.Second * 60)
	for {
		select {
		case <-timer.C:
			if l.getCurrentBlock() != 0 {
				return
			}
			if l.getFinalizedBlock() != 0 {
				return
			}
		case <-ctx.Done():
			log.Debug("[RandomRequestListener][Started] call done")
			timer.Stop()
			timeout.Stop()
			return
		case <-timeout.C:
			panic(fmt.Errorf("could not fetch latest block number"))
		}
	}
}

func (l *RandomRequestListener) Stopped(context.Context) {
	log.Info("[RandomRequestListener][Stopped]")
}

func (l *RandomRequestListener) Process(context.Context) {
	l.syncBatchLogs()
}

// Returns the event looper of the random request listener.
func NewRandomRequestListener(
	chainInfoId string,
	blockListener *event.EventLooper,
	contract *contract.RandomBeaconCoordinator,
	db *db.Client,
	startingBlock uint64,
	blockSize uint64,
	interval uint64) event.EventLooper {
	processor := RandomRequestListener{
		ChainInfoId:    chainInfoId,
		ChainListener:  blockListener,
		contract:       contract,
		database:       db,
		startingBlock:  startingBlock,
		blockSize:      blockSize,
		processedBlock: 0,
	}
	return event.NewEventLooper(&processor, interval)
}

// Returns the current block.
func (l *RandomRequestListener) getCurrentBlock() uint64 {
	currentBlock := l.ChainListener.Processor.(*ChainListener).GetCurrentBlock()
	if currentBlock == 0 {
		log.Warn("[RandomRequestListener][getCurrentBlock] current block equals to 0")
	}
	return currentBlock
}

// Returns the finalized block.
func (l *RandomRequestListener) getFinalizedBlock() uint64 {
	finalizedBlock := l.ChainListener.Processor.(*ChainListener).GetFinalizedBlockNumber()
	return finalizedBlock
}

// Returns the processed block.
// If the tracker does not work before, it returns the starting block.
func (l *RandomRequestListener) getProcessedBlock(ctx context.Context) uint64 {
	if l.processedBlock == 0 {
		result, err := l.database.ChainInfo.Get(ctx, l.ChainInfoId)
		if err == nil {
			return result.ProcessedBlock
		}
		log.Error("[RandomRequestListener][getProcessedBlock] could not query processed block", "err", err)
		return l.startingBlock
	}
	return l.processedBlock
}

// Updates the processed block for both postgres db and current cache.
func (l *RandomRequestListener) setProcessedBlock(block uint64, ctx context.Context) error {
	err := l.database.ChainInfo.Create().
		SetID(l.ChainInfoId).
		SetProcessedBlock(block).
		OnConflict(sql.ConflictColumns(chaininfo.FieldID)).
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		return err
	}
	l.processedBlock = block
	return nil
}

func (l *RandomRequestListener) syncUp() {
	ctx := context.Background()
	finalizedBlock := l.getFinalizedBlock()
	if currentBlock := l.getCurrentBlock(); currentBlock == 0 {
		return
	}
	log.Info("[RandomRequestListener][SyncUp] Keep try sync Processed block to finalizedBlock before shutting down.")
	for {
		log.Debug("[RandomRequestListener][syncUp] start looping")
		processedBlock := l.getProcessedBlock(ctx)
		if processedBlock >= finalizedBlock {
			log.Info("[RandomRequestListener][syncUp] processed block is greater than finalized block stop Looping", "processedBlock", processedBlock, "finalizedBlock", finalizedBlock)
			return
		}
		l.syncBatchLogs()
	}

}
func (l *RandomRequestListener) getEventsFromBlockRange(from uint64, to uint64) ([]*contract.RandomBeaconCoordinatorRandomSeedRequested, error) {
	log.Debug("[RandomRequestListener][getEventsFromBlockRange]", "from", from, "to", to)
	iter, err := l.contract.FilterRandomSeedRequested(&bind.FilterOpts{
		Start: from,
		End:   &to,
	}, make([]*big.Int, 0), make([][32]byte, 0))

	if err != nil {
		return nil, err
	}
	var events []*contract.RandomBeaconCoordinatorRandomSeedRequested

	for iter.Next() {
		log.Info("[RandomRequestListener][getEventsFromBlockRange]", "event", iter.Event)
		events = append(events, iter.Event)
	}

	if err := iter.Error(); err != nil {
		return events, err
	}

	return events, nil
}

func (l *RandomRequestListener) storeEvents(events []*contract.RandomBeaconCoordinatorRandomSeedRequested, ctx context.Context) error {
	if len(events) == 0 {
		return nil
	}

	var bulk []*db.RandomRequestCreate

	for _, event := range events {
		raw, err := json.Marshal(event)
		if err != nil {
			return err
		}

		bulk = append(bulk, l.database.RandomRequest.Create().
			SetID(hex.EncodeToString(event.ReqHash[:])).
			SetBlockNumber(event.Raw.BlockNumber).
			SetLogIndex(event.Raw.Index).
			SetRaw(string(raw)))
	}

	_, err := l.database.RandomRequest.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (l *RandomRequestListener) syncBatchLogs() error {
	ctx := context.Background()
	currentBlock := l.getCurrentBlock()
	processedBlock := l.getProcessedBlock(ctx)
	finalizedBlock := l.getFinalizedBlock()

	log.Debug("[RandomRequestListener][syncBatchLogs] function triggered", "processedBlock", processedBlock, "currentBlock", currentBlock)

	log.Debug("[RandomRequestListener][syncBatchLogs] loop condition", "block", finalizedBlock, "currentBlock", currentBlock)

	var processChunk uint64 = 0

	for processedBlock < finalizedBlock && processChunk < 10 {
		processChunk++
		from := processedBlock + 1
		// Min (processedBlock + l.blockSize, finalizedBlock)
		to := utils.Min(processedBlock+l.blockSize, finalizedBlock)
		if from > to {
			break
		}
		log.Info("[RandomRequestListener][syncBatchLogs] processing logs", "from", from, "to", to)

		events, err := l.getEventsFromBlockRange(from, to)
		if err != nil {
			log.Error("[RandomRequestListener][getBatchLogs] error while pulling events", "events", len(events), "error", err)
			return err
		}
		log.Info("[RandomRequestListener][getBatchLogs] events pulled", "events", len(events), "fromBlock", from, "toBlock", to)

		err = l.storeEvents(events, ctx)
		if err != nil {
			log.Error("[RandomRequestListener][getBatchLogs] could not store batch events", "events", len(events), "error", err)
			return err
		}
		log.Info("[RandomRequestListener][getBatchLogs] events stored", "events", len(events), "fromBlock", from, "toBlock", to)

		processedBlock = to
		err = l.setProcessedBlock(processedBlock, ctx)
		if err != nil {
			log.Error("[RandomRequestListener][getBatchLogs] could not update processed block", "processedBlock", processedBlock, "error", err)
			return err
		}
		log.Info("[RandomRequestListener][getBatchLogs] updated processed block")

	}
	return nil
}
