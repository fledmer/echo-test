package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RequestLog holds the schema definition for the RequestLog entity.
type RequestLog struct {
	ent.Schema
}

// Fields of the RequestLog.
func (RequestLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("method").
			NotEmpty().
			Comment("HTTP method (GET, POST, etc.)"),
		field.String("path").
			NotEmpty().
			Comment("Request path"),
		field.Text("headers").
			Optional().
			Comment("Request headers as JSON string"),
		field.Text("body").
			Optional().
			Comment("Request body"),
		field.String("ip").
			Optional().
			Comment("Client IP address"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("Timestamp of the request"),
	}
}

// Edges of the RequestLog.
func (RequestLog) Edges() []ent.Edge {
	return nil
}
