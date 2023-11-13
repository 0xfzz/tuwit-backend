// Code generated by ent, DO NOT EDIT.

package usercount

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the usercount type in the database.
	Label = "user_count"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFollowerCount holds the string denoting the follower_count field in the database.
	FieldFollowerCount = "follower_count"
	// FieldFollowingsCount holds the string denoting the followings_count field in the database.
	FieldFollowingsCount = "followings_count"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the usercount in the database.
	Table = "user_account"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_account"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_account_user_count_info"
)

// Columns holds all SQL columns for usercount fields.
var Columns = []string{
	FieldID,
	FieldFollowerCount,
	FieldFollowingsCount,
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
	// DefaultFollowerCount holds the default value on creation for the "follower_count" field.
	DefaultFollowerCount int
	// DefaultFollowingsCount holds the default value on creation for the "followings_count" field.
	DefaultFollowingsCount int
)

// OrderOption defines the ordering options for the UserCount queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFollowerCount orders the results by the follower_count field.
func ByFollowerCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFollowerCount, opts...).ToFunc()
}

// ByFollowingsCount orders the results by the followings_count field.
func ByFollowingsCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFollowingsCount, opts...).ToFunc()
}

// ByUserCount orders the results by user count.
func ByUserCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserStep(), opts...)
	}
}

// ByUser orders the results by user terms.
func ByUser(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, UserTable, UserColumn),
	)
}