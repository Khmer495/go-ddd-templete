// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/auth"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/user"
)

// AuthCreate is the builder for creating a Auth model.
type AuthCreate struct {
	config
	mutation *AuthMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (ac *AuthCreate) SetUserID(i int) *AuthCreate {
	ac.mutation.SetUserID(i)
	return ac
}

// SetFirebaseUserID sets the "firebase_user_id" field.
func (ac *AuthCreate) SetFirebaseUserID(s string) *AuthCreate {
	ac.mutation.SetFirebaseUserID(s)
	return ac
}

// SetUser sets the "user" edge to the User model.
func (ac *AuthCreate) SetUser(u *User) *AuthCreate {
	return ac.SetUserID(u.ID)
}

// Mutation returns the AuthMutation object of the builder.
func (ac *AuthCreate) Mutation() *AuthMutation {
	return ac.mutation
}

// Save creates the Auth in the database.
func (ac *AuthCreate) Save(ctx context.Context) (*Auth, error) {
	var (
		err  error
		node *Auth
	)
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AuthCreate) SaveX(ctx context.Context) *Auth {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (ac *AuthCreate) check() error {
	if _, ok := ac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New("ent: missing required field \"user_id\"")}
	}
	if _, ok := ac.mutation.FirebaseUserID(); !ok {
		return &ValidationError{Name: "firebase_user_id", err: errors.New("ent: missing required field \"firebase_user_id\"")}
	}
	if _, ok := ac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	return nil
}

func (ac *AuthCreate) sqlSave(ctx context.Context) (*Auth, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AuthCreate) createSpec() (*Auth, *sqlgraph.CreateSpec) {
	var (
		_node = &Auth{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: auth.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: auth.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.FirebaseUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: auth.FieldFirebaseUserID,
		})
		_node.FirebaseUserID = value
	}
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   auth.UserTable,
			Columns: []string{auth.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AuthCreateBulk is the builder for creating many Auth entities in bulk.
type AuthCreateBulk struct {
	config
	builders []*AuthCreate
}

// Save creates the Auth entities in the database.
func (acb *AuthCreateBulk) Save(ctx context.Context) ([]*Auth, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Auth, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AuthCreateBulk) SaveX(ctx context.Context) []*Auth {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
