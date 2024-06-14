package task

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db"
	dbTask "github.com/ronin-chain/ronin-random-beacon/pkg/db/task"
	"github.com/ronin-chain/ronin-random-beacon/pkg/event"
	"github.com/ronin-chain/ronin-random-beacon/pkg/listener"
)

type TaskSyncerSent struct {
	database      *db.Client
	batchSize     int
	client        *ethclient.Client
	chainListener *listener.ChainListener
}

// Implements EventProcessor interface.
func (w *TaskSyncerSent) Started(context.Context) {
	log.Info("[TaskSyncerSent][Started]")
}

func (w *TaskSyncerSent) Stopped(context.Context) {
	log.Info("[TaskSyncerSent][Stopped]")
}

func (w *TaskSyncerSent) Process(context.Context) {
	err := w.syncStateTasks()
	if err != nil {
		log.Error("[TaskSyncerSent][Process] could not send task", "error", err)
	}
}

func NewTaskSyncerSent(database *db.Client, batchSize int, client *ethclient.Client, chainListener *listener.ChainListener, interval uint64) event.EventLooper {
	processor := TaskSyncerSent{database, batchSize, client, chainListener}
	return event.NewEventLooper(&processor, interval)
}

func (w *TaskSyncerSent) syncStateTasks() error {
	ctx := context.Background()

	timestamp := int(time.Now().Unix()) - 60 // 1 minute ago
	tasks, err := w.database.Task.Query().Where(
		dbTask.StatusEQ(dbTask.StatusSent),
		dbTask.LastSentLT(timestamp),
	).Order(db.Asc(dbTask.FieldLastSent)).Limit(w.batchSize).All(ctx)

	if err != nil {
		return fmt.Errorf("error while querying task, error=%v", err)
	}

	log.Info("[TaskSyncerSent][handleSyncerSent]", "timestamp", timestamp, "tasks", len(tasks))
	var pendingIds []int
	var successIds []int

	for _, task := range tasks {
		receipt, err := w.client.TransactionReceipt(ctx, common.HexToHash(task.TxHash))
		if err == ethereum.NotFound || receipt.BlockNumber == nil {
			log.Error("[TaskSyncerSent][handleSyncerSent] could not found tx receipt, retry", "taskId", task.ID, "tx", receipt)
			pendingIds = append(pendingIds, task.ID)
			continue
		}

		if receipt.Status == types.ReceiptStatusFailed {
			log.Error("[TaskSyncerSent][handleSyncerSent] transaction failed", "taskId", task.ID, "receipt", receipt)
			_, err := task.Update().
				SetLastError("Previous receipt status  is failed").
				SetAttempts(task.Attempts + 1).
				SetStatus(dbTask.StatusPending).
				Save(ctx)
			if err != nil {
				log.Error("[TaskSyncerSent][sendTasks] got error on update task", "taskId", task.ID, "err", err)
			}
		} else if receipt.Status == types.ReceiptStatusSuccessful {
			log.Info("[TaskSyncerSent][handleSyncerSent] task successful", "taskId", task.ID)
			successIds = append(successIds, task.ID)
		}
	}

	_, err = w.database.Task.Update().Where(dbTask.IDIn(pendingIds...)).SetStatus(dbTask.StatusPending).Save(ctx)
	if err != nil {
		log.Error("[TaskSyncerSent][handleSyncerSent] could not update task status", "pendingIds", pendingIds, "err", err)
	}

	_, err = w.database.Task.Update().Where(dbTask.IDIn(successIds...)).SetStatus(dbTask.StatusSuccess).Save(ctx)
	if err != nil {
		log.Error("[TaskSyncerSent][handleSyncerSent] could not update task status", "successIds", successIds, "err", err)
	}
	return nil
}
