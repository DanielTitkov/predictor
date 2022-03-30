package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("sid").NotEmpty().Unique().Immutable(),
		field.String("ip"),
		field.String("user_agent"),
		field.Time("last_activity").Default(time.Now),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("user", User.Type).Ref("sessions").Unique().Required(),
	}
}

func (Session) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
