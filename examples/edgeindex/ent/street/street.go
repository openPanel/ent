// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package street

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the street type in the database.
	Label = "street"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeCity holds the string denoting the city edge name in mutations.
	EdgeCity = "city"
	// Table holds the table name of the street in the database.
	Table = "streets"
	// CityTable is the table that holds the city relation/edge.
	CityTable = "streets"
	// CityInverseTable is the table name for the City entity.
	// It exists in this package in order to avoid circular dependency with the "city" package.
	CityInverseTable = "cities"
	// CityColumn is the table column denoting the city relation/edge.
	CityColumn = "city_streets"
)

// Columns holds all SQL columns for street fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "streets"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"city_streets",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Order defines the ordering method for the Street queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCityField orders the results by city field.
func ByCityField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCityStep(), sql.OrderByField(field, opts...))
	}
}
func newCityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CityTable, CityColumn),
	)
}
