// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0xfzz/tuwitt/ent/media"
	"github.com/0xfzz/tuwitt/ent/predicate"
	"github.com/0xfzz/tuwitt/ent/useraccount"
	"github.com/0xfzz/tuwitt/ent/userprofile"
)

// UserProfileUpdate is the builder for updating UserProfile entities.
type UserProfileUpdate struct {
	config
	hooks    []Hook
	mutation *UserProfileMutation
}

// Where appends a list predicates to the UserProfileUpdate builder.
func (upu *UserProfileUpdate) Where(ps ...predicate.UserProfile) *UserProfileUpdate {
	upu.mutation.Where(ps...)
	return upu
}

// SetDisplayName sets the "display_name" field.
func (upu *UserProfileUpdate) SetDisplayName(s string) *UserProfileUpdate {
	upu.mutation.SetDisplayName(s)
	return upu
}

// SetBio sets the "bio" field.
func (upu *UserProfileUpdate) SetBio(s string) *UserProfileUpdate {
	upu.mutation.SetBio(s)
	return upu
}

// SetAccountID sets the "account" edge to the UserAccount entity by ID.
func (upu *UserProfileUpdate) SetAccountID(id int) *UserProfileUpdate {
	upu.mutation.SetAccountID(id)
	return upu
}

// SetAccount sets the "account" edge to the UserAccount entity.
func (upu *UserProfileUpdate) SetAccount(u *UserAccount) *UserProfileUpdate {
	return upu.SetAccountID(u.ID)
}

// SetProfilePictureID sets the "profile_picture" edge to the Media entity by ID.
func (upu *UserProfileUpdate) SetProfilePictureID(id int) *UserProfileUpdate {
	upu.mutation.SetProfilePictureID(id)
	return upu
}

// SetNillableProfilePictureID sets the "profile_picture" edge to the Media entity by ID if the given value is not nil.
func (upu *UserProfileUpdate) SetNillableProfilePictureID(id *int) *UserProfileUpdate {
	if id != nil {
		upu = upu.SetProfilePictureID(*id)
	}
	return upu
}

// SetProfilePicture sets the "profile_picture" edge to the Media entity.
func (upu *UserProfileUpdate) SetProfilePicture(m *Media) *UserProfileUpdate {
	return upu.SetProfilePictureID(m.ID)
}

// SetBannerID sets the "banner" edge to the Media entity by ID.
func (upu *UserProfileUpdate) SetBannerID(id int) *UserProfileUpdate {
	upu.mutation.SetBannerID(id)
	return upu
}

// SetNillableBannerID sets the "banner" edge to the Media entity by ID if the given value is not nil.
func (upu *UserProfileUpdate) SetNillableBannerID(id *int) *UserProfileUpdate {
	if id != nil {
		upu = upu.SetBannerID(*id)
	}
	return upu
}

// SetBanner sets the "banner" edge to the Media entity.
func (upu *UserProfileUpdate) SetBanner(m *Media) *UserProfileUpdate {
	return upu.SetBannerID(m.ID)
}

// Mutation returns the UserProfileMutation object of the builder.
func (upu *UserProfileUpdate) Mutation() *UserProfileMutation {
	return upu.mutation
}

// ClearAccount clears the "account" edge to the UserAccount entity.
func (upu *UserProfileUpdate) ClearAccount() *UserProfileUpdate {
	upu.mutation.ClearAccount()
	return upu
}

// ClearProfilePicture clears the "profile_picture" edge to the Media entity.
func (upu *UserProfileUpdate) ClearProfilePicture() *UserProfileUpdate {
	upu.mutation.ClearProfilePicture()
	return upu
}

// ClearBanner clears the "banner" edge to the Media entity.
func (upu *UserProfileUpdate) ClearBanner() *UserProfileUpdate {
	upu.mutation.ClearBanner()
	return upu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (upu *UserProfileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, upu.sqlSave, upu.mutation, upu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (upu *UserProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := upu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (upu *UserProfileUpdate) Exec(ctx context.Context) error {
	_, err := upu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upu *UserProfileUpdate) ExecX(ctx context.Context) {
	if err := upu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upu *UserProfileUpdate) check() error {
	if _, ok := upu.mutation.AccountID(); upu.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserProfile.account"`)
	}
	return nil
}

func (upu *UserProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := upu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userprofile.Table, userprofile.Columns, sqlgraph.NewFieldSpec(userprofile.FieldID, field.TypeInt))
	if ps := upu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upu.mutation.DisplayName(); ok {
		_spec.SetField(userprofile.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := upu.mutation.Bio(); ok {
		_spec.SetField(userprofile.FieldBio, field.TypeString, value)
	}
	if upu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userprofile.AccountTable,
			Columns: []string{userprofile.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(useraccount.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := upu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userprofile.AccountTable,
			Columns: []string{userprofile.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(useraccount.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if upu.mutation.ProfilePictureCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userprofile.ProfilePictureTable,
			Columns: []string{userprofile.ProfilePictureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := upu.mutation.ProfilePictureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userprofile.ProfilePictureTable,
			Columns: []string{userprofile.ProfilePictureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if upu.mutation.BannerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userprofile.BannerTable,
			Columns: []string{userprofile.BannerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := upu.mutation.BannerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userprofile.BannerTable,
			Columns: []string{userprofile.BannerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, upu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userprofile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	upu.mutation.done = true
	return n, nil
}

// UserProfileUpdateOne is the builder for updating a single UserProfile entity.
type UserProfileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserProfileMutation
}

// SetDisplayName sets the "display_name" field.
func (upuo *UserProfileUpdateOne) SetDisplayName(s string) *UserProfileUpdateOne {
	upuo.mutation.SetDisplayName(s)
	return upuo
}

// SetBio sets the "bio" field.
func (upuo *UserProfileUpdateOne) SetBio(s string) *UserProfileUpdateOne {
	upuo.mutation.SetBio(s)
	return upuo
}

// SetAccountID sets the "account" edge to the UserAccount entity by ID.
func (upuo *UserProfileUpdateOne) SetAccountID(id int) *UserProfileUpdateOne {
	upuo.mutation.SetAccountID(id)
	return upuo
}

// SetAccount sets the "account" edge to the UserAccount entity.
func (upuo *UserProfileUpdateOne) SetAccount(u *UserAccount) *UserProfileUpdateOne {
	return upuo.SetAccountID(u.ID)
}

// SetProfilePictureID sets the "profile_picture" edge to the Media entity by ID.
func (upuo *UserProfileUpdateOne) SetProfilePictureID(id int) *UserProfileUpdateOne {
	upuo.mutation.SetProfilePictureID(id)
	return upuo
}

// SetNillableProfilePictureID sets the "profile_picture" edge to the Media entity by ID if the given value is not nil.
func (upuo *UserProfileUpdateOne) SetNillableProfilePictureID(id *int) *UserProfileUpdateOne {
	if id != nil {
		upuo = upuo.SetProfilePictureID(*id)
	}
	return upuo
}

// SetProfilePicture sets the "profile_picture" edge to the Media entity.
func (upuo *UserProfileUpdateOne) SetProfilePicture(m *Media) *UserProfileUpdateOne {
	return upuo.SetProfilePictureID(m.ID)
}

// SetBannerID sets the "banner" edge to the Media entity by ID.
func (upuo *UserProfileUpdateOne) SetBannerID(id int) *UserProfileUpdateOne {
	upuo.mutation.SetBannerID(id)
	return upuo
}

// SetNillableBannerID sets the "banner" edge to the Media entity by ID if the given value is not nil.
func (upuo *UserProfileUpdateOne) SetNillableBannerID(id *int) *UserProfileUpdateOne {
	if id != nil {
		upuo = upuo.SetBannerID(*id)
	}
	return upuo
}

// SetBanner sets the "banner" edge to the Media entity.
func (upuo *UserProfileUpdateOne) SetBanner(m *Media) *UserProfileUpdateOne {
	return upuo.SetBannerID(m.ID)
}

// Mutation returns the UserProfileMutation object of the builder.
func (upuo *UserProfileUpdateOne) Mutation() *UserProfileMutation {
	return upuo.mutation
}

// ClearAccount clears the "account" edge to the UserAccount entity.
func (upuo *UserProfileUpdateOne) ClearAccount() *UserProfileUpdateOne {
	upuo.mutation.ClearAccount()
	return upuo
}

// ClearProfilePicture clears the "profile_picture" edge to the Media entity.
func (upuo *UserProfileUpdateOne) ClearProfilePicture() *UserProfileUpdateOne {
	upuo.mutation.ClearProfilePicture()
	return upuo
}

// ClearBanner clears the "banner" edge to the Media entity.
func (upuo *UserProfileUpdateOne) ClearBanner() *UserProfileUpdateOne {
	upuo.mutation.ClearBanner()
	return upuo
}

// Where appends a list predicates to the UserProfileUpdate builder.
func (upuo *UserProfileUpdateOne) Where(ps ...predicate.UserProfile) *UserProfileUpdateOne {
	upuo.mutation.Where(ps...)
	return upuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (upuo *UserProfileUpdateOne) Select(field string, fields ...string) *UserProfileUpdateOne {
	upuo.fields = append([]string{field}, fields...)
	return upuo
}

// Save executes the query and returns the updated UserProfile entity.
func (upuo *UserProfileUpdateOne) Save(ctx context.Context) (*UserProfile, error) {
	return withHooks(ctx, upuo.sqlSave, upuo.mutation, upuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (upuo *UserProfileUpdateOne) SaveX(ctx context.Context) *UserProfile {
	node, err := upuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (upuo *UserProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := upuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upuo *UserProfileUpdateOne) ExecX(ctx context.Context) {
	if err := upuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upuo *UserProfileUpdateOne) check() error {
	if _, ok := upuo.mutation.AccountID(); upuo.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserProfile.account"`)
	}
	return nil
}

func (upuo *UserProfileUpdateOne) sqlSave(ctx context.Context) (_node *UserProfile, err error) {
	if err := upuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userprofile.Table, userprofile.Columns, sqlgraph.NewFieldSpec(userprofile.FieldID, field.TypeInt))
	id, ok := upuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserProfile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := upuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userprofile.FieldID)
		for _, f := range fields {
			if !userprofile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userprofile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := upuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upuo.mutation.DisplayName(); ok {
		_spec.SetField(userprofile.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := upuo.mutation.Bio(); ok {
		_spec.SetField(userprofile.FieldBio, field.TypeString, value)
	}
	if upuo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userprofile.AccountTable,
			Columns: []string{userprofile.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(useraccount.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := upuo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userprofile.AccountTable,
			Columns: []string{userprofile.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(useraccount.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if upuo.mutation.ProfilePictureCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userprofile.ProfilePictureTable,
			Columns: []string{userprofile.ProfilePictureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := upuo.mutation.ProfilePictureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userprofile.ProfilePictureTable,
			Columns: []string{userprofile.ProfilePictureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if upuo.mutation.BannerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userprofile.BannerTable,
			Columns: []string{userprofile.BannerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := upuo.mutation.BannerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userprofile.BannerTable,
			Columns: []string{userprofile.BannerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserProfile{config: upuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, upuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userprofile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	upuo.mutation.done = true
	return _node, nil
}
