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
	Table = "user_count"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_count"
	// UserInverseTable is the table name for the UserAccount entity.
	// It exists in this package in order to avoid circular dependency with the "useraccount" package.
	UserInverseTable = "user_account"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_account_user_count"
)

// Columns holds all SQL columns for usercount fields.
var Columns = []string{
	FieldID,
	FieldFollowerCount,
	FieldFollowingsCount,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "user_count"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_account_user_count",
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

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
	)
}
