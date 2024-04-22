package task

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ronin-chain/ronin-random-beacon/pkg/config"
	"github.com/ronin-chain/ronin-random-beacon/pkg/contract"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db"
	dbTask "github.com/ronin-chain/ronin-random-beacon/pkg/db/task"
	"github.com/ronin-chain/ronin-random-beacon/pkg/event"
	"github.com/ronin-chain/ronin-random-beacon/pkg/listener"
	"github.com/ronin-chain/ronin-random-beacon/pkg/utils"
)

type TaskSender struct {
	dbClient      *db.Client
	batchSize     int
	coordinator   *contract.RandomBeaconCoordinatorSession
	chainListener *listener.ChainListener
}

// Implements EventProcessor interface.
func (w *TaskSender) Started(context.Context) {
	log.Info("[TaskSender][Started]")
}

func (w *TaskSender) Stopped(context.Context) {
	log.Info("[TaskSender][Stopped]")
}

func (w *TaskSender) Process(context.Context) {
	err := w.sendTasks()
	if err != nil {
		log.Error("[TaskSender][Process] could not send task", "error", err)
	}
}

func NewTaskSender(dbClient *db.Client, batchSize int, contract *contract.RandomBeaconCoordinatorSession, chainListener *listener.ChainListener, interval uint64) event.EventLooper {
	processor := TaskSender{dbClient, batchSize, contract, chainListener}
	return event.NewEventLooper(&processor, interval)
}

// Sends tasks. // TODO: improve flow, send tx to contract ASAP when received the event
func (w *TaskSender) sendTasks() error {
	ctx := context.Background()
	timestamp := int(time.Now().Unix()) - 60*10 // 10 minute ago
	// Query full tasks with pending stats and last_sent in recent 10 minute ago.
	tasks, err := w.dbClient.Task.Query().Where(
		dbTask.And(dbTask.StatusEQ(dbTask.StatusPending)), dbTask.LastSentGT(timestamp)).
		Order(db.Asc(dbTask.FieldID)).
		Limit(w.batchSize).
		All(ctx)

	if err != nil {
		return fmt.Errorf("error while querying task, error=%v", err)
	}

	log.Info("[TaskSender][sendTasks] sending tasks", "tasks", len(tasks))
	for _, task := range tasks {
		log.Debug("[TaskSender][sendTasks] handling task", "taskId", task.ID)

		if task.Attempts > config.FailedMaxAttempts {
			log.Error("[TaskSender][sendTasks] mark task as failed", "taskId", task.ID)
			task.Update().SetStatus(dbTask.StatusFailed).Save(ctx)
			continue
		}

		var params FulfillRandomSeedParams
		if err := json.Unmarshal([]byte(task.Data), &params); err != nil {
			return fmt.Errorf("could not unmarshal task data, taskId=%v, error=%v", task.ID, err)
		}

		updatingTask := task.Update()

		w.coordinator.TransactOpts.GasPrice = utils.ConvertStringToBigInt("20000000000")
		tx, err := w.coordinator.FulfillRandomSeed(params.Req, params.Proof)
		if err != nil {
			log.Error("[TaskSender][sendTasks] got error on sending task", "taskId", task.ID, "err", err)
			updatingTask = updatingTask.SetLastError(err.Error())
		} else {
			updatingTask = updatingTask.
				SetTxHash(tx.Hash().String()).
				SetStatus(dbTask.StatusSent).
				SetSentBlock(int(w.chainListener.GetCurrentBlock())).
				SetLastSent(int(time.Now().Unix()))
		}
		if tx != nil {
			log.Info("[TaskSender][sendTasks] task sent", "taskId", task.ID, "hash", tx.Hash().String())
		}

		task, err = updatingTask.Save(ctx)
		if err != nil {
			log.Error("[TaskSender][sendTasks] got error on update task", "taskId", task.ID, "err", err)
		}
	}

	return nil
}
