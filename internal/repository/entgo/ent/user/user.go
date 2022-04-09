// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPicture holds the string denoting the picture field in the database.
	FieldPicture = "picture"
	// FieldAdmin holds the string denoting the admin field in the database.
	FieldAdmin = "admin"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldLocale holds the string denoting the locale field in the database.
	FieldLocale = "locale"
	// FieldMeta holds the string denoting the meta field in the database.
	FieldMeta = "meta"
	// EdgePredictions holds the string denoting the predictions edge name in mutations.
	EdgePredictions = "predictions"
	// EdgeSessions holds the string denoting the sessions edge name in mutations.
	EdgeSessions = "sessions"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PredictionsTable is the table that holds the predictions relation/edge.
	PredictionsTable = "predictions"
	// PredictionsInverseTable is the table name for the Prediction entity.
	// It exists in this package in order to avoid circular dependency with the "prediction" package.
	PredictionsInverseTable = "predictions"
	// PredictionsColumn is the table column denoting the predictions relation/edge.
	PredictionsColumn = "user_predictions"
	// SessionsTable is the table that holds the sessions relation/edge.
	SessionsTable = "user_sessions"
	// SessionsInverseTable is the table name for the UserSession entity.
	// It exists in this package in order to avoid circular dependency with the "usersession" package.
	SessionsInverseTable = "user_sessions"
	// SessionsColumn is the table column denoting the sessions relation/edge.
	SessionsColumn = "user_sessions"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldEmail,
	FieldPicture,
	FieldAdmin,
	FieldPasswordHash,
	FieldLocale,
	FieldMeta,
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

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultPicture holds the default value on creation for the "picture" field.
	DefaultPicture string
	// DefaultAdmin holds the default value on creation for the "admin" field.
	DefaultAdmin bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Locale defines the type for the "locale" enum field.
type Locale string

// LocaleRu is the default value of the Locale enum.
const DefaultLocale = LocaleRu

// Locale values.
const (
	LocaleEn Locale = "en"
	LocaleRu Locale = "ru"
)

func (l Locale) String() string {
	return string(l)
}

// LocaleValidator is a validator for the "locale" field enum values. It is called by the builders before save.
func LocaleValidator(l Locale) error {
	switch l {
	case LocaleEn, LocaleRu:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for locale field: %q", l)
	}
}
