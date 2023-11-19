// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/0xfzz/tuwit-backend/ent/media"
	"github.com/0xfzz/tuwit-backend/ent/useraccount"
	"github.com/0xfzz/tuwit-backend/ent/userprofile"
)

// UserProfile is the model entity for the UserProfile schema.
type UserProfile struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// Bio holds the value of the "bio" field.
	Bio string `json:"bio,omitempty"`
	// ProfilePictureID holds the value of the "profile_picture_id" field.
	ProfilePictureID int `json:"profile_picture_id,omitempty"`
	// BannerID holds the value of the "banner_id" field.
	BannerID int `json:"banner_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserProfileQuery when eager-loading is set.
	Edges                UserProfileEdges `json:"edges"`
	user_account_profile *int
	selectValues         sql.SelectValues
}

// UserProfileEdges holds the relations/edges for other nodes in the graph.
type UserProfileEdges struct {
	// Account holds the value of the account edge.
	Account *UserAccount `json:"account,omitempty"`
	// ProfilePicture holds the value of the profile_picture edge.
	ProfilePicture *Media `json:"profile_picture,omitempty"`
	// Banner holds the value of the banner edge.
	Banner *Media `json:"banner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserProfileEdges) AccountOrErr() (*UserAccount, error) {
	if e.loadedTypes[0] {
		if e.Account == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: useraccount.Label}
		}
		return e.Account, nil
	}
	return nil, &NotLoadedError{edge: "account"}
}

// ProfilePictureOrErr returns the ProfilePicture value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserProfileEdges) ProfilePictureOrErr() (*Media, error) {
	if e.loadedTypes[1] {
		if e.ProfilePicture == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: media.Label}
		}
		return e.ProfilePicture, nil
	}
	return nil, &NotLoadedError{edge: "profile_picture"}
}

// BannerOrErr returns the Banner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserProfileEdges) BannerOrErr() (*Media, error) {
	if e.loadedTypes[2] {
		if e.Banner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: media.Label}
		}
		return e.Banner, nil
	}
	return nil, &NotLoadedError{edge: "banner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserProfile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userprofile.FieldID, userprofile.FieldProfilePictureID, userprofile.FieldBannerID:
			values[i] = new(sql.NullInt64)
		case userprofile.FieldDisplayName, userprofile.FieldBio:
			values[i] = new(sql.NullString)
		case userprofile.ForeignKeys[0]: // user_account_profile
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserProfile fields.
func (up *UserProfile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userprofile.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			up.ID = int(value.Int64)
		case userprofile.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				up.DisplayName = value.String
			}
		case userprofile.FieldBio:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bio", values[i])
			} else if value.Valid {
				up.Bio = value.String
			}
		case userprofile.FieldProfilePictureID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field profile_picture_id", values[i])
			} else if value.Valid {
				up.ProfilePictureID = int(value.Int64)
			}
		case userprofile.FieldBannerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field banner_id", values[i])
			} else if value.Valid {
				up.BannerID = int(value.Int64)
			}
		case userprofile.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_account_profile", value)
			} else if value.Valid {
				up.user_account_profile = new(int)
				*up.user_account_profile = int(value.Int64)
			}
		default:
			up.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserProfile.
// This includes values selected through modifiers, order, etc.
func (up *UserProfile) Value(name string) (ent.Value, error) {
	return up.selectValues.Get(name)
}

// QueryAccount queries the "account" edge of the UserProfile entity.
func (up *UserProfile) QueryAccount() *UserAccountQuery {
	return NewUserProfileClient(up.config).QueryAccount(up)
}

// QueryProfilePicture queries the "profile_picture" edge of the UserProfile entity.
func (up *UserProfile) QueryProfilePicture() *MediaQuery {
	return NewUserProfileClient(up.config).QueryProfilePicture(up)
}

// QueryBanner queries the "banner" edge of the UserProfile entity.
func (up *UserProfile) QueryBanner() *MediaQuery {
	return NewUserProfileClient(up.config).QueryBanner(up)
}

// Update returns a builder for updating this UserProfile.
// Note that you need to call UserProfile.Unwrap() before calling this method if this UserProfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (up *UserProfile) Update() *UserProfileUpdateOne {
	return NewUserProfileClient(up.config).UpdateOne(up)
}

// Unwrap unwraps the UserProfile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (up *UserProfile) Unwrap() *UserProfile {
	_tx, ok := up.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserProfile is not a transactional entity")
	}
	up.config.driver = _tx.drv
	return up
}

// String implements the fmt.Stringer.
func (up *UserProfile) String() string {
	var builder strings.Builder
	builder.WriteString("UserProfile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", up.ID))
	builder.WriteString("display_name=")
	builder.WriteString(up.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("bio=")
	builder.WriteString(up.Bio)
	builder.WriteString(", ")
	builder.WriteString("profile_picture_id=")
	builder.WriteString(fmt.Sprintf("%v", up.ProfilePictureID))
	builder.WriteString(", ")
	builder.WriteString("banner_id=")
	builder.WriteString(fmt.Sprintf("%v", up.BannerID))
	builder.WriteByte(')')
	return builder.String()
}

// UserProfiles is a parsable slice of UserProfile.
type UserProfiles []*UserProfile
