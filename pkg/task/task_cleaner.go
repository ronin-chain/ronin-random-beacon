package task

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/task"
	"github.com/ronin-chain/ronin-random-beacon/pkg/event"
)

type TaskCleaner struct {
	dbClient *db.Client
	interval uint64
}

// Implements EventProcessor interface.
func (w *TaskCleaner) Started(context.Context) {
	log.Info("[TaskCleaner][Started]")
}

func (w *TaskCleaner) Stopped(context.Context) {
	log.Info("[TaskCleaner][Stopped]")
}

func (w *TaskCleaner) Process(context.Context) {
	err := w.cleanupOldTasks()
	if err != nil {
		log.Error("[TaskCleaner][Process] could not cleanup Old tasks", "error", err)
	}
}

func NewTaskCleaner(dbClient *db.Client, interval uint64) event.EventLooper {
	processor := TaskCleaner{dbClient, interval}
	return event.NewEventLooper(&processor, interval)
}

func (w *TaskCleaner) cleanupOldTasks() error {
	ctx := context.Background()
	// Get the time 2 months ago
	twoMonthsAgo := int(time.Now().Unix()) - 60*60*24*60 // 2 months ago
	// Delete all tasks where lastSent is older than 2 months
	affected, err := w.dbClient.Task.
		Delete().
		Where(task.LastSentLT(twoMonthsAgo)).
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("error while deleting old tasks, error=%v", err)
	}
	log.Info("[TaskCleaner][cleanupOldTasks] Deleted old tasks", "affected", affected)
	return nil
}
