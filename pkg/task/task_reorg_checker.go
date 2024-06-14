package task

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ronin-chain/ronin-random-beacon/pkg/contract"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db"
	dbTask "github.com/ronin-chain/ronin-random-beacon/pkg/db/task"
	"github.com/ronin-chain/ronin-random-beacon/pkg/event"
	"github.com/ronin-chain/ronin-random-beacon/pkg/listener"
)

type TaskReorgChecker struct {
	dbClient      *db.Client
	ethClient     *ethclient.Client
	coordinator   *contract.RandomBeaconCoordinatorSession
	chainListener *listener.ChainListener
	batchSize     int
}

// Implements EventProcessor interface.
func (w *TaskReorgChecker) Started(context.Context) {
	log.Info("[TaskReorgChecker][Started]")
}

func (w *TaskReorgChecker) Stopped(context.Context) {
	log.Info("[TaskReorgChecker][Stopped]")
}

func (w *TaskReorgChecker) Process(context.Context) {
	err := w.reorgCheckerTasks()
	if err != nil {
		log.Error("[TaskReorgChecker][Process] could not send task", "error", err)
	}
}

func NewTaskReorgChecker(dbClient *db.Client, ethClient *ethclient.Client, coordinator *contract.RandomBeaconCoordinatorSession,
	chainListener *listener.ChainListener, batchSize int, interval uint64) event.EventLooper {
	processor := TaskReorgChecker{dbClient, ethClient, coordinator, chainListener, batchSize}
	return event.NewEventLooper(&processor, interval)
}

func (w *TaskReorgChecker) reorgCheckerTasks() error {
	// Check all success tasks that are not finalized and before the finalized Block.
	ctx := context.Background()
	finalizedBlock := w.chainListener.GetFinalizedBlockNumber()

	tasks, err := w.dbClient.Task.Query().Where(
		dbTask.StatusEQ(dbTask.StatusSuccess),
		dbTask.IsFinalizedEQ(false),
		dbTask.SentBlockLT(int(finalizedBlock)),
	).Order(db.Asc(dbTask.FieldLastSent)).Limit(w.batchSize).All(ctx)
	if err != nil {
		return fmt.Errorf("[TaskReorgChecker][reorgCheckerTasks] Error while querying task, error=%v", err)
	}
	log.Info("[TaskReorgChecker][reorgCheckerTasks] Checking tasks Success", "totalTasks", len(tasks), "finalizedBlock", finalizedBlock)
	reOrg := 0
	for _, task := range tasks {
		receipt, err := w.ethClient.TransactionReceipt(ctx, common.HexToHash(task.TxHash))
		if receipt.Status == types.ReceiptStatusSuccessful {
			task.Update().SetIsFinalized(true).Save(ctx)
			continue
		}

		if err == ethereum.NotFound || receipt.BlockNumber == nil {
			log.Error("[TaskReorgChecker][reorgCheckerTasks] Could not get receipt for task", "taskId", task.ID, "txHash", task.TxHash)
			task.Update().SetStatus(dbTask.StatusPending).SetReOrg(true).Save(ctx)
			reOrg++
			continue
		}
	}

	log.Info("[TaskReorgChecker][reorgCheckerTasks] Total Reorg", "totalReOrg", reOrg, "finalizedBlock", finalizedBlock)

	return nil
}
