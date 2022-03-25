// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/challenge"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/prediction"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/user"
	"github.com/google/uuid"
)

// Prediction is the model entity for the Prediction schema.
type Prediction struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Prognosis holds the value of the "prognosis" field.
	Prognosis bool `json:"prognosis,omitempty"`
	// Meta holds the value of the "meta" field.
	Meta map[string]interface{} `json:"meta,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PredictionQuery when eager-loading is set.
	Edges                 PredictionEdges `json:"edges"`
	challenge_predictions *uuid.UUID
	user_predictions      *uuid.UUID
}

// PredictionEdges holds the relations/edges for other nodes in the graph.
type PredictionEdges struct {
	// Challenge holds the value of the challenge edge.
	Challenge *Challenge `json:"challenge,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ChallengeOrErr returns the Challenge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PredictionEdges) ChallengeOrErr() (*Challenge, error) {
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

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PredictionEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Prediction) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case prediction.FieldMeta:
			values[i] = new([]byte)
		case prediction.FieldPrognosis:
			values[i] = new(sql.NullBool)
		case prediction.FieldCreateTime, prediction.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case prediction.FieldID:
			values[i] = new(uuid.UUID)
		case prediction.ForeignKeys[0]: // challenge_predictions
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case prediction.ForeignKeys[1]: // user_predictions
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Prediction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Prediction fields.
func (pr *Prediction) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case prediction.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case prediction.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pr.CreateTime = value.Time
			}
		case prediction.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pr.UpdateTime = value.Time
			}
		case prediction.FieldPrognosis:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field prognosis", values[i])
			} else if value.Valid {
				pr.Prognosis = value.Bool
			}
		case prediction.FieldMeta:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field meta", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Meta); err != nil {
					return fmt.Errorf("unmarshal field meta: %w", err)
				}
			}
		case prediction.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field challenge_predictions", values[i])
			} else if value.Valid {
				pr.challenge_predictions = new(uuid.UUID)
				*pr.challenge_predictions = *value.S.(*uuid.UUID)
			}
		case prediction.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_predictions", values[i])
			} else if value.Valid {
				pr.user_predictions = new(uuid.UUID)
				*pr.user_predictions = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryChallenge queries the "challenge" edge of the Prediction entity.
func (pr *Prediction) QueryChallenge() *ChallengeQuery {
	return (&PredictionClient{config: pr.config}).QueryChallenge(pr)
}

// QueryUser queries the "user" edge of the Prediction entity.
func (pr *Prediction) QueryUser() *UserQuery {
	return (&PredictionClient{config: pr.config}).QueryUser(pr)
}

// Update returns a builder for updating this Prediction.
// Note that you need to call Prediction.Unwrap() before calling this method if this Prediction
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Prediction) Update() *PredictionUpdateOne {
	return (&PredictionClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Prediction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Prediction) Unwrap() *Prediction {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Prediction is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Prediction) String() string {
	var builder strings.Builder
	builder.WriteString("Prediction(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(pr.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(pr.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", prognosis=")
	builder.WriteString(fmt.Sprintf("%v", pr.Prognosis))
	builder.WriteString(", meta=")
	builder.WriteString(fmt.Sprintf("%v", pr.Meta))
	builder.WriteByte(')')
	return builder.String()
}

// Predictions is a parsable slice of Prediction.
type Predictions []*Prediction

func (pr Predictions) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
