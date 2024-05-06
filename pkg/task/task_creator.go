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
	"github.com/ronin-chain/ronin-random-beacon/pkg/core"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/predicate"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/randomrequest"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/task"
	"github.com/ronin-chain/ronin-random-beacon/pkg/event"
)

type TaskCreator struct {
	dbClient  *db.Client
	config    *config.AppConfig
	batchSize int
}

type FulfillRandomSeedParams struct {
	ReqHash common.Hash
	Proof   contract.VRFProof
	Req     contract.RandomRequest
}

// Implements EventProcessor interface.
func (w *TaskCreator) Started(context.Context) {
	log.Info("[TaskCreator][Started]")
}

func (w *TaskCreator) Stopped(context.Context) {
	log.Info("[TaskCreator][Stopped]")
}

func (w *TaskCreator) Process(context.Context) {
	err := w.createTasks()
	if err != nil {
		log.Error("[TaskCreator][Process] could not create task", "error", err)
	}
}

func NewTaskCreator(dbClient *db.Client, config *config.AppConfig, batchSize int, interval uint64) event.EventLooper {
	processor := TaskCreator{dbClient, config, batchSize}
	return event.NewEventLooper(&processor, interval)
}

// Creates tasks. // TODO: improve flow, send tx to contract ASAP when received the event
func (w *TaskCreator) createTasks() error {
	ctx := context.Background()
	// Query rows from RandomRequest where tasks are null
	events, err := w.dbClient.RandomRequest.Query().
		Where(predicate.RandomRequest(func(s *sql.Selector) {
			s.Where(sql.IsNull(randomrequest.TaskColumn))
		})).
		Order(db.Asc(randomrequest.FieldBlockNumber, randomrequest.FieldLogIndex)).
		Limit(w.batchSize).
		All(ctx)
	if err != nil {
		return fmt.Errorf("error while querying event, error=%v", err)
	}

	tasks := make([]*db.TaskCreate, len(events))

	// For per row Randow Requests -> add a task , generateProofResponse => FulfillrandomSeed
	for i, event := range events {
		var trackedEvent contract.RandomBeaconCoordinatorRandomSeedRequested
		if err := json.Unmarshal([]byte(event.Raw), &trackedEvent); err != nil {
			return fmt.Errorf("error while un-marshalling event, event=%v, error=%v", event, err)
		}
		// Gen a Proof based on tracked event Request
		res, err := core.GenerateProofResponse(w.config.VRFConfig.Key, trackedEvent.Period, trackedEvent.Req.PrevBeacon,
			w.config.VRFConfig.KeyHash)

		if err != nil {
			return fmt.Errorf("error while generate random proof, error=%v", err)
		}

		// Create  struct FulFillRandomSee
		request := FulfillRandomSeedParams{trackedEvent.ReqHash, res.Proof, res.RandomRequest}
		data, err := json.Marshal(request)
		if err != nil {
			return fmt.Errorf("error while marshaling params, params=%v, error=%v", request, err)
		}

		// Create tasks and with status pending.
		tasks[i] = w.dbClient.Task.Create().
			SetData(string(data)).
			SetStatus(task.StatusPending).
			AddRandomrequestIDs(event.ID).
			SetCreatedAt(int(time.Now().Unix())).
			SetLastSent(int(time.Now().Unix()))
	}

	_, err = w.dbClient.Task.CreateBulk(tasks...).Save(ctx)
	if err != nil {
		return fmt.Errorf("error while creating tasks, error=%v", err)
	}

	return nil
}
