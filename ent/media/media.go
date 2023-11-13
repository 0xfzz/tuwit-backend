// Code generated by ent, DO NOT EDIT.

package media

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the media type in the database.
	Label = "media"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// EdgeThreads holds the string denoting the threads edge name in mutations.
	EdgeThreads = "threads"
	// EdgeOwnerProfilePicture holds the string denoting the owner_profile_picture edge name in mutations.
	EdgeOwnerProfilePicture = "owner_profile_picture"
	// EdgeOwnerBanner holds the string denoting the owner_banner edge name in mutations.
	EdgeOwnerBanner = "owner_banner"
	// Table holds the table name of the media in the database.
	Table = "media"
	// ThreadsTable is the table that holds the threads relation/edge. The primary key declared below.
	ThreadsTable = "thread_images"
	// ThreadsInverseTable is the table name for the Thread entity.
	// It exists in this package in order to avoid circular dependency with the "thread" package.
	ThreadsInverseTable = "threads"
	// OwnerProfilePictureTable is the table that holds the owner_profile_picture relation/edge.
	OwnerProfilePictureTable = "user_profile"
	// OwnerProfilePictureInverseTable is the table name for the UserProfile entity.
	// It exists in this package in order to avoid circular dependency with the "userprofile" package.
	OwnerProfilePictureInverseTable = "user_profile"
	// OwnerProfilePictureColumn is the table column denoting the owner_profile_picture relation/edge.
	OwnerProfilePictureColumn = "user_profile_profile_picture"
	// OwnerBannerTable is the table that holds the owner_banner relation/edge.
	OwnerBannerTable = "user_profile"
	// OwnerBannerInverseTable is the table name for the UserProfile entity.
	// It exists in this package in order to avoid circular dependency with the "userprofile" package.
	OwnerBannerInverseTable = "user_profile"
	// OwnerBannerColumn is the table column denoting the owner_banner relation/edge.
	OwnerBannerColumn = "user_profile_banner"
)

// Columns holds all SQL columns for media fields.
var Columns = []string{
	FieldID,
	FieldPath,
	FieldCategory,
}

var (
	// ThreadsPrimaryKey and ThreadsColumn2 are the table columns denoting the
	// primary key for the threads relation (M2M).
	ThreadsPrimaryKey = []string{"thread_id", "media_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Category defines the type for the "category" enum field.
type Category string

// Category values.
const (
	CategoryIMAGE    Category = "IMAGE"
	CategoryDOCUMENT Category = "DOCUMENT"
	CategoryOTHER    Category = "OTHER"
)

func (c Category) String() string {
	return string(c)
}

// CategoryValidator is a validator for the "category" field enum values. It is called by the builders before save.
func CategoryValidator(c Category) error {
	switch c {
	case CategoryIMAGE, CategoryDOCUMENT, CategoryOTHER:
		return nil
	default:
		return fmt.Errorf("media: invalid enum value for category field: %q", c)
	}
}

// OrderOption defines the ordering options for the Media queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
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

// ByOwnerProfilePictureCount orders the results by owner_profile_picture count.
func ByOwnerProfilePictureCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOwnerProfilePictureStep(), opts...)
	}
}

// ByOwnerProfilePicture orders the results by owner_profile_picture terms.
func ByOwnerProfilePicture(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerProfilePictureStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOwnerBannerCount orders the results by owner_banner count.
func ByOwnerBannerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOwnerBannerStep(), opts...)
	}
}

// ByOwnerBanner orders the results by owner_banner terms.
func ByOwnerBanner(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerBannerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newThreadsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ThreadsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ThreadsTable, ThreadsPrimaryKey...),
	)
}
func newOwnerProfilePictureStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerProfilePictureInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, OwnerProfilePictureTable, OwnerProfilePictureColumn),
	)
}
func newOwnerBannerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerBannerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, OwnerBannerTable, OwnerBannerColumn),
	)
}
