package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		// ULID Implementation
		// ...
		// field.String("id").
		// 	GoType(ulid.ID("")).
		// 	DefaultFunc(func () ulid.ID  {
		// 		return ulid.MustNew("")
		// 	}),
		// field.String("user_id").
		// 	GoType(ulid.ID("")).
		// 	DefaultFunc(func () ulid.ID {
		// 		return ulid.MustNew("")
		// 	}),
		field.Int("user_id").Optional(),
		field.String("title").Default(""),
		field.Bool("completed").Default(false),
		field.Time("created_at").
			Default(time.Now()).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime DEFAULT CURRENT_TIMESTAMP",
			}).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
			}).
			Immutable(),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("todo").
			Unique().
			Field("user_id"),
	}
}
