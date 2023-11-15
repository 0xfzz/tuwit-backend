// Code generated by ent, DO NOT EDIT.

package thread

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/0xfzz/tuwitt/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Thread {
	return predicate.Thread(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Thread {
	return predicate.Thread(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Thread {
	return predicate.Thread(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Thread {
	return predicate.Thread(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Thread {
	return predicate.Thread(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Thread {
	return predicate.Thread(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Thread {
	return predicate.Thread(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldUpdatedAt, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldContent, v))
}

// IsCommentDisabled applies equality check predicate on the "is_comment_disabled" field. It's identical to IsCommentDisabledEQ.
func IsCommentDisabled(v bool) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldIsCommentDisabled, v))
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v bool) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldIsDeleted, v))
}

// AuthorID applies equality check predicate on the "author_id" field. It's identical to AuthorIDEQ.
func AuthorID(v int) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldAuthorID, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v int) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldParentID, v))
}

// RepostThreadID applies equality check predicate on the "repost_thread_id" field. It's identical to RepostThreadIDEQ.
func RepostThreadID(v int) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldRepostThreadID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Thread {
	return predicate.Thread(sql.FieldLTE(FieldUpdatedAt, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Thread {
	return predicate.Thread(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Thread {
	return predicate.Thread(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Thread {
	return predicate.Thread(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Thread {
	return predicate.Thread(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Thread {
	return predicate.Thread(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Thread {
	return predicate.Thread(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Thread {
	return predicate.Thread(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Thread {
	return predicate.Thread(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Thread {
	return predicate.Thread(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Thread {
	return predicate.Thread(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Thread {
	return predicate.Thread(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Thread {
	return predicate.Thread(sql.FieldContainsFold(FieldContent, v))
}

// IsCommentDisabledEQ applies the EQ predicate on the "is_comment_disabled" field.
func IsCommentDisabledEQ(v bool) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldIsCommentDisabled, v))
}

// IsCommentDisabledNEQ applies the NEQ predicate on the "is_comment_disabled" field.
func IsCommentDisabledNEQ(v bool) predicate.Thread {
	return predicate.Thread(sql.FieldNEQ(FieldIsCommentDisabled, v))
}

// VisibilityEQ applies the EQ predicate on the "visibility" field.
func VisibilityEQ(v Visibility) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldVisibility, v))
}

// VisibilityNEQ applies the NEQ predicate on the "visibility" field.
func VisibilityNEQ(v Visibility) predicate.Thread {
	return predicate.Thread(sql.FieldNEQ(FieldVisibility, v))
}

// VisibilityIn applies the In predicate on the "visibility" field.
func VisibilityIn(vs ...Visibility) predicate.Thread {
	return predicate.Thread(sql.FieldIn(FieldVisibility, vs...))
}

// VisibilityNotIn applies the NotIn predicate on the "visibility" field.
func VisibilityNotIn(vs ...Visibility) predicate.Thread {
	return predicate.Thread(sql.FieldNotIn(FieldVisibility, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Thread {
	return predicate.Thread(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Thread {
	return predicate.Thread(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Thread {
	return predicate.Thread(sql.FieldNotIn(FieldStatus, vs...))
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v bool) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldIsDeleted, v))
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v bool) predicate.Thread {
	return predicate.Thread(sql.FieldNEQ(FieldIsDeleted, v))
}

// AuthorIDEQ applies the EQ predicate on the "author_id" field.
func AuthorIDEQ(v int) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldAuthorID, v))
}

// AuthorIDNEQ applies the NEQ predicate on the "author_id" field.
func AuthorIDNEQ(v int) predicate.Thread {
	return predicate.Thread(sql.FieldNEQ(FieldAuthorID, v))
}

// AuthorIDIn applies the In predicate on the "author_id" field.
func AuthorIDIn(vs ...int) predicate.Thread {
	return predicate.Thread(sql.FieldIn(FieldAuthorID, vs...))
}

// AuthorIDNotIn applies the NotIn predicate on the "author_id" field.
func AuthorIDNotIn(vs ...int) predicate.Thread {
	return predicate.Thread(sql.FieldNotIn(FieldAuthorID, vs...))
}

// AuthorIDIsNil applies the IsNil predicate on the "author_id" field.
func AuthorIDIsNil() predicate.Thread {
	return predicate.Thread(sql.FieldIsNull(FieldAuthorID))
}

// AuthorIDNotNil applies the NotNil predicate on the "author_id" field.
func AuthorIDNotNil() predicate.Thread {
	return predicate.Thread(sql.FieldNotNull(FieldAuthorID))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v int) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v int) predicate.Thread {
	return predicate.Thread(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...int) predicate.Thread {
	return predicate.Thread(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...int) predicate.Thread {
	return predicate.Thread(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.Thread {
	return predicate.Thread(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.Thread {
	return predicate.Thread(sql.FieldNotNull(FieldParentID))
}

// RepostThreadIDEQ applies the EQ predicate on the "repost_thread_id" field.
func RepostThreadIDEQ(v int) predicate.Thread {
	return predicate.Thread(sql.FieldEQ(FieldRepostThreadID, v))
}

// RepostThreadIDNEQ applies the NEQ predicate on the "repost_thread_id" field.
func RepostThreadIDNEQ(v int) predicate.Thread {
	return predicate.Thread(sql.FieldNEQ(FieldRepostThreadID, v))
}

// RepostThreadIDIn applies the In predicate on the "repost_thread_id" field.
func RepostThreadIDIn(vs ...int) predicate.Thread {
	return predicate.Thread(sql.FieldIn(FieldRepostThreadID, vs...))
}

// RepostThreadIDNotIn applies the NotIn predicate on the "repost_thread_id" field.
func RepostThreadIDNotIn(vs ...int) predicate.Thread {
	return predicate.Thread(sql.FieldNotIn(FieldRepostThreadID, vs...))
}

// RepostThreadIDIsNil applies the IsNil predicate on the "repost_thread_id" field.
func RepostThreadIDIsNil() predicate.Thread {
	return predicate.Thread(sql.FieldIsNull(FieldRepostThreadID))
}

// RepostThreadIDNotNil applies the NotNil predicate on the "repost_thread_id" field.
func RepostThreadIDNotNil() predicate.Thread {
	return predicate.Thread(sql.FieldNotNull(FieldRepostThreadID))
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Thread {
	return predicate.Thread(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.UserAccount) predicate.Thread {
	return predicate.Thread(func(s *sql.Selector) {
		step := newAuthorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Thread {
	return predicate.Thread(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Thread) predicate.Thread {
	return predicate.Thread(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Thread {
	return predicate.Thread(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Thread) predicate.Thread {
	return predicate.Thread(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasThreadCount applies the HasEdge predicate on the "thread_count" edge.
func HasThreadCount() predicate.Thread {
	return predicate.Thread(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ThreadCountTable, ThreadCountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasThreadCountWith applies the HasEdge predicate on the "thread_count" edge with a given conditions (other predicates).
func HasThreadCountWith(preds ...predicate.ThreadCount) predicate.Thread {
	return predicate.Thread(func(s *sql.Selector) {
		step := newThreadCountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReposted applies the HasEdge predicate on the "reposted" edge.
func HasReposted() predicate.Thread {
	return predicate.Thread(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, RepostedTable, RepostedColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepostedWith applies the HasEdge predicate on the "reposted" edge with a given conditions (other predicates).
func HasRepostedWith(preds ...predicate.Thread) predicate.Thread {
	return predicate.Thread(func(s *sql.Selector) {
		step := newRepostedStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRepost applies the HasEdge predicate on the "repost" edge.
func HasRepost() predicate.Thread {
	return predicate.Thread(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, RepostTable, RepostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepostWith applies the HasEdge predicate on the "repost" edge with a given conditions (other predicates).
func HasRepostWith(preds ...predicate.Thread) predicate.Thread {
	return predicate.Thread(func(s *sql.Selector) {
		step := newRepostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasImages applies the HasEdge predicate on the "images" edge.
func HasImages() predicate.Thread {
	return predicate.Thread(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ImagesTable, ImagesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImagesWith applies the HasEdge predicate on the "images" edge with a given conditions (other predicates).
func HasImagesWith(preds ...predicate.Media) predicate.Thread {
	return predicate.Thread(func(s *sql.Selector) {
		step := newImagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Thread) predicate.Thread {
	return predicate.Thread(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Thread) predicate.Thread {
	return predicate.Thread(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Thread) predicate.Thread {
	return predicate.Thread(sql.NotPredicates(p))
}
