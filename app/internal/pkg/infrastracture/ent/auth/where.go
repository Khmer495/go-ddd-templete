// Code generated by entc, DO NOT EDIT.

package auth

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// FirebaseUserID applies equality check predicate on the "firebase_user_id" field. It's identical to FirebaseUserIDEQ.
func FirebaseUserID(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirebaseUserID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// FirebaseUserIDEQ applies the EQ predicate on the "firebase_user_id" field.
func FirebaseUserIDEQ(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirebaseUserID), v))
	})
}

// FirebaseUserIDNEQ applies the NEQ predicate on the "firebase_user_id" field.
func FirebaseUserIDNEQ(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFirebaseUserID), v))
	})
}

// FirebaseUserIDIn applies the In predicate on the "firebase_user_id" field.
func FirebaseUserIDIn(vs ...string) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFirebaseUserID), v...))
	})
}

// FirebaseUserIDNotIn applies the NotIn predicate on the "firebase_user_id" field.
func FirebaseUserIDNotIn(vs ...string) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFirebaseUserID), v...))
	})
}

// FirebaseUserIDGT applies the GT predicate on the "firebase_user_id" field.
func FirebaseUserIDGT(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFirebaseUserID), v))
	})
}

// FirebaseUserIDGTE applies the GTE predicate on the "firebase_user_id" field.
func FirebaseUserIDGTE(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFirebaseUserID), v))
	})
}

// FirebaseUserIDLT applies the LT predicate on the "firebase_user_id" field.
func FirebaseUserIDLT(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFirebaseUserID), v))
	})
}

// FirebaseUserIDLTE applies the LTE predicate on the "firebase_user_id" field.
func FirebaseUserIDLTE(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFirebaseUserID), v))
	})
}

// FirebaseUserIDContains applies the Contains predicate on the "firebase_user_id" field.
func FirebaseUserIDContains(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFirebaseUserID), v))
	})
}

// FirebaseUserIDHasPrefix applies the HasPrefix predicate on the "firebase_user_id" field.
func FirebaseUserIDHasPrefix(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFirebaseUserID), v))
	})
}

// FirebaseUserIDHasSuffix applies the HasSuffix predicate on the "firebase_user_id" field.
func FirebaseUserIDHasSuffix(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFirebaseUserID), v))
	})
}

// FirebaseUserIDEqualFold applies the EqualFold predicate on the "firebase_user_id" field.
func FirebaseUserIDEqualFold(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFirebaseUserID), v))
	})
}

// FirebaseUserIDContainsFold applies the ContainsFold predicate on the "firebase_user_id" field.
func FirebaseUserIDContainsFold(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFirebaseUserID), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Auth) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Auth) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Auth) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		p(s.Not())
	})
}
