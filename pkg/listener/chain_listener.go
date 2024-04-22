package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ronin-chain/ronin-random-beacon/pkg/contract"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/chaininfo"
	"github.com/ronin-chain/ronin-random-beacon/pkg/event"
)

type FinalizedStruct struct {
	Result struct {
		BlockNumber uint64 `json:"blockNumber"`
	} `json:"result"`
}

type ChainListener struct {
	ChainInfoId    string
	currentBlock   uint64
	finalizedBlock uint64
	dbClient       *db.Client
	rpcClient      *ethclient.Client
	skynetEndpoint string
	contract       *contract.RandomBeaconCoordinatorSession
}

// Implements EventProcessor interface.
func (l *ChainListener) Started(context.Context) {
	log.Info("[ChainListener][Started]")
	l.syncLatestBlock()
	l.syncFinalizedBlock()
}
func (l *ChainListener) Stopped(context.Context) {
	log.Info("[ChainListener][Stopped]")
}

func (l *ChainListener) Process(context.Context) {
	l.syncLatestBlock()
	l.syncFinalizedBlock()
}

func (l *ChainListener) getFinalizedBlockNumber(skynetEndpoint string, path string) (*FinalizedStruct, error) {
	client := &http.Client{Transport: &http.Transport{}, Timeout: 15 * time.Second}
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", skynetEndpoint, path), nil)
	if err != nil {
		return &FinalizedStruct{}, err
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return &FinalizedStruct{}, err
	}

	if response.StatusCode != 200 {
		return &FinalizedStruct{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &FinalizedStruct{}, err
	}

	var rs *FinalizedStruct
	if err := json.Unmarshal(body, &rs); err != nil {
		return &FinalizedStruct{}, err
	}

	if rs.Result.BlockNumber <= 0 {
		return &FinalizedStruct{}, err
	}

	return rs, nil
}

// Syncs latest block number.
func (l *ChainListener) syncLatestBlock() {
	log.Debug("[ChainListener][syncLatestBlock] fetching latest block", "currentBlock", l.currentBlock)
	block, err := l.rpcClient.BlockNumber(context.Background())
	if err != nil {
		log.Error("[ChainListener][syncLatestBlock] could not get latest block number", "err", err)
		return
	}
	l.currentBlock = block
}

// Syncs finalized block number.
func (l *ChainListener) syncFinalizedBlock() error {
	log.Debug("[ChainListener][syncLatestBlock] fetching finalized block", "currentBlock", l.currentBlock)
	finalizedStruc, err := l.getFinalizedBlockNumber(l.skynetEndpoint, "ronin/blocks/finalized")
	if err != nil {
		log.Error("[ChainListener][syncLatestBlock] could not get finalized block number", "err", err)
		return err
	}

	l.finalizedBlock = finalizedStruc.Result.BlockNumber

	err = l.dbClient.ChainInfo.Create().
		SetID(l.ChainInfoId).
		SetFinalizedBlock(finalizedStruc.Result.BlockNumber).
		OnConflict(sql.ConflictColumns(chaininfo.FieldID)).
		UpdateNewValues().
		Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}

// Exposes current block.
func (l *ChainListener) GetCurrentBlock() uint64 {
	return l.currentBlock
}

// Expose current  MaxRepsonseBlock
func (l *ChainListener) GetFinalizedBlockNumber() uint64 {
	if l.finalizedBlock == 0 {
		log.Warn("[RandomRequestListener][getFinalizedBlock] finalized block equals to 0")
		result, err := l.dbClient.ChainInfo.Get(context.Background(), l.ChainInfoId)
		if err == nil {
			return result.FinalizedBlock
		}
		log.Error("[RandomRequestListener][getFinalizedBlock] could not query finalized block", "err", err)

	}
	return l.finalizedBlock
}

// Returns the event looper of latest block listener.
func NewChainListener(chainInfoId string, dbClient *db.Client, rpcClient *ethclient.Client, contract *contract.RandomBeaconCoordinatorSession, skynetEndpoint string, interval uint64) (event.EventLooper, *ChainListener) {
	processor := ChainListener{chainInfoId, 0, 0, dbClient, rpcClient, skynetEndpoint, contract}
	event := event.NewEventLooper(&processor, interval)
	return event, &processor
}
