// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/predicate"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/team"
)

// TeamDelete is the builder for deleting a Team model.
type TeamDelete struct {
	config
	hooks    []Hook
	mutation *TeamMutation
}

// Where adds a new predicate to the TeamDelete builder.
func (td *TeamDelete) Where(ps ...predicate.Team) *TeamDelete {
	td.mutation.predicates = append(td.mutation.predicates, ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TeamDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(td.hooks) == 0 {
		affected, err = td.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeamMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			td.mutation = mutation
			affected, err = td.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(td.hooks) - 1; i >= 0; i-- {
			mut = td.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, td.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TeamDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TeamDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: team.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: team.FieldID,
			},
		},
	}
	if ps := td.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, td.driver, _spec)
}

// TeamDeleteOne is the builder for deleting a single Team model.
type TeamDeleteOne struct {
	td *TeamDelete
}

// Exec executes the deletion query.
func (tdo *TeamDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{team.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TeamDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}
