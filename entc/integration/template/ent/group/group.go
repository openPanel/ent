// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package group

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMaxUsers holds the string denoting the max_users field in the database.
	FieldMaxUsers = "max_users"
	// Table holds the table name of the group in the database.
	Table = "groups"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldMaxUsers,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Order defines the ordering method for the Group queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMaxUsers orders the results by the max_users field.
func ByMaxUsers(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldMaxUsers, opts...).ToFunc()
}
