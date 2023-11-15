// Code generated by ent, DO NOT EDIT.

package useraccount

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the useraccount type in the database.
	Label = "user_account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldIsVerified holds the string denoting the is_verified field in the database.
	FieldIsVerified = "is_verified"
	// FieldIsPrivate holds the string denoting the is_private field in the database.
	FieldIsPrivate = "is_private"
	// FieldIsEmailVerified holds the string denoting the is_email_verified field in the database.
	FieldIsEmailVerified = "is_email_verified"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// EdgeFollowers holds the string denoting the followers edge name in mutations.
	EdgeFollowers = "followers"
	// EdgeFollowings holds the string denoting the followings edge name in mutations.
	EdgeFollowings = "followings"
	// EdgeBlockedBy holds the string denoting the blocked_by edge name in mutations.
	EdgeBlockedBy = "blocked_by"
	// EdgeBlockedUsers holds the string denoting the blocked_users edge name in mutations.
	EdgeBlockedUsers = "blocked_users"
	// EdgeUserCount holds the string denoting the user_count edge name in mutations.
	EdgeUserCount = "user_count"
	// EdgeThreads holds the string denoting the threads edge name in mutations.
	EdgeThreads = "threads"
	// Table holds the table name of the useraccount in the database.
	Table = "user_account"
	// ProfileTable is the table that holds the profile relation/edge.
	ProfileTable = "user_profile"
	// ProfileInverseTable is the table name for the UserProfile entity.
	// It exists in this package in order to avoid circular dependency with the "userprofile" package.
	ProfileInverseTable = "user_profile"
	// ProfileColumn is the table column denoting the profile relation/edge.
	ProfileColumn = "user_account_profile"
	// FollowersTable is the table that holds the followers relation/edge.
	FollowersTable = "user_follower_relationships"
	// FollowersInverseTable is the table name for the UserFollowerRelationship entity.
	// It exists in this package in order to avoid circular dependency with the "userfollowerrelationship" package.
	FollowersInverseTable = "user_follower_relationships"
	// FollowersColumn is the table column denoting the followers relation/edge.
	FollowersColumn = "user_id"
	// FollowingsTable is the table that holds the followings relation/edge.
	FollowingsTable = "user_follower_relationships"
	// FollowingsInverseTable is the table name for the UserFollowerRelationship entity.
	// It exists in this package in order to avoid circular dependency with the "userfollowerrelationship" package.
	FollowingsInverseTable = "user_follower_relationships"
	// FollowingsColumn is the table column denoting the followings relation/edge.
	FollowingsColumn = "follower_id"
	// BlockedByTable is the table that holds the blocked_by relation/edge.
	BlockedByTable = "blocked_users_relationships"
	// BlockedByInverseTable is the table name for the BlockedUsersRelationship entity.
	// It exists in this package in order to avoid circular dependency with the "blockedusersrelationship" package.
	BlockedByInverseTable = "blocked_users_relationships"
	// BlockedByColumn is the table column denoting the blocked_by relation/edge.
	BlockedByColumn = "user_id"
	// BlockedUsersTable is the table that holds the blocked_users relation/edge.
	BlockedUsersTable = "blocked_users_relationships"
	// BlockedUsersInverseTable is the table name for the BlockedUsersRelationship entity.
	// It exists in this package in order to avoid circular dependency with the "blockedusersrelationship" package.
	BlockedUsersInverseTable = "blocked_users_relationships"
	// BlockedUsersColumn is the table column denoting the blocked_users relation/edge.
	BlockedUsersColumn = "blocker_id"
	// UserCountTable is the table that holds the user_count relation/edge.
	UserCountTable = "user_count"
	// UserCountInverseTable is the table name for the UserCount entity.
	// It exists in this package in order to avoid circular dependency with the "usercount" package.
	UserCountInverseTable = "user_count"
	// UserCountColumn is the table column denoting the user_count relation/edge.
	UserCountColumn = "user_account_user_count"
	// ThreadsTable is the table that holds the threads relation/edge.
	ThreadsTable = "threads"
	// ThreadsInverseTable is the table name for the Thread entity.
	// It exists in this package in order to avoid circular dependency with the "thread" package.
	ThreadsInverseTable = "threads"
	// ThreadsColumn is the table column denoting the threads relation/edge.
	ThreadsColumn = "author_id"
)

// Columns holds all SQL columns for useraccount fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldUsername,
	FieldPassword,
	FieldIsVerified,
	FieldIsPrivate,
	FieldIsEmailVerified,
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
	// DefaultIsVerified holds the default value on creation for the "is_verified" field.
	DefaultIsVerified bool
	// DefaultIsPrivate holds the default value on creation for the "is_private" field.
	DefaultIsPrivate bool
	// DefaultIsEmailVerified holds the default value on creation for the "is_email_verified" field.
	DefaultIsEmailVerified bool
)

// OrderOption defines the ordering options for the UserAccount queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByIsVerified orders the results by the is_verified field.
func ByIsVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsVerified, opts...).ToFunc()
}

// ByIsPrivate orders the results by the is_private field.
func ByIsPrivate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPrivate, opts...).ToFunc()
}

// ByIsEmailVerified orders the results by the is_email_verified field.
func ByIsEmailVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsEmailVerified, opts...).ToFunc()
}

// ByProfileField orders the results by profile field.
func ByProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProfileStep(), sql.OrderByField(field, opts...))
	}
}

// ByFollowersCount orders the results by followers count.
func ByFollowersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFollowersStep(), opts...)
	}
}

// ByFollowers orders the results by followers terms.
func ByFollowers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFollowersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFollowingsCount orders the results by followings count.
func ByFollowingsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFollowingsStep(), opts...)
	}
}

// ByFollowings orders the results by followings terms.
func ByFollowings(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFollowingsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBlockedByCount orders the results by blocked_by count.
func ByBlockedByCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBlockedByStep(), opts...)
	}
}

// ByBlockedBy orders the results by blocked_by terms.
func ByBlockedBy(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBlockedByStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBlockedUsersCount orders the results by blocked_users count.
func ByBlockedUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBlockedUsersStep(), opts...)
	}
}

// ByBlockedUsers orders the results by blocked_users terms.
func ByBlockedUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBlockedUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserCountField orders the results by user_count field.
func ByUserCountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserCountStep(), sql.OrderByField(field, opts...))
	}
}

// ByThreadsCount orders the results by threads count.
func ByThreadsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newThreadsStep(), opts...)
	}
}

// ByThreads orders the results by threads terms.
func ByThreads(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newThreadsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ProfileTable, ProfileColumn),
	)
}
func newFollowersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FollowersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FollowersTable, FollowersColumn),
	)
}
func newFollowingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FollowingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FollowingsTable, FollowingsColumn),
	)
}
func newBlockedByStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BlockedByInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BlockedByTable, BlockedByColumn),
	)
}
func newBlockedUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BlockedUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BlockedUsersTable, BlockedUsersColumn),
	)
}
func newUserCountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserCountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, UserCountTable, UserCountColumn),
	)
}
func newThreadsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ThreadsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ThreadsTable, ThreadsColumn),
	)
}
