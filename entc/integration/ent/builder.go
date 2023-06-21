// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/builder"
)

// Builder is the model entity for the Builder schema.
type Builder struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Builder) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case builder.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Builder fields.
func (b *Builder) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case builder.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Builder.
// This includes values selected through modifiers, order, etc.
func (b *Builder) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// Update returns a builder for updating this Builder.
// Note that you need to call Builder.Unwrap() before calling this method if this Builder
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Builder) Update() *BuilderUpdateOne {
	return NewBuilderClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Builder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Builder) Unwrap() *Builder {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Builder is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Builder) String() string {
	var builder strings.Builder
	builder.WriteString("Builder(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Builders is a parsable slice of Builder.
type Builders []*Builder
