// Code generated by ent, DO NOT EDIT.

package media

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/0xfzz/tuwitt/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldID, id))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldPath, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Media {
	return predicate.Media(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Media {
	return predicate.Media(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Media {
	return predicate.Media(sql.FieldContainsFold(FieldPath, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v Category) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v Category) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...Category) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...Category) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldCategory, vs...))
}

// HasThreads applies the HasEdge predicate on the "threads" edge.
func HasThreads() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ThreadsTable, ThreadsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasThreadsWith applies the HasEdge predicate on the "threads" edge with a given conditions (other predicates).
func HasThreadsWith(preds ...predicate.Thread) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := newThreadsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwnerProfilePicture applies the HasEdge predicate on the "owner_profile_picture" edge.
func HasOwnerProfilePicture() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, OwnerProfilePictureTable, OwnerProfilePictureColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerProfilePictureWith applies the HasEdge predicate on the "owner_profile_picture" edge with a given conditions (other predicates).
func HasOwnerProfilePictureWith(preds ...predicate.UserProfile) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := newOwnerProfilePictureStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwnerBanner applies the HasEdge predicate on the "owner_banner" edge.
func HasOwnerBanner() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, OwnerBannerTable, OwnerBannerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerBannerWith applies the HasEdge predicate on the "owner_banner" edge with a given conditions (other predicates).
func HasOwnerBannerWith(preds ...predicate.UserProfile) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := newOwnerBannerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Media) predicate.Media {
	return predicate.Media(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Media) predicate.Media {
	return predicate.Media(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Media) predicate.Media {
	return predicate.Media(sql.NotPredicates(p))
}