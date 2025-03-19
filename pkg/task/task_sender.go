package task

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ronin-chain/ronin-random-beacon/pkg/config"
	"github.com/ronin-chain/ronin-random-beacon/pkg/contract"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db"
	dbTask "github.com/ronin-chain/ronin-random-beacon/pkg/db/task"
	"github.com/ronin-chain/ronin-random-beacon/pkg/event"
	"github.com/ronin-chain/ronin-random-beacon/pkg/listener"
)

type TaskSender struct {
	dbClient      *db.Client
	batchSize     int
	coordinator   *contract.RandomBeaconCoordinatorSession
	chainListener *listener.ChainListener
	KeyHash       common.Hash
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

func NewTaskSender(dbClient *db.Client, batchSize int, contract *contract.RandomBeaconCoordinatorSession, chainListener *listener.ChainListener, interval uint64, keyHash common.Hash) event.EventLooper {
	processor := TaskSender{dbClient, batchSize, contract, chainListener, keyHash}
	return event.NewEventLooper(&processor, interval)
}

// Sends tasks. // TODO: improve flow, send tx to contract ASAP when received the event
func (w *TaskSender) sendTasks() error {
	ctx := context.Background()

	tasks, err := w.dbClient.Task.Query().Where(
		func(s *sql.Selector) {
			s.Where(
				sql.And(
					sql.ExprP("last_sent + POW(2, attempts) < $1", time.Now().Unix()), // last_sent + 2^attempts < now
					sql.EQ(dbTask.FieldStatus, dbTask.StatusPending),
				),
			)
		}).Order(db.Asc(dbTask.FieldID)).
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
		// Get last Finalized period
		lastedPeriod, err := w.coordinator.GetLastFinalizedPeriod()
		if err != nil {
			return fmt.Errorf("could not get last finalized period, error=%v", err)
		}
		if params.Req.Period.Cmp(lastedPeriod) <= 0 {
			log.Info("[TaskSender][sendTasks] period is already finalized", "taskId", task.ID)
			task, err := task.Update().SetStatus(dbTask.StatusFinalized).Save(ctx)
			if err != nil {
				log.Error("[TaskSender][sendTasks] got error on update task", "taskId", task.ID, "err", err)
			}
			continue
		}
		// Check isSubmitted at By Consensus
		isSubmittedAt, err := w.coordinator.IsSubmittedAtByKeyHash(params.Req.Period, w.KeyHash)
		if err != nil {
			log.Error("[TaskSender][sendTasks] could not check IsSubmittedAtByKeyHash", "error", err, "period", params.Req.Period.String(), "keyHash", w.KeyHash.String())
			continue
		}

		if isSubmittedAt {
			log.Info("[TaskSender][sendTasks] period is already submitted", "taskId", task.ID)
			task, err := task.Update().SetStatus(dbTask.StatusFinalized).Save(ctx)
			if err != nil {
				log.Error("[TaskSender][sendTasks] got error on update task", "taskId", task.ID, "err", err)
			}
			continue
		}

		updatingTask := task.Update()

		tx, err := w.coordinator.FulfillRandomSeed(params.Req, params.Proof)
		if err != nil {
			log.Error("[TaskSender][sendTasks] got error on sending task", "taskId", task.ID, "err", err)
			updatingTask = updatingTask.SetLastError(err.Error()).
				SetAttempts(task.Attempts + 1)
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
