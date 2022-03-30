// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChallengesColumns holds the columns for the "challenges" table.
	ChallengesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString, Unique: true, Size: 140},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 280},
		{Name: "outcome", Type: field.TypeBool, Nullable: true},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"bool"}, Default: "bool"},
	}
	// ChallengesTable holds the schema information for the "challenges" table.
	ChallengesTable = &schema.Table{
		Name:       "challenges",
		Columns:    ChallengesColumns,
		PrimaryKey: []*schema.Column{ChallengesColumns[0]},
	}
	// PredictionsColumns holds the columns for the "predictions" table.
	PredictionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "prognosis", Type: field.TypeBool},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
		{Name: "challenge_predictions", Type: field.TypeUUID},
		{Name: "user_predictions", Type: field.TypeUUID},
	}
	// PredictionsTable holds the schema information for the "predictions" table.
	PredictionsTable = &schema.Table{
		Name:       "predictions",
		Columns:    PredictionsColumns,
		PrimaryKey: []*schema.Column{PredictionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "predictions_challenges_predictions",
				Columns:    []*schema.Column{PredictionsColumns[5]},
				RefColumns: []*schema.Column{ChallengesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "predictions_users_predictions",
				Columns:    []*schema.Column{PredictionsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "prediction_challenge_predictions_user_predictions",
				Unique:  true,
				Columns: []*schema.Column{PredictionsColumns[5], PredictionsColumns[6]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "admin", Type: field.TypeBool, Default: false},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserSessionsColumns holds the columns for the "user_sessions" table.
	UserSessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "sid", Type: field.TypeString, Unique: true},
		{Name: "ip", Type: field.TypeString},
		{Name: "user_agent", Type: field.TypeString},
		{Name: "last_activity", Type: field.TypeTime},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
		{Name: "user_sessions", Type: field.TypeUUID},
	}
	// UserSessionsTable holds the schema information for the "user_sessions" table.
	UserSessionsTable = &schema.Table{
		Name:       "user_sessions",
		Columns:    UserSessionsColumns,
		PrimaryKey: []*schema.Column{UserSessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_sessions_users_sessions",
				Columns:    []*schema.Column{UserSessionsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChallengesTable,
		PredictionsTable,
		UsersTable,
		UserSessionsTable,
	}
)

func init() {
	PredictionsTable.ForeignKeys[0].RefTable = ChallengesTable
	PredictionsTable.ForeignKeys[1].RefTable = UsersTable
	UserSessionsTable.ForeignKeys[0].RefTable = UsersTable
}
