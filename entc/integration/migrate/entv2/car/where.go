// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package car

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv2/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Car {
	return predicate.Car(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Car {
	return predicate.Car(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Car {
	return predicate.Car(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Car {
	return predicate.Car(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Car {
	return predicate.Car(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Car {
	return predicate.Car(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Car {
	return predicate.Car(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldName, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Car {
	return predicate.Car(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Car {
	return predicate.Car(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Car {
	return predicate.Car(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Car {
	return predicate.Car(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Car {
	return predicate.Car(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Car {
	return predicate.Car(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Car {
	return predicate.Car(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Car {
	return predicate.Car(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Car {
	return predicate.Car(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Car {
	return predicate.Car(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Car {
	return predicate.Car(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Car {
	return predicate.Car(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Car {
	return predicate.Car(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Car {
	return predicate.Car(sql.FieldContainsFold(FieldName, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
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
func Not(p predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		p(s.Not())
	})
}
