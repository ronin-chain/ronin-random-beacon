package listener

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ronin-chain/ronin-random-beacon/pkg/contract"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db"
	"github.com/ronin-chain/ronin-random-beacon/pkg/event"
	"github.com/ronin-chain/ronin-random-beacon/pkg/utils"
)

const (
	BufferBlock = 2
)

type GetFinalized struct {
	ID      int           `json:"id"`
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type FinalizedStruct struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BaseFeePerGas    string        `json:"baseFeePerGas"`
		Difficulty       string        `json:"difficulty"`
		ExtraData        string        `json:"extraData"`
		GasLimit         string        `json:"gasLimit"`
		GasUsed          string        `json:"gasUsed"`
		Hash             string        `json:"hash"`
		LogsBloom        string        `json:"logsBloom"`
		Miner            string        `json:"miner"`
		MixHash          string        `json:"mixHash"`
		Nonce            string        `json:"nonce"`
		Number           string        `json:"number"`
		ParentHash       string        `json:"parentHash"`
		ReceiptsRoot     string        `json:"receiptsRoot"`
		Sha3Uncles       string        `json:"sha3Uncles"`
		Size             string        `json:"size"`
		StateRoot        string        `json:"stateRoot"`
		Timestamp        string        `json:"timestamp"`
		TotalDifficulty  string        `json:"totalDifficulty"`
		Transactions     []string      `json:"transactions"`
		TransactionsRoot string        `json:"transactionsRoot"`
		Uncles           []interface{} `json:"uncles"`
	} `json:"result"`
}

type ChainListener struct {
	ChainInfoId    string
	currentBlock   uint64
	finalizedBlock uint64
	dbClient       *db.Client
	rpcClient      *ethclient.Client
	rpcEndpoint    string
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

func (l *ChainListener) getFinalizedBlockNumber(rpcEndpoint string) (uint64, error) {
	client := &http.Client{Transport: &http.Transport{}, Timeout: 15 * time.Second}
	data := GetFinalized{
		ID:      1,
		Jsonrpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{"finalized", false},
	}
	// Encode the struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest(http.MethodPost, rpcEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, err
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}

	if response.StatusCode != http.StatusOK {
		return 0, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var rs *FinalizedStruct
	if err := json.Unmarshal(body, &rs); err != nil {
		return 0, err
	}

	number, err := utils.ParseBlockNumberWithOptionalDelay(rs.Result.Number, BufferBlock)
	return number, err
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
	finalizedBlock, err := l.getFinalizedBlockNumber(l.rpcEndpoint)
	log.Info("[ChainListener][syncLatestBlock] fetching finalized block", "finalizedBlock", finalizedBlock, "currentFinalizedBlock", l.finalizedBlock)
	if err != nil {
		log.Error("[ChainListener][syncLatestBlock] could not get finalized block number", "err", err)
		return err
	}
	if finalizedBlock < l.finalizedBlock {
		log.Warn("[ChainListener][syncLatestBlock] finalized block from rpc is lower than current", "newfinalizedBlock", finalizedBlock, "currentFinalizedBlock", l.finalizedBlock)
		return nil
	}
	l.finalizedBlock = finalizedBlock
	log.Info("[ChainListener][syncLatestBlock] updating finalized block", "finalizedBlock", finalizedBlock, "currentFinalizedBlock", l.finalizedBlock)
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
		l.syncFinalizedBlock()
	}
	return l.finalizedBlock
}

// Returns the event looper of latest block listener.
func NewChainListener(chainInfoId string, dbClient *db.Client, rpcClient *ethclient.Client, contract *contract.RandomBeaconCoordinatorSession, rpcEndpoint string, interval uint64) (event.EventLooper, *ChainListener) {
	processor := ChainListener{chainInfoId, 0, 0, dbClient, rpcClient, rpcEndpoint, contract}
	event := event.NewEventLooper(&processor, interval)
	return event, &processor
}
