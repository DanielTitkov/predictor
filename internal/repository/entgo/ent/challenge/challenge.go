// Code generated by entc, DO NOT EDIT.

package challenge

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the challenge type in the database.
	Label = "challenge"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldOutcome holds the string denoting the outcome field in the database.
	FieldOutcome = "outcome"
	// FieldPublished holds the string denoting the published field in the database.
	FieldPublished = "published"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgePredictions holds the string denoting the predictions edge name in mutations.
	EdgePredictions = "predictions"
	// EdgeProofs holds the string denoting the proofs edge name in mutations.
	EdgeProofs = "proofs"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// Table holds the table name of the challenge in the database.
	Table = "challenges"
	// PredictionsTable is the table that holds the predictions relation/edge.
	PredictionsTable = "predictions"
	// PredictionsInverseTable is the table name for the Prediction entity.
	// It exists in this package in order to avoid circular dependency with the "prediction" package.
	PredictionsInverseTable = "predictions"
	// PredictionsColumn is the table column denoting the predictions relation/edge.
	PredictionsColumn = "challenge_predictions"
	// ProofsTable is the table that holds the proofs relation/edge.
	ProofsTable = "proofs"
	// ProofsInverseTable is the table name for the Proof entity.
	// It exists in this package in order to avoid circular dependency with the "proof" package.
	ProofsInverseTable = "proofs"
	// ProofsColumn is the table column denoting the proofs relation/edge.
	ProofsColumn = "challenge_proofs"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "challenges"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "user_challenges"
)

// Columns holds all SQL columns for challenge fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldContent,
	FieldDescription,
	FieldOutcome,
	FieldPublished,
	FieldStartTime,
	FieldEndTime,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "challenges"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_challenges",
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

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultPublished holds the default value on creation for the "published" field.
	DefaultPublished bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// TypeBool is the default value of the Type enum.
const DefaultType = TypeBool

// Type values.
const (
	TypeBool Type = "bool"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeBool:
		return nil
	default:
		return fmt.Errorf("challenge: invalid enum value for type field: %q", _type)
	}
}
