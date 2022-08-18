package schema

import (
	"go-gqlgen-template/ent/schema/ulid"
	"go-gqlgen-template/pkg/const/globalid"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// ULID Implementation
		// ...
		field.String("id").
			GoType(ulid.ID("")).
			DefaultFunc(func() ulid.ID {
				return ulid.MustNew(globalid.New().User.Prefix)
			}),
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.String("username").
			Unique().
			NotEmpty(),
		field.String("email"),
		field.String("phone"),
		field.String("website").Optional(),
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

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todo", Todo.Type),
	}
}
