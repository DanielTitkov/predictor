// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/challenge"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/proof"
	"github.com/google/uuid"
)

// Proof is the model entity for the Proof schema.
type Proof struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Link holds the value of the "link" field.
	Link string `json:"link,omitempty"`
	// Meta holds the value of the "meta" field.
	Meta map[string]interface{} `json:"meta,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProofQuery when eager-loading is set.
	Edges            ProofEdges `json:"edges"`
	challenge_proofs *uuid.UUID
}

// ProofEdges holds the relations/edges for other nodes in the graph.
type ProofEdges struct {
	// Challenge holds the value of the challenge edge.
	Challenge *Challenge `json:"challenge,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ChallengeOrErr returns the Challenge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProofEdges) ChallengeOrErr() (*Challenge, error) {
	if e.loadedTypes[0] {
		if e.Challenge == nil {
			// The edge challenge was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: challenge.Label}
		}
		return e.Challenge, nil
	}
	return nil, &NotLoadedError{edge: "challenge"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Proof) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case proof.FieldMeta:
			values[i] = new([]byte)
		case proof.FieldContent, proof.FieldLink:
			values[i] = new(sql.NullString)
		case proof.FieldCreateTime, proof.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case proof.FieldID:
			values[i] = new(uuid.UUID)
		case proof.ForeignKeys[0]: // challenge_proofs
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Proof", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Proof fields.
func (pr *Proof) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case proof.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case proof.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pr.CreateTime = value.Time
			}
		case proof.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pr.UpdateTime = value.Time
			}
		case proof.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				pr.Content = value.String
			}
		case proof.FieldLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link", values[i])
			} else if value.Valid {
				pr.Link = value.String
			}
		case proof.FieldMeta:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field meta", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Meta); err != nil {
					return fmt.Errorf("unmarshal field meta: %w", err)
				}
			}
		case proof.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field challenge_proofs", values[i])
			} else if value.Valid {
				pr.challenge_proofs = new(uuid.UUID)
				*pr.challenge_proofs = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryChallenge queries the "challenge" edge of the Proof entity.
func (pr *Proof) QueryChallenge() *ChallengeQuery {
	return (&ProofClient{config: pr.config}).QueryChallenge(pr)
}

// Update returns a builder for updating this Proof.
// Note that you need to call Proof.Unwrap() before calling this method if this Proof
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Proof) Update() *ProofUpdateOne {
	return (&ProofClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Proof entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Proof) Unwrap() *Proof {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Proof is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Proof) String() string {
	var builder strings.Builder
	builder.WriteString("Proof(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(pr.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(pr.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", content=")
	builder.WriteString(pr.Content)
	builder.WriteString(", link=")
	builder.WriteString(pr.Link)
	builder.WriteString(", meta=")
	builder.WriteString(fmt.Sprintf("%v", pr.Meta))
	builder.WriteByte(')')
	return builder.String()
}

// Proofs is a parsable slice of Proof.
type Proofs []*Proof

func (pr Proofs) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}