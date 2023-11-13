// Code generated by ent, DO NOT EDIT.

package thread

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the thread type in the database.
	Label = "thread"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldIsCommentDisabled holds the string denoting the is_comment_disabled field in the database.
	FieldIsCommentDisabled = "is_comment_disabled"
	// FieldVisibility holds the string denoting the visibility field in the database.
	FieldVisibility = "visibility"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// EdgeParentThread holds the string denoting the parent_thread edge name in mutations.
	EdgeParentThread = "parent_thread"
	// EdgeChildThreads holds the string denoting the child_threads edge name in mutations.
	EdgeChildThreads = "child_threads"
	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"
	// Table holds the table name of the thread in the database.
	Table = "threads"
	// ParentThreadTable is the table that holds the parent_thread relation/edge.
	ParentThreadTable = "threads"
	// ParentThreadColumn is the table column denoting the parent_thread relation/edge.
	ParentThreadColumn = "thread_child_threads"
	// ChildThreadsTable is the table that holds the child_threads relation/edge.
	ChildThreadsTable = "threads"
	// ChildThreadsColumn is the table column denoting the child_threads relation/edge.
	ChildThreadsColumn = "thread_child_threads"
	// ImagesTable is the table that holds the images relation/edge. The primary key declared below.
	ImagesTable = "thread_images"
	// ImagesInverseTable is the table name for the Media entity.
	// It exists in this package in order to avoid circular dependency with the "media" package.
	ImagesInverseTable = "media"
)

// Columns holds all SQL columns for thread fields.
var Columns = []string{
	FieldID,
	FieldContent,
	FieldIsCommentDisabled,
	FieldVisibility,
	FieldStatus,
	FieldIsDeleted,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "threads"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"thread_child_threads",
}

var (
	// ImagesPrimaryKey and ImagesColumn2 are the table columns denoting the
	// primary key for the images relation (M2M).
	ImagesPrimaryKey = []string{"thread_id", "media_id"}
)

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
	// DefaultIsCommentDisabled holds the default value on creation for the "is_comment_disabled" field.
	DefaultIsCommentDisabled bool
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted bool
)

// Visibility defines the type for the "visibility" enum field.
type Visibility string

// Visibility values.
const (
	VisibilityPRIVATE Visibility = "PRIVATE"
	VisibilityPUBLIC  Visibility = "PUBLIC"
)

func (v Visibility) String() string {
	return string(v)
}

// VisibilityValidator is a validator for the "visibility" field enum values. It is called by the builders before save.
func VisibilityValidator(v Visibility) error {
	switch v {
	case VisibilityPRIVATE, VisibilityPUBLIC:
		return nil
	default:
		return fmt.Errorf("thread: invalid enum value for visibility field: %q", v)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusPUBLISHED Status = "PUBLISHED"
	StatusDRAFT     Status = "DRAFT"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPUBLISHED, StatusDRAFT:
		return nil
	default:
		return fmt.Errorf("thread: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Thread queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByIsCommentDisabled orders the results by the is_comment_disabled field.
func ByIsCommentDisabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsCommentDisabled, opts...).ToFunc()
}

// ByVisibility orders the results by the visibility field.
func ByVisibility(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVisibility, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByIsDeleted orders the results by the is_deleted field.
func ByIsDeleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDeleted, opts...).ToFunc()
}

// ByParentThreadField orders the results by parent_thread field.
func ByParentThreadField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentThreadStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildThreadsCount orders the results by child_threads count.
func ByChildThreadsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildThreadsStep(), opts...)
	}
}

// ByChildThreads orders the results by child_threads terms.
func ByChildThreads(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildThreadsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByImagesCount orders the results by images count.
func ByImagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newImagesStep(), opts...)
	}
}

// ByImages orders the results by images terms.
func ByImages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newImagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newParentThreadStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentThreadTable, ParentThreadColumn),
	)
}
func newChildThreadsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildThreadsTable, ChildThreadsColumn),
	)
}
func newImagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ImagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ImagesTable, ImagesPrimaryKey...),
	)
}