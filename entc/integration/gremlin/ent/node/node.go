// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package node

import (
	"time"

	"entgo.io/ent/dialect/gremlin/graph/dsl"
)

const (
	// Label holds the string label denoting the node type in the database.
	Label = "node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgePrev holds the string denoting the prev edge name in mutations.
	EdgePrev = "prev"
	// EdgeNext holds the string denoting the next edge name in mutations.
	EdgeNext = "next"
	// PrevInverseLabel holds the string label denoting the prev inverse edge type in the database.
	PrevInverseLabel = "node_next"
	// NextLabel holds the string label denoting the next edge type in the database.
	NextLabel = "node_next"
)

var (
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Order defines the ordering method for the Node queries.
type Order func(*dsl.Traversal)

// comment from another template.
