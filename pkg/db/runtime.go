// Code generated by ent, DO NOT EDIT.

package db

import (
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/chaininfo"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/randomrequest"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/task"
	"github.com/ronin-chain/ronin-random-beacon/pkg/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	chaininfoFields := schema.ChainInfo{}.Fields()
	_ = chaininfoFields
	// chaininfoDescProcessedBlock is the schema descriptor for processed_block field.
	chaininfoDescProcessedBlock := chaininfoFields[1].Descriptor()
	// chaininfo.DefaultProcessedBlock holds the default value on creation for the processed_block field.
	chaininfo.DefaultProcessedBlock = chaininfoDescProcessedBlock.Default.(uint64)
	// chaininfoDescID is the schema descriptor for id field.
	chaininfoDescID := chaininfoFields[0].Descriptor()
	// chaininfo.IDValidator is a validator for the "id" field. It is called by the builders before save.
	chaininfo.IDValidator = chaininfoDescID.Validators[0].(func(string) error)
	randomrequestFields := schema.RandomRequest{}.Fields()
	_ = randomrequestFields
	// randomrequestDescRaw is the schema descriptor for raw field.
	randomrequestDescRaw := randomrequestFields[3].Descriptor()
	// randomrequest.RawValidator is a validator for the "raw" field. It is called by the builders before save.
	randomrequest.RawValidator = randomrequestDescRaw.Validators[0].(func(string) error)
	// randomrequestDescID is the schema descriptor for id field.
	randomrequestDescID := randomrequestFields[0].Descriptor()
	// randomrequest.IDValidator is a validator for the "id" field. It is called by the builders before save.
	randomrequest.IDValidator = randomrequestDescID.Validators[0].(func(string) error)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescData is the schema descriptor for data field.
	taskDescData := taskFields[1].Descriptor()
	// task.DataValidator is a validator for the "data" field. It is called by the builders before save.
	task.DataValidator = taskDescData.Validators[0].(func(string) error)
	// taskDescReOrg is the schema descriptor for reOrg field.
	taskDescReOrg := taskFields[4].Descriptor()
	// task.DefaultReOrg holds the default value on creation for the reOrg field.
	task.DefaultReOrg = taskDescReOrg.Default.(bool)
	// taskDescAttempts is the schema descriptor for attempts field.
	taskDescAttempts := taskFields[5].Descriptor()
	// task.DefaultAttempts holds the default value on creation for the attempts field.
	task.DefaultAttempts = taskDescAttempts.Default.(int)
	// taskDescIsFinalized is the schema descriptor for isFinalized field.
	taskDescIsFinalized := taskFields[7].Descriptor()
	// task.DefaultIsFinalized holds the default value on creation for the isFinalized field.
	task.DefaultIsFinalized = taskDescIsFinalized.Default.(bool)
	// taskDescSentBlock is the schema descriptor for sentBlock field.
	taskDescSentBlock := taskFields[8].Descriptor()
	// task.DefaultSentBlock holds the default value on creation for the sentBlock field.
	task.DefaultSentBlock = taskDescSentBlock.Default.(int)
	// taskDescCreatedAt is the schema descriptor for createdAt field.
	taskDescCreatedAt := taskFields[9].Descriptor()
	// task.CreatedAtValidator is a validator for the "createdAt" field. It is called by the builders before save.
	task.CreatedAtValidator = taskDescCreatedAt.Validators[0].(func(int) error)
}
