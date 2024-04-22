// Code generated by ent, DO NOT EDIT.

package task

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTxHash holds the string denoting the txhash field in the database.
	FieldTxHash = "tx_hash"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldLastError holds the string denoting the lasterror field in the database.
	FieldLastError = "last_error"
	// FieldReOrg holds the string denoting the reorg field in the database.
	FieldReOrg = "re_org"
	// FieldAttempts holds the string denoting the attempts field in the database.
	FieldAttempts = "attempts"
	// FieldLastSent holds the string denoting the lastsent field in the database.
	FieldLastSent = "last_sent"
	// FieldIsFinalized holds the string denoting the isfinalized field in the database.
	FieldIsFinalized = "is_finalized"
	// FieldSentBlock holds the string denoting the sentblock field in the database.
	FieldSentBlock = "sent_block"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// EdgeRandomrequest holds the string denoting the randomrequest edge name in mutations.
	EdgeRandomrequest = "randomrequest"
	// Table holds the table name of the task in the database.
	Table = "tasks"
	// RandomrequestTable is the table that holds the randomrequest relation/edge.
	RandomrequestTable = "random_requests"
	// RandomrequestInverseTable is the table name for the RandomRequest entity.
	// It exists in this package in order to avoid circular dependency with the "randomrequest" package.
	RandomrequestInverseTable = "random_requests"
	// RandomrequestColumn is the table column denoting the randomrequest relation/edge.
	RandomrequestColumn = "task_randomrequest"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldTxHash,
	FieldData,
	FieldStatus,
	FieldLastError,
	FieldReOrg,
	FieldAttempts,
	FieldLastSent,
	FieldIsFinalized,
	FieldSentBlock,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DataValidator is a validator for the "data" field. It is called by the builders before save.
	DataValidator func(string) error
	// DefaultReOrg holds the default value on creation for the "reOrg" field.
	DefaultReOrg bool
	// DefaultAttempts holds the default value on creation for the "attempts" field.
	DefaultAttempts int
	// DefaultIsFinalized holds the default value on creation for the "isFinalized" field.
	DefaultIsFinalized bool
	// DefaultSentBlock holds the default value on creation for the "sentBlock" field.
	DefaultSentBlock int
	// CreatedAtValidator is a validator for the "createdAt" field. It is called by the builders before save.
	CreatedAtValidator func(int) error
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusPending Status = "pending"
	StatusSent    Status = "sent"
	StatusSuccess Status = "success"
	StatusFailed  Status = "failed"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPending, StatusSent, StatusSuccess, StatusFailed:
		return nil
	default:
		return fmt.Errorf("task: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Task queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTxHash orders the results by the txHash field.
func ByTxHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTxHash, opts...).ToFunc()
}

// ByData orders the results by the data field.
func ByData(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldData, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByLastError orders the results by the lastError field.
func ByLastError(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastError, opts...).ToFunc()
}

// ByReOrg orders the results by the reOrg field.
func ByReOrg(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReOrg, opts...).ToFunc()
}

// ByAttempts orders the results by the attempts field.
func ByAttempts(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAttempts, opts...).ToFunc()
}

// ByLastSent orders the results by the lastSent field.
func ByLastSent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastSent, opts...).ToFunc()
}

// ByIsFinalized orders the results by the isFinalized field.
func ByIsFinalized(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsFinalized, opts...).ToFunc()
}

// BySentBlock orders the results by the sentBlock field.
func BySentBlock(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSentBlock, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByRandomrequestCount orders the results by randomrequest count.
func ByRandomrequestCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRandomrequestStep(), opts...)
	}
}

// ByRandomrequest orders the results by randomrequest terms.
func ByRandomrequest(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRandomrequestStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRandomrequestStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RandomrequestInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RandomrequestTable, RandomrequestColumn),
	)
}
