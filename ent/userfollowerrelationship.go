// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/0xfzz/tuwit-backend/ent/useraccount"
	"github.com/0xfzz/tuwit-backend/ent/userfollowerrelationship"
)

// UserFollowerRelationship is the model entity for the UserFollowerRelationship schema.
type UserFollowerRelationship struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// FollowerID holds the value of the "follower_id" field.
	FollowerID int `json:"follower_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserFollowerRelationshipQuery when eager-loading is set.
	Edges        UserFollowerRelationshipEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserFollowerRelationshipEdges holds the relations/edges for other nodes in the graph.
type UserFollowerRelationshipEdges struct {
	// Follower holds the value of the follower edge.
	Follower *UserAccount `json:"follower,omitempty"`
	// Following holds the value of the following edge.
	Following *UserAccount `json:"following,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FollowerOrErr returns the Follower value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserFollowerRelationshipEdges) FollowerOrErr() (*UserAccount, error) {
	if e.loadedTypes[0] {
		if e.Follower == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: useraccount.Label}
		}
		return e.Follower, nil
	}
	return nil, &NotLoadedError{edge: "follower"}
}

// FollowingOrErr returns the Following value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserFollowerRelationshipEdges) FollowingOrErr() (*UserAccount, error) {
	if e.loadedTypes[1] {
		if e.Following == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: useraccount.Label}
		}
		return e.Following, nil
	}
	return nil, &NotLoadedError{edge: "following"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserFollowerRelationship) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userfollowerrelationship.FieldID, userfollowerrelationship.FieldUserID, userfollowerrelationship.FieldFollowerID:
			values[i] = new(sql.NullInt64)
		case userfollowerrelationship.FieldCreatedAt, userfollowerrelationship.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserFollowerRelationship fields.
func (ufr *UserFollowerRelationship) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userfollowerrelationship.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ufr.ID = int(value.Int64)
		case userfollowerrelationship.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ufr.CreatedAt = value.Time
			}
		case userfollowerrelationship.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ufr.UpdatedAt = value.Time
			}
		case userfollowerrelationship.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ufr.UserID = int(value.Int64)
			}
		case userfollowerrelationship.FieldFollowerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field follower_id", values[i])
			} else if value.Valid {
				ufr.FollowerID = int(value.Int64)
			}
		default:
			ufr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserFollowerRelationship.
// This includes values selected through modifiers, order, etc.
func (ufr *UserFollowerRelationship) Value(name string) (ent.Value, error) {
	return ufr.selectValues.Get(name)
}

// QueryFollower queries the "follower" edge of the UserFollowerRelationship entity.
func (ufr *UserFollowerRelationship) QueryFollower() *UserAccountQuery {
	return NewUserFollowerRelationshipClient(ufr.config).QueryFollower(ufr)
}

// QueryFollowing queries the "following" edge of the UserFollowerRelationship entity.
func (ufr *UserFollowerRelationship) QueryFollowing() *UserAccountQuery {
	return NewUserFollowerRelationshipClient(ufr.config).QueryFollowing(ufr)
}

// Update returns a builder for updating this UserFollowerRelationship.
// Note that you need to call UserFollowerRelationship.Unwrap() before calling this method if this UserFollowerRelationship
// was returned from a transaction, and the transaction was committed or rolled back.
func (ufr *UserFollowerRelationship) Update() *UserFollowerRelationshipUpdateOne {
	return NewUserFollowerRelationshipClient(ufr.config).UpdateOne(ufr)
}

// Unwrap unwraps the UserFollowerRelationship entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ufr *UserFollowerRelationship) Unwrap() *UserFollowerRelationship {
	_tx, ok := ufr.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserFollowerRelationship is not a transactional entity")
	}
	ufr.config.driver = _tx.drv
	return ufr
}

// String implements the fmt.Stringer.
func (ufr *UserFollowerRelationship) String() string {
	var builder strings.Builder
	builder.WriteString("UserFollowerRelationship(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ufr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ufr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ufr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ufr.UserID))
	builder.WriteString(", ")
	builder.WriteString("follower_id=")
	builder.WriteString(fmt.Sprintf("%v", ufr.FollowerID))
	builder.WriteByte(')')
	return builder.String()
}

// UserFollowerRelationships is a parsable slice of UserFollowerRelationship.
type UserFollowerRelationships []*UserFollowerRelationship
