// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/badge"
)

// Badge is the model entity for the Badge schema.
type Badge struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// Meta holds the value of the "meta" field.
	Meta map[string]interface{} `json:"meta,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BadgeQuery when eager-loading is set.
	Edges BadgeEdges `json:"edges"`
}

// BadgeEdges holds the relations/edges for other nodes in the graph.
type BadgeEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e BadgeEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Badge) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case badge.FieldMeta:
			values[i] = new([]byte)
		case badge.FieldActive:
			values[i] = new(sql.NullBool)
		case badge.FieldID:
			values[i] = new(sql.NullInt64)
		case badge.FieldType:
			values[i] = new(sql.NullString)
		case badge.FieldCreateTime, badge.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Badge", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Badge fields.
func (b *Badge) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case badge.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case badge.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				b.CreateTime = value.Time
			}
		case badge.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				b.UpdateTime = value.Time
			}
		case badge.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				b.Type = value.String
			}
		case badge.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				b.Active = value.Bool
			}
		case badge.FieldMeta:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field meta", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &b.Meta); err != nil {
					return fmt.Errorf("unmarshal field meta: %w", err)
				}
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Badge entity.
func (b *Badge) QueryUsers() *UserQuery {
	return (&BadgeClient{config: b.config}).QueryUsers(b)
}

// Update returns a builder for updating this Badge.
// Note that you need to call Badge.Unwrap() before calling this method if this Badge
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Badge) Update() *BadgeUpdateOne {
	return (&BadgeClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Badge entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Badge) Unwrap() *Badge {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Badge is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Badge) String() string {
	var builder strings.Builder
	builder.WriteString("Badge(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(b.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(b.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", type=")
	builder.WriteString(b.Type)
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", b.Active))
	builder.WriteString(", meta=")
	builder.WriteString(fmt.Sprintf("%v", b.Meta))
	builder.WriteByte(')')
	return builder.String()
}

// Badges is a parsable slice of Badge.
type Badges []*Badge

func (b Badges) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
