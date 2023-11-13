// Code generated by ent, DO NOT EDIT.

package threadcount

import (
	"entgo.io/ent/dialect/sql"
	"github.com/0xfzz/tuwitt/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldLTE(FieldID, id))
}

// ReplyCount applies equality check predicate on the "reply_count" field. It's identical to ReplyCountEQ.
func ReplyCount(v int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldEQ(FieldReplyCount, v))
}

// LikeCount applies equality check predicate on the "like_count" field. It's identical to LikeCountEQ.
func LikeCount(v int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldEQ(FieldLikeCount, v))
}

// ReplyCountEQ applies the EQ predicate on the "reply_count" field.
func ReplyCountEQ(v int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldEQ(FieldReplyCount, v))
}

// ReplyCountNEQ applies the NEQ predicate on the "reply_count" field.
func ReplyCountNEQ(v int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldNEQ(FieldReplyCount, v))
}

// ReplyCountIn applies the In predicate on the "reply_count" field.
func ReplyCountIn(vs ...int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldIn(FieldReplyCount, vs...))
}

// ReplyCountNotIn applies the NotIn predicate on the "reply_count" field.
func ReplyCountNotIn(vs ...int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldNotIn(FieldReplyCount, vs...))
}

// ReplyCountGT applies the GT predicate on the "reply_count" field.
func ReplyCountGT(v int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldGT(FieldReplyCount, v))
}

// ReplyCountGTE applies the GTE predicate on the "reply_count" field.
func ReplyCountGTE(v int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldGTE(FieldReplyCount, v))
}

// ReplyCountLT applies the LT predicate on the "reply_count" field.
func ReplyCountLT(v int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldLT(FieldReplyCount, v))
}

// ReplyCountLTE applies the LTE predicate on the "reply_count" field.
func ReplyCountLTE(v int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldLTE(FieldReplyCount, v))
}

// LikeCountEQ applies the EQ predicate on the "like_count" field.
func LikeCountEQ(v int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldEQ(FieldLikeCount, v))
}

// LikeCountNEQ applies the NEQ predicate on the "like_count" field.
func LikeCountNEQ(v int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldNEQ(FieldLikeCount, v))
}

// LikeCountIn applies the In predicate on the "like_count" field.
func LikeCountIn(vs ...int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldIn(FieldLikeCount, vs...))
}

// LikeCountNotIn applies the NotIn predicate on the "like_count" field.
func LikeCountNotIn(vs ...int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldNotIn(FieldLikeCount, vs...))
}

// LikeCountGT applies the GT predicate on the "like_count" field.
func LikeCountGT(v int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldGT(FieldLikeCount, v))
}

// LikeCountGTE applies the GTE predicate on the "like_count" field.
func LikeCountGTE(v int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldGTE(FieldLikeCount, v))
}

// LikeCountLT applies the LT predicate on the "like_count" field.
func LikeCountLT(v int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldLT(FieldLikeCount, v))
}

// LikeCountLTE applies the LTE predicate on the "like_count" field.
func LikeCountLTE(v int) predicate.ThreadCount {
	return predicate.ThreadCount(sql.FieldLTE(FieldLikeCount, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ThreadCount) predicate.ThreadCount {
	return predicate.ThreadCount(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ThreadCount) predicate.ThreadCount {
	return predicate.ThreadCount(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ThreadCount) predicate.ThreadCount {
	return predicate.ThreadCount(sql.NotPredicates(p))
}