// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/adshao/ordinals-indexer/internal/data/ent/collection"
	"github.com/adshao/ordinals-indexer/internal/data/ent/predicate"
	"github.com/adshao/ordinals-indexer/internal/data/ent/token"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// CollectionUpdate is the builder for updating Collection entities.
type CollectionUpdate struct {
	config
	hooks    []Hook
	mutation *CollectionMutation
}

// Where appends a list predicates to the CollectionUpdate builder.
func (cu *CollectionUpdate) Where(ps ...predicate.Collection) *CollectionUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CollectionUpdate) SetUpdatedAt(t time.Time) *CollectionUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetTick sets the "tick" field.
func (cu *CollectionUpdate) SetTick(s string) *CollectionUpdate {
	cu.mutation.SetTick(s)
	return cu
}

// SetP sets the "p" field.
func (cu *CollectionUpdate) SetP(s string) *CollectionUpdate {
	cu.mutation.SetP(s)
	return cu
}

// SetNillableP sets the "p" field if the given value is not nil.
func (cu *CollectionUpdate) SetNillableP(s *string) *CollectionUpdate {
	if s != nil {
		cu.SetP(*s)
	}
	return cu
}

// SetMax sets the "max" field.
func (cu *CollectionUpdate) SetMax(u uint64) *CollectionUpdate {
	cu.mutation.ResetMax()
	cu.mutation.SetMax(u)
	return cu
}

// AddMax adds u to the "max" field.
func (cu *CollectionUpdate) AddMax(u int64) *CollectionUpdate {
	cu.mutation.AddMax(u)
	return cu
}

// SetSupply sets the "supply" field.
func (cu *CollectionUpdate) SetSupply(u uint64) *CollectionUpdate {
	cu.mutation.ResetSupply()
	cu.mutation.SetSupply(u)
	return cu
}

// AddSupply adds u to the "supply" field.
func (cu *CollectionUpdate) AddSupply(u int64) *CollectionUpdate {
	cu.mutation.AddSupply(u)
	return cu
}

// SetBaseURI sets the "base_uri" field.
func (cu *CollectionUpdate) SetBaseURI(s string) *CollectionUpdate {
	cu.mutation.SetBaseURI(s)
	return cu
}

// SetName sets the "name" field.
func (cu *CollectionUpdate) SetName(s string) *CollectionUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetDescription sets the "description" field.
func (cu *CollectionUpdate) SetDescription(s string) *CollectionUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetImage sets the "image" field.
func (cu *CollectionUpdate) SetImage(s string) *CollectionUpdate {
	cu.mutation.SetImage(s)
	return cu
}

// SetAttributes sets the "attributes" field.
func (cu *CollectionUpdate) SetAttributes(m []map[string]interface{}) *CollectionUpdate {
	cu.mutation.SetAttributes(m)
	return cu
}

// AppendAttributes appends m to the "attributes" field.
func (cu *CollectionUpdate) AppendAttributes(m []map[string]interface{}) *CollectionUpdate {
	cu.mutation.AppendAttributes(m)
	return cu
}

// SetTxHash sets the "tx_hash" field.
func (cu *CollectionUpdate) SetTxHash(s string) *CollectionUpdate {
	cu.mutation.SetTxHash(s)
	return cu
}

// SetBlockHeight sets the "block_height" field.
func (cu *CollectionUpdate) SetBlockHeight(u uint64) *CollectionUpdate {
	cu.mutation.ResetBlockHeight()
	cu.mutation.SetBlockHeight(u)
	return cu
}

// AddBlockHeight adds u to the "block_height" field.
func (cu *CollectionUpdate) AddBlockHeight(u int64) *CollectionUpdate {
	cu.mutation.AddBlockHeight(u)
	return cu
}

// SetBlockTime sets the "block_time" field.
func (cu *CollectionUpdate) SetBlockTime(t time.Time) *CollectionUpdate {
	cu.mutation.SetBlockTime(t)
	return cu
}

// SetAddress sets the "address" field.
func (cu *CollectionUpdate) SetAddress(s string) *CollectionUpdate {
	cu.mutation.SetAddress(s)
	return cu
}

// SetInscriptionID sets the "inscription_id" field.
func (cu *CollectionUpdate) SetInscriptionID(i int64) *CollectionUpdate {
	cu.mutation.ResetInscriptionID()
	cu.mutation.SetInscriptionID(i)
	return cu
}

// AddInscriptionID adds i to the "inscription_id" field.
func (cu *CollectionUpdate) AddInscriptionID(i int64) *CollectionUpdate {
	cu.mutation.AddInscriptionID(i)
	return cu
}

// SetInscriptionUID sets the "inscription_uid" field.
func (cu *CollectionUpdate) SetInscriptionUID(s string) *CollectionUpdate {
	cu.mutation.SetInscriptionUID(s)
	return cu
}

// AddTokenIDs adds the "tokens" edge to the Token entity by IDs.
func (cu *CollectionUpdate) AddTokenIDs(ids ...int) *CollectionUpdate {
	cu.mutation.AddTokenIDs(ids...)
	return cu
}

// AddTokens adds the "tokens" edges to the Token entity.
func (cu *CollectionUpdate) AddTokens(t ...*Token) *CollectionUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddTokenIDs(ids...)
}

// Mutation returns the CollectionMutation object of the builder.
func (cu *CollectionUpdate) Mutation() *CollectionMutation {
	return cu.mutation
}

// ClearTokens clears all "tokens" edges to the Token entity.
func (cu *CollectionUpdate) ClearTokens() *CollectionUpdate {
	cu.mutation.ClearTokens()
	return cu
}

// RemoveTokenIDs removes the "tokens" edge to Token entities by IDs.
func (cu *CollectionUpdate) RemoveTokenIDs(ids ...int) *CollectionUpdate {
	cu.mutation.RemoveTokenIDs(ids...)
	return cu
}

// RemoveTokens removes "tokens" edges to Token entities.
func (cu *CollectionUpdate) RemoveTokens(t ...*Token) *CollectionUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveTokenIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CollectionUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CollectionUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CollectionUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CollectionUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CollectionUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := collection.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *CollectionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(collection.Table, collection.Columns, sqlgraph.NewFieldSpec(collection.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(collection.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Tick(); ok {
		_spec.SetField(collection.FieldTick, field.TypeString, value)
	}
	if value, ok := cu.mutation.P(); ok {
		_spec.SetField(collection.FieldP, field.TypeString, value)
	}
	if value, ok := cu.mutation.Max(); ok {
		_spec.SetField(collection.FieldMax, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.AddedMax(); ok {
		_spec.AddField(collection.FieldMax, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.Supply(); ok {
		_spec.SetField(collection.FieldSupply, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.AddedSupply(); ok {
		_spec.AddField(collection.FieldSupply, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.BaseURI(); ok {
		_spec.SetField(collection.FieldBaseURI, field.TypeString, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(collection.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(collection.FieldDescription, field.TypeString, value)
	}
	if value, ok := cu.mutation.Image(); ok {
		_spec.SetField(collection.FieldImage, field.TypeString, value)
	}
	if value, ok := cu.mutation.Attributes(); ok {
		_spec.SetField(collection.FieldAttributes, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedAttributes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, collection.FieldAttributes, value)
		})
	}
	if value, ok := cu.mutation.TxHash(); ok {
		_spec.SetField(collection.FieldTxHash, field.TypeString, value)
	}
	if value, ok := cu.mutation.BlockHeight(); ok {
		_spec.SetField(collection.FieldBlockHeight, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.AddedBlockHeight(); ok {
		_spec.AddField(collection.FieldBlockHeight, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.BlockTime(); ok {
		_spec.SetField(collection.FieldBlockTime, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Address(); ok {
		_spec.SetField(collection.FieldAddress, field.TypeString, value)
	}
	if value, ok := cu.mutation.InscriptionID(); ok {
		_spec.SetField(collection.FieldInscriptionID, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedInscriptionID(); ok {
		_spec.AddField(collection.FieldInscriptionID, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.InscriptionUID(); ok {
		_spec.SetField(collection.FieldInscriptionUID, field.TypeString, value)
	}
	if cu.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   collection.TokensTable,
			Columns: []string{collection.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedTokensIDs(); len(nodes) > 0 && !cu.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   collection.TokensTable,
			Columns: []string{collection.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   collection.TokensTable,
			Columns: []string{collection.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{collection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CollectionUpdateOne is the builder for updating a single Collection entity.
type CollectionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CollectionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CollectionUpdateOne) SetUpdatedAt(t time.Time) *CollectionUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetTick sets the "tick" field.
func (cuo *CollectionUpdateOne) SetTick(s string) *CollectionUpdateOne {
	cuo.mutation.SetTick(s)
	return cuo
}

// SetP sets the "p" field.
func (cuo *CollectionUpdateOne) SetP(s string) *CollectionUpdateOne {
	cuo.mutation.SetP(s)
	return cuo
}

// SetNillableP sets the "p" field if the given value is not nil.
func (cuo *CollectionUpdateOne) SetNillableP(s *string) *CollectionUpdateOne {
	if s != nil {
		cuo.SetP(*s)
	}
	return cuo
}

// SetMax sets the "max" field.
func (cuo *CollectionUpdateOne) SetMax(u uint64) *CollectionUpdateOne {
	cuo.mutation.ResetMax()
	cuo.mutation.SetMax(u)
	return cuo
}

// AddMax adds u to the "max" field.
func (cuo *CollectionUpdateOne) AddMax(u int64) *CollectionUpdateOne {
	cuo.mutation.AddMax(u)
	return cuo
}

// SetSupply sets the "supply" field.
func (cuo *CollectionUpdateOne) SetSupply(u uint64) *CollectionUpdateOne {
	cuo.mutation.ResetSupply()
	cuo.mutation.SetSupply(u)
	return cuo
}

// AddSupply adds u to the "supply" field.
func (cuo *CollectionUpdateOne) AddSupply(u int64) *CollectionUpdateOne {
	cuo.mutation.AddSupply(u)
	return cuo
}

// SetBaseURI sets the "base_uri" field.
func (cuo *CollectionUpdateOne) SetBaseURI(s string) *CollectionUpdateOne {
	cuo.mutation.SetBaseURI(s)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CollectionUpdateOne) SetName(s string) *CollectionUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CollectionUpdateOne) SetDescription(s string) *CollectionUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetImage sets the "image" field.
func (cuo *CollectionUpdateOne) SetImage(s string) *CollectionUpdateOne {
	cuo.mutation.SetImage(s)
	return cuo
}

// SetAttributes sets the "attributes" field.
func (cuo *CollectionUpdateOne) SetAttributes(m []map[string]interface{}) *CollectionUpdateOne {
	cuo.mutation.SetAttributes(m)
	return cuo
}

// AppendAttributes appends m to the "attributes" field.
func (cuo *CollectionUpdateOne) AppendAttributes(m []map[string]interface{}) *CollectionUpdateOne {
	cuo.mutation.AppendAttributes(m)
	return cuo
}

// SetTxHash sets the "tx_hash" field.
func (cuo *CollectionUpdateOne) SetTxHash(s string) *CollectionUpdateOne {
	cuo.mutation.SetTxHash(s)
	return cuo
}

// SetBlockHeight sets the "block_height" field.
func (cuo *CollectionUpdateOne) SetBlockHeight(u uint64) *CollectionUpdateOne {
	cuo.mutation.ResetBlockHeight()
	cuo.mutation.SetBlockHeight(u)
	return cuo
}

// AddBlockHeight adds u to the "block_height" field.
func (cuo *CollectionUpdateOne) AddBlockHeight(u int64) *CollectionUpdateOne {
	cuo.mutation.AddBlockHeight(u)
	return cuo
}

// SetBlockTime sets the "block_time" field.
func (cuo *CollectionUpdateOne) SetBlockTime(t time.Time) *CollectionUpdateOne {
	cuo.mutation.SetBlockTime(t)
	return cuo
}

// SetAddress sets the "address" field.
func (cuo *CollectionUpdateOne) SetAddress(s string) *CollectionUpdateOne {
	cuo.mutation.SetAddress(s)
	return cuo
}

// SetInscriptionID sets the "inscription_id" field.
func (cuo *CollectionUpdateOne) SetInscriptionID(i int64) *CollectionUpdateOne {
	cuo.mutation.ResetInscriptionID()
	cuo.mutation.SetInscriptionID(i)
	return cuo
}

// AddInscriptionID adds i to the "inscription_id" field.
func (cuo *CollectionUpdateOne) AddInscriptionID(i int64) *CollectionUpdateOne {
	cuo.mutation.AddInscriptionID(i)
	return cuo
}

// SetInscriptionUID sets the "inscription_uid" field.
func (cuo *CollectionUpdateOne) SetInscriptionUID(s string) *CollectionUpdateOne {
	cuo.mutation.SetInscriptionUID(s)
	return cuo
}

// AddTokenIDs adds the "tokens" edge to the Token entity by IDs.
func (cuo *CollectionUpdateOne) AddTokenIDs(ids ...int) *CollectionUpdateOne {
	cuo.mutation.AddTokenIDs(ids...)
	return cuo
}

// AddTokens adds the "tokens" edges to the Token entity.
func (cuo *CollectionUpdateOne) AddTokens(t ...*Token) *CollectionUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddTokenIDs(ids...)
}

// Mutation returns the CollectionMutation object of the builder.
func (cuo *CollectionUpdateOne) Mutation() *CollectionMutation {
	return cuo.mutation
}

// ClearTokens clears all "tokens" edges to the Token entity.
func (cuo *CollectionUpdateOne) ClearTokens() *CollectionUpdateOne {
	cuo.mutation.ClearTokens()
	return cuo
}

// RemoveTokenIDs removes the "tokens" edge to Token entities by IDs.
func (cuo *CollectionUpdateOne) RemoveTokenIDs(ids ...int) *CollectionUpdateOne {
	cuo.mutation.RemoveTokenIDs(ids...)
	return cuo
}

// RemoveTokens removes "tokens" edges to Token entities.
func (cuo *CollectionUpdateOne) RemoveTokens(t ...*Token) *CollectionUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveTokenIDs(ids...)
}

// Where appends a list predicates to the CollectionUpdate builder.
func (cuo *CollectionUpdateOne) Where(ps ...predicate.Collection) *CollectionUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CollectionUpdateOne) Select(field string, fields ...string) *CollectionUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Collection entity.
func (cuo *CollectionUpdateOne) Save(ctx context.Context) (*Collection, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CollectionUpdateOne) SaveX(ctx context.Context) *Collection {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CollectionUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CollectionUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CollectionUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := collection.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *CollectionUpdateOne) sqlSave(ctx context.Context) (_node *Collection, err error) {
	_spec := sqlgraph.NewUpdateSpec(collection.Table, collection.Columns, sqlgraph.NewFieldSpec(collection.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Collection.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, collection.FieldID)
		for _, f := range fields {
			if !collection.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != collection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(collection.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Tick(); ok {
		_spec.SetField(collection.FieldTick, field.TypeString, value)
	}
	if value, ok := cuo.mutation.P(); ok {
		_spec.SetField(collection.FieldP, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Max(); ok {
		_spec.SetField(collection.FieldMax, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.AddedMax(); ok {
		_spec.AddField(collection.FieldMax, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.Supply(); ok {
		_spec.SetField(collection.FieldSupply, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.AddedSupply(); ok {
		_spec.AddField(collection.FieldSupply, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.BaseURI(); ok {
		_spec.SetField(collection.FieldBaseURI, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(collection.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(collection.FieldDescription, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Image(); ok {
		_spec.SetField(collection.FieldImage, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Attributes(); ok {
		_spec.SetField(collection.FieldAttributes, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedAttributes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, collection.FieldAttributes, value)
		})
	}
	if value, ok := cuo.mutation.TxHash(); ok {
		_spec.SetField(collection.FieldTxHash, field.TypeString, value)
	}
	if value, ok := cuo.mutation.BlockHeight(); ok {
		_spec.SetField(collection.FieldBlockHeight, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.AddedBlockHeight(); ok {
		_spec.AddField(collection.FieldBlockHeight, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.BlockTime(); ok {
		_spec.SetField(collection.FieldBlockTime, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Address(); ok {
		_spec.SetField(collection.FieldAddress, field.TypeString, value)
	}
	if value, ok := cuo.mutation.InscriptionID(); ok {
		_spec.SetField(collection.FieldInscriptionID, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedInscriptionID(); ok {
		_spec.AddField(collection.FieldInscriptionID, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.InscriptionUID(); ok {
		_spec.SetField(collection.FieldInscriptionUID, field.TypeString, value)
	}
	if cuo.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   collection.TokensTable,
			Columns: []string{collection.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedTokensIDs(); len(nodes) > 0 && !cuo.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   collection.TokensTable,
			Columns: []string{collection.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   collection.TokensTable,
			Columns: []string{collection.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Collection{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{collection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
