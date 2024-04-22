// Code generated by ent, DO NOT EDIT.

package chaininfo

import (
	"entgo.io/ent/dialect/sql"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldContainsFold(FieldID, id))
}

// ProcessedBlock applies equality check predicate on the "processed_block" field. It's identical to ProcessedBlockEQ.
func ProcessedBlock(v uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldEQ(FieldProcessedBlock, v))
}

// FinalizedBlock applies equality check predicate on the "finalized_block" field. It's identical to FinalizedBlockEQ.
func FinalizedBlock(v uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldEQ(FieldFinalizedBlock, v))
}

// ProcessedBlockEQ applies the EQ predicate on the "processed_block" field.
func ProcessedBlockEQ(v uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldEQ(FieldProcessedBlock, v))
}

// ProcessedBlockNEQ applies the NEQ predicate on the "processed_block" field.
func ProcessedBlockNEQ(v uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldNEQ(FieldProcessedBlock, v))
}

// ProcessedBlockIn applies the In predicate on the "processed_block" field.
func ProcessedBlockIn(vs ...uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldIn(FieldProcessedBlock, vs...))
}

// ProcessedBlockNotIn applies the NotIn predicate on the "processed_block" field.
func ProcessedBlockNotIn(vs ...uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldNotIn(FieldProcessedBlock, vs...))
}

// ProcessedBlockGT applies the GT predicate on the "processed_block" field.
func ProcessedBlockGT(v uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldGT(FieldProcessedBlock, v))
}

// ProcessedBlockGTE applies the GTE predicate on the "processed_block" field.
func ProcessedBlockGTE(v uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldGTE(FieldProcessedBlock, v))
}

// ProcessedBlockLT applies the LT predicate on the "processed_block" field.
func ProcessedBlockLT(v uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldLT(FieldProcessedBlock, v))
}

// ProcessedBlockLTE applies the LTE predicate on the "processed_block" field.
func ProcessedBlockLTE(v uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldLTE(FieldProcessedBlock, v))
}

// FinalizedBlockEQ applies the EQ predicate on the "finalized_block" field.
func FinalizedBlockEQ(v uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldEQ(FieldFinalizedBlock, v))
}

// FinalizedBlockNEQ applies the NEQ predicate on the "finalized_block" field.
func FinalizedBlockNEQ(v uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldNEQ(FieldFinalizedBlock, v))
}

// FinalizedBlockIn applies the In predicate on the "finalized_block" field.
func FinalizedBlockIn(vs ...uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldIn(FieldFinalizedBlock, vs...))
}

// FinalizedBlockNotIn applies the NotIn predicate on the "finalized_block" field.
func FinalizedBlockNotIn(vs ...uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldNotIn(FieldFinalizedBlock, vs...))
}

// FinalizedBlockGT applies the GT predicate on the "finalized_block" field.
func FinalizedBlockGT(v uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldGT(FieldFinalizedBlock, v))
}

// FinalizedBlockGTE applies the GTE predicate on the "finalized_block" field.
func FinalizedBlockGTE(v uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldGTE(FieldFinalizedBlock, v))
}

// FinalizedBlockLT applies the LT predicate on the "finalized_block" field.
func FinalizedBlockLT(v uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldLT(FieldFinalizedBlock, v))
}

// FinalizedBlockLTE applies the LTE predicate on the "finalized_block" field.
func FinalizedBlockLTE(v uint64) predicate.ChainInfo {
	return predicate.ChainInfo(sql.FieldLTE(FieldFinalizedBlock, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ChainInfo) predicate.ChainInfo {
	return predicate.ChainInfo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ChainInfo) predicate.ChainInfo {
	return predicate.ChainInfo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ChainInfo) predicate.ChainInfo {
	return predicate.ChainInfo(sql.NotPredicates(p))
}
