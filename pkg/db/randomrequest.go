// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/randomrequest"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/task"
)

// RandomRequest is the model entity for the RandomRequest schema.
type RandomRequest struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// BlockNumber holds the value of the "blockNumber" field.
	BlockNumber uint64 `json:"blockNumber,omitempty"`
	// LogIndex holds the value of the "logIndex" field.
	LogIndex uint `json:"logIndex,omitempty"`
	// Raw holds the value of the "raw" field.
	Raw string `json:"raw,omitempty"`
	// Period holds the value of the "period" field.
	Period uint64 `json:"period,omitempty"`
	// PreviousBeacon holds the value of the "previousBeacon" field.
	PreviousBeacon uint64 `json:"previousBeacon,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RandomRequestQuery when eager-loading is set.
	Edges              RandomRequestEdges `json:"edges"`
	task_randomrequest *int
	selectValues       sql.SelectValues
}

// RandomRequestEdges holds the relations/edges for other nodes in the graph.
type RandomRequestEdges struct {
	// Task holds the value of the task edge.
	Task *Task `json:"task,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RandomRequestEdges) TaskOrErr() (*Task, error) {
	if e.loadedTypes[0] {
		if e.Task == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: task.Label}
		}
		return e.Task, nil
	}
	return nil, &NotLoadedError{edge: "task"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RandomRequest) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case randomrequest.FieldBlockNumber, randomrequest.FieldLogIndex, randomrequest.FieldPeriod, randomrequest.FieldPreviousBeacon:
			values[i] = new(sql.NullInt64)
		case randomrequest.FieldID, randomrequest.FieldRaw:
			values[i] = new(sql.NullString)
		case randomrequest.ForeignKeys[0]: // task_randomrequest
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RandomRequest fields.
func (rr *RandomRequest) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case randomrequest.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				rr.ID = value.String
			}
		case randomrequest.FieldBlockNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field blockNumber", values[i])
			} else if value.Valid {
				rr.BlockNumber = uint64(value.Int64)
			}
		case randomrequest.FieldLogIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field logIndex", values[i])
			} else if value.Valid {
				rr.LogIndex = uint(value.Int64)
			}
		case randomrequest.FieldRaw:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field raw", values[i])
			} else if value.Valid {
				rr.Raw = value.String
			}
		case randomrequest.FieldPeriod:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field period", values[i])
			} else if value.Valid {
				rr.Period = uint64(value.Int64)
			}
		case randomrequest.FieldPreviousBeacon:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field previousBeacon", values[i])
			} else if value.Valid {
				rr.PreviousBeacon = uint64(value.Int64)
			}
		case randomrequest.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field task_randomrequest", value)
			} else if value.Valid {
				rr.task_randomrequest = new(int)
				*rr.task_randomrequest = int(value.Int64)
			}
		default:
			rr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RandomRequest.
// This includes values selected through modifiers, order, etc.
func (rr *RandomRequest) Value(name string) (ent.Value, error) {
	return rr.selectValues.Get(name)
}

// QueryTask queries the "task" edge of the RandomRequest entity.
func (rr *RandomRequest) QueryTask() *TaskQuery {
	return NewRandomRequestClient(rr.config).QueryTask(rr)
}

// Update returns a builder for updating this RandomRequest.
// Note that you need to call RandomRequest.Unwrap() before calling this method if this RandomRequest
// was returned from a transaction, and the transaction was committed or rolled back.
func (rr *RandomRequest) Update() *RandomRequestUpdateOne {
	return NewRandomRequestClient(rr.config).UpdateOne(rr)
}

// Unwrap unwraps the RandomRequest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rr *RandomRequest) Unwrap() *RandomRequest {
	_tx, ok := rr.config.driver.(*txDriver)
	if !ok {
		panic("db: RandomRequest is not a transactional entity")
	}
	rr.config.driver = _tx.drv
	return rr
}

// String implements the fmt.Stringer.
func (rr *RandomRequest) String() string {
	var builder strings.Builder
	builder.WriteString("RandomRequest(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rr.ID))
	builder.WriteString("blockNumber=")
	builder.WriteString(fmt.Sprintf("%v", rr.BlockNumber))
	builder.WriteString(", ")
	builder.WriteString("logIndex=")
	builder.WriteString(fmt.Sprintf("%v", rr.LogIndex))
	builder.WriteString(", ")
	builder.WriteString("raw=")
	builder.WriteString(rr.Raw)
	builder.WriteString(", ")
	builder.WriteString("period=")
	builder.WriteString(fmt.Sprintf("%v", rr.Period))
	builder.WriteString(", ")
	builder.WriteString("previousBeacon=")
	builder.WriteString(fmt.Sprintf("%v", rr.PreviousBeacon))
	builder.WriteByte(')')
	return builder.String()
}

// RandomRequests is a parsable slice of RandomRequest.
type RandomRequests []*RandomRequest