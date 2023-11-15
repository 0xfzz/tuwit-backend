// Code generated by ent, DO NOT EDIT.

package thread

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the thread type in the database.
	Label = "thread"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
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
	// FieldAuthorID holds the string denoting the author_id field in the database.
	FieldAuthorID = "author_id"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldRepostThreadID holds the string denoting the repost_thread_id field in the database.
	FieldRepostThreadID = "repost_thread_id"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeThreadCount holds the string denoting the thread_count edge name in mutations.
	EdgeThreadCount = "thread_count"
	// EdgeReposted holds the string denoting the reposted edge name in mutations.
	EdgeReposted = "reposted"
	// EdgeRepost holds the string denoting the repost edge name in mutations.
	EdgeRepost = "repost"
	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"
	// Table holds the table name of the thread in the database.
	Table = "threads"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "threads"
	// AuthorInverseTable is the table name for the UserAccount entity.
	// It exists in this package in order to avoid circular dependency with the "useraccount" package.
	AuthorInverseTable = "user_account"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "author_id"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "threads"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "threads"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
	// ThreadCountTable is the table that holds the thread_count relation/edge.
	ThreadCountTable = "thread_account"
	// ThreadCountInverseTable is the table name for the ThreadCount entity.
	// It exists in this package in order to avoid circular dependency with the "threadcount" package.
	ThreadCountInverseTable = "thread_account"
	// ThreadCountColumn is the table column denoting the thread_count relation/edge.
	ThreadCountColumn = "thread_thread_count"
	// RepostedTable is the table that holds the reposted relation/edge.
	RepostedTable = "threads"
	// RepostedColumn is the table column denoting the reposted relation/edge.
	RepostedColumn = "repost_thread_id"
	// RepostTable is the table that holds the repost relation/edge.
	RepostTable = "threads"
	// RepostColumn is the table column denoting the repost relation/edge.
	RepostColumn = "repost_thread_id"
	// ImagesTable is the table that holds the images relation/edge. The primary key declared below.
	ImagesTable = "thread_images"
	// ImagesInverseTable is the table name for the Media entity.
	// It exists in this package in order to avoid circular dependency with the "media" package.
	ImagesInverseTable = "media"
)

// Columns holds all SQL columns for thread fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldContent,
	FieldIsCommentDisabled,
	FieldVisibility,
	FieldStatus,
	FieldIsDeleted,
	FieldAuthorID,
	FieldParentID,
	FieldRepostThreadID,
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
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultIsCommentDisabled holds the default value on creation for the "is_comment_disabled" field.
	DefaultIsCommentDisabled bool
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted bool
)

// Visibility defines the type for the "visibility" enum field.
type Visibility string

// VisibilityPUBLIC is the default value of the Visibility enum.
const DefaultVisibility = VisibilityPUBLIC

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

// StatusPUBLISHED is the default value of the Status enum.
const DefaultStatus = StatusPUBLISHED

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

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
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

// ByAuthorID orders the results by the author_id field.
func ByAuthorID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthorID, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByRepostThreadID orders the results by the repost_thread_id field.
func ByRepostThreadID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRepostThreadID, opts...).ToFunc()
}

// ByAuthorField orders the results by author field.
func ByAuthorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), sql.OrderByField(field, opts...))
	}
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByThreadCountCount orders the results by thread_count count.
func ByThreadCountCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newThreadCountStep(), opts...)
	}
}

// ByThreadCount orders the results by thread_count terms.
func ByThreadCount(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newThreadCountStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRepostedField orders the results by reposted field.
func ByRepostedField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRepostedStep(), sql.OrderByField(field, opts...))
	}
}

// ByRepostField orders the results by repost field.
func ByRepostField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRepostStep(), sql.OrderByField(field, opts...))
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
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
	)
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
	)
}
func newThreadCountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ThreadCountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ThreadCountTable, ThreadCountColumn),
	)
}
func newRepostedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, RepostedTable, RepostedColumn),
	)
}
func newRepostStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, RepostTable, RepostColumn),
	)
}
func newImagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ImagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ImagesTable, ImagesPrimaryKey...),
	)
}
