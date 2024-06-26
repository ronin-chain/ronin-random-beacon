// Code generated by ent, DO NOT EDIT.

package randomrequest

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the randomrequest type in the database.
	Label = "random_request"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBlockNumber holds the string denoting the blocknumber field in the database.
	FieldBlockNumber = "block_number"
	// FieldLogIndex holds the string denoting the logindex field in the database.
	FieldLogIndex = "log_index"
	// FieldRaw holds the string denoting the raw field in the database.
	FieldRaw = "raw"
	// FieldPeriod holds the string denoting the period field in the database.
	FieldPeriod = "period"
	// EdgeTask holds the string denoting the task edge name in mutations.
	EdgeTask = "task"
	// Table holds the table name of the randomrequest in the database.
	Table = "random_requests"
	// TaskTable is the table that holds the task relation/edge.
	TaskTable = "random_requests"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "tasks"
	// TaskColumn is the table column denoting the task relation/edge.
	TaskColumn = "task_randomrequest"
)

// Columns holds all SQL columns for randomrequest fields.
var Columns = []string{
	FieldID,
	FieldBlockNumber,
	FieldLogIndex,
	FieldRaw,
	FieldPeriod,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "random_requests"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"task_randomrequest",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// RawValidator is a validator for the "raw" field. It is called by the builders before save.
	RawValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the RandomRequest queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBlockNumber orders the results by the blockNumber field.
func ByBlockNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlockNumber, opts...).ToFunc()
}

// ByLogIndex orders the results by the logIndex field.
func ByLogIndex(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogIndex, opts...).ToFunc()
}

// ByRaw orders the results by the raw field.
func ByRaw(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRaw, opts...).ToFunc()
}

// ByPeriod orders the results by the period field.
func ByPeriod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPeriod, opts...).ToFunc()
}

// ByTaskField orders the results by task field.
func ByTaskField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTaskStep(), sql.OrderByField(field, opts...))
	}
}
func newTaskStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TaskInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TaskTable, TaskColumn),
	)
}
