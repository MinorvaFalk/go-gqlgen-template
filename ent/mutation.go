// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-gqlgen-template/ent/predicate"
	"go-gqlgen-template/ent/schema/ulid"
	"go-gqlgen-template/ent/todo"
	"go-gqlgen-template/ent/user"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTodo = "Todo"
	TypeUser = "User"
)

// TodoMutation represents an operation that mutates the Todo nodes in the graph.
type TodoMutation struct {
	config
	op            Op
	typ           string
	id            *ulid.ID
	title         *string
	completed     *bool
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	user          *ulid.ID
	cleareduser   bool
	done          bool
	oldValue      func(context.Context) (*Todo, error)
	predicates    []predicate.Todo
}

var _ ent.Mutation = (*TodoMutation)(nil)

// todoOption allows management of the mutation configuration using functional options.
type todoOption func(*TodoMutation)

// newTodoMutation creates new mutation for the Todo entity.
func newTodoMutation(c config, op Op, opts ...todoOption) *TodoMutation {
	m := &TodoMutation{
		config:        c,
		op:            op,
		typ:           TypeTodo,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTodoID sets the ID field of the mutation.
func withTodoID(id ulid.ID) todoOption {
	return func(m *TodoMutation) {
		var (
			err   error
			once  sync.Once
			value *Todo
		)
		m.oldValue = func(ctx context.Context) (*Todo, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Todo.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTodo sets the old Todo of the mutation.
func withTodo(node *Todo) todoOption {
	return func(m *TodoMutation) {
		m.oldValue = func(context.Context) (*Todo, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TodoMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TodoMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Todo entities.
func (m *TodoMutation) SetID(id ulid.ID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TodoMutation) ID() (id ulid.ID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TodoMutation) IDs(ctx context.Context) ([]ulid.ID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []ulid.ID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Todo.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUserID sets the "user_id" field.
func (m *TodoMutation) SetUserID(u ulid.ID) {
	m.user = &u
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *TodoMutation) UserID() (r ulid.ID, exists bool) {
	v := m.user
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldUserID(ctx context.Context) (v ulid.ID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// ClearUserID clears the value of the "user_id" field.
func (m *TodoMutation) ClearUserID() {
	m.user = nil
	m.clearedFields[todo.FieldUserID] = struct{}{}
}

// UserIDCleared returns if the "user_id" field was cleared in this mutation.
func (m *TodoMutation) UserIDCleared() bool {
	_, ok := m.clearedFields[todo.FieldUserID]
	return ok
}

// ResetUserID resets all changes to the "user_id" field.
func (m *TodoMutation) ResetUserID() {
	m.user = nil
	delete(m.clearedFields, todo.FieldUserID)
}

// SetTitle sets the "title" field.
func (m *TodoMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *TodoMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *TodoMutation) ResetTitle() {
	m.title = nil
}

// SetCompleted sets the "completed" field.
func (m *TodoMutation) SetCompleted(b bool) {
	m.completed = &b
}

// Completed returns the value of the "completed" field in the mutation.
func (m *TodoMutation) Completed() (r bool, exists bool) {
	v := m.completed
	if v == nil {
		return
	}
	return *v, true
}

// OldCompleted returns the old "completed" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldCompleted(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCompleted is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCompleted requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCompleted: %w", err)
	}
	return oldValue.Completed, nil
}

// ResetCompleted resets all changes to the "completed" field.
func (m *TodoMutation) ResetCompleted() {
	m.completed = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *TodoMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TodoMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TodoMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TodoMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TodoMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TodoMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// ClearUser clears the "user" edge to the User entity.
func (m *TodoMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared reports if the "user" edge to the User entity was cleared.
func (m *TodoMutation) UserCleared() bool {
	return m.UserIDCleared() || m.cleareduser
}

// UserIDs returns the "user" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// UserID instead. It exists only for internal usage by the builders.
func (m *TodoMutation) UserIDs() (ids []ulid.ID) {
	if id := m.user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *TodoMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
}

// Where appends a list predicates to the TodoMutation builder.
func (m *TodoMutation) Where(ps ...predicate.Todo) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *TodoMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Todo).
func (m *TodoMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TodoMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.user != nil {
		fields = append(fields, todo.FieldUserID)
	}
	if m.title != nil {
		fields = append(fields, todo.FieldTitle)
	}
	if m.completed != nil {
		fields = append(fields, todo.FieldCompleted)
	}
	if m.created_at != nil {
		fields = append(fields, todo.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, todo.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TodoMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case todo.FieldUserID:
		return m.UserID()
	case todo.FieldTitle:
		return m.Title()
	case todo.FieldCompleted:
		return m.Completed()
	case todo.FieldCreatedAt:
		return m.CreatedAt()
	case todo.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TodoMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case todo.FieldUserID:
		return m.OldUserID(ctx)
	case todo.FieldTitle:
		return m.OldTitle(ctx)
	case todo.FieldCompleted:
		return m.OldCompleted(ctx)
	case todo.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case todo.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Todo field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TodoMutation) SetField(name string, value ent.Value) error {
	switch name {
	case todo.FieldUserID:
		v, ok := value.(ulid.ID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case todo.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case todo.FieldCompleted:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCompleted(v)
		return nil
	case todo.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case todo.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Todo field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TodoMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TodoMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TodoMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Todo numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TodoMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(todo.FieldUserID) {
		fields = append(fields, todo.FieldUserID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TodoMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TodoMutation) ClearField(name string) error {
	switch name {
	case todo.FieldUserID:
		m.ClearUserID()
		return nil
	}
	return fmt.Errorf("unknown Todo nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TodoMutation) ResetField(name string) error {
	switch name {
	case todo.FieldUserID:
		m.ResetUserID()
		return nil
	case todo.FieldTitle:
		m.ResetTitle()
		return nil
	case todo.FieldCompleted:
		m.ResetCompleted()
		return nil
	case todo.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case todo.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Todo field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TodoMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, todo.EdgeUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TodoMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case todo.EdgeUser:
		if id := m.user; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TodoMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TodoMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TodoMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, todo.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TodoMutation) EdgeCleared(name string) bool {
	switch name {
	case todo.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TodoMutation) ClearEdge(name string) error {
	switch name {
	case todo.EdgeUser:
		m.ClearUser()
		return nil
	}
	return fmt.Errorf("unknown Todo unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TodoMutation) ResetEdge(name string) error {
	switch name {
	case todo.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown Todo edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *ulid.ID
	name          *string
	username      *string
	email         *string
	phone         *string
	website       *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	todo          map[ulid.ID]struct{}
	removedtodo   map[ulid.ID]struct{}
	clearedtodo   bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id ulid.ID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id ulid.ID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id ulid.ID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]ulid.ID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []ulid.ID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetUsername sets the "username" field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetPhone sets the "phone" field.
func (m *UserMutation) SetPhone(s string) {
	m.phone = &s
}

// Phone returns the value of the "phone" field in the mutation.
func (m *UserMutation) Phone() (r string, exists bool) {
	v := m.phone
	if v == nil {
		return
	}
	return *v, true
}

// OldPhone returns the old "phone" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPhone(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPhone is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPhone requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhone: %w", err)
	}
	return oldValue.Phone, nil
}

// ResetPhone resets all changes to the "phone" field.
func (m *UserMutation) ResetPhone() {
	m.phone = nil
}

// SetWebsite sets the "website" field.
func (m *UserMutation) SetWebsite(s string) {
	m.website = &s
}

// Website returns the value of the "website" field in the mutation.
func (m *UserMutation) Website() (r string, exists bool) {
	v := m.website
	if v == nil {
		return
	}
	return *v, true
}

// OldWebsite returns the old "website" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldWebsite(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldWebsite is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldWebsite requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldWebsite: %w", err)
	}
	return oldValue.Website, nil
}

// ClearWebsite clears the value of the "website" field.
func (m *UserMutation) ClearWebsite() {
	m.website = nil
	m.clearedFields[user.FieldWebsite] = struct{}{}
}

// WebsiteCleared returns if the "website" field was cleared in this mutation.
func (m *UserMutation) WebsiteCleared() bool {
	_, ok := m.clearedFields[user.FieldWebsite]
	return ok
}

// ResetWebsite resets all changes to the "website" field.
func (m *UserMutation) ResetWebsite() {
	m.website = nil
	delete(m.clearedFields, user.FieldWebsite)
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// AddTodoIDs adds the "todo" edge to the Todo entity by ids.
func (m *UserMutation) AddTodoIDs(ids ...ulid.ID) {
	if m.todo == nil {
		m.todo = make(map[ulid.ID]struct{})
	}
	for i := range ids {
		m.todo[ids[i]] = struct{}{}
	}
}

// ClearTodo clears the "todo" edge to the Todo entity.
func (m *UserMutation) ClearTodo() {
	m.clearedtodo = true
}

// TodoCleared reports if the "todo" edge to the Todo entity was cleared.
func (m *UserMutation) TodoCleared() bool {
	return m.clearedtodo
}

// RemoveTodoIDs removes the "todo" edge to the Todo entity by IDs.
func (m *UserMutation) RemoveTodoIDs(ids ...ulid.ID) {
	if m.removedtodo == nil {
		m.removedtodo = make(map[ulid.ID]struct{})
	}
	for i := range ids {
		delete(m.todo, ids[i])
		m.removedtodo[ids[i]] = struct{}{}
	}
}

// RemovedTodo returns the removed IDs of the "todo" edge to the Todo entity.
func (m *UserMutation) RemovedTodoIDs() (ids []ulid.ID) {
	for id := range m.removedtodo {
		ids = append(ids, id)
	}
	return
}

// TodoIDs returns the "todo" edge IDs in the mutation.
func (m *UserMutation) TodoIDs() (ids []ulid.ID) {
	for id := range m.todo {
		ids = append(ids, id)
	}
	return
}

// ResetTodo resets all changes to the "todo" edge.
func (m *UserMutation) ResetTodo() {
	m.todo = nil
	m.clearedtodo = false
	m.removedtodo = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.phone != nil {
		fields = append(fields, user.FieldPhone)
	}
	if m.website != nil {
		fields = append(fields, user.FieldWebsite)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	case user.FieldUsername:
		return m.Username()
	case user.FieldEmail:
		return m.Email()
	case user.FieldPhone:
		return m.Phone()
	case user.FieldWebsite:
		return m.Website()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldUsername:
		return m.OldUsername(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPhone:
		return m.OldPhone(ctx)
	case user.FieldWebsite:
		return m.OldWebsite(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPhone:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhone(v)
		return nil
	case user.FieldWebsite:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetWebsite(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldWebsite) {
		fields = append(fields, user.FieldWebsite)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldWebsite:
		m.ClearWebsite()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPhone:
		m.ResetPhone()
		return nil
	case user.FieldWebsite:
		m.ResetWebsite()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.todo != nil {
		edges = append(edges, user.EdgeTodo)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTodo:
		ids := make([]ent.Value, 0, len(m.todo))
		for id := range m.todo {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtodo != nil {
		edges = append(edges, user.EdgeTodo)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTodo:
		ids := make([]ent.Value, 0, len(m.removedtodo))
		for id := range m.removedtodo {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtodo {
		edges = append(edges, user.EdgeTodo)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeTodo:
		return m.clearedtodo
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeTodo:
		m.ResetTodo()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
