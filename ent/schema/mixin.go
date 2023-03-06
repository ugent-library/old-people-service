package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/oklog/ulid/v2"
)

type TimeMixin struct {
	ent.Mixin
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date_created").Default(time.Now).Immutable(),
		field.Time("date_updated").Default(time.Now).UpdateDefault(time.Now),
	}
}

type UUIDMixin struct {
	ent.Mixin
}

func (UUIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable().Unique().DefaultFunc(func() string {
			return ulid.Make().String()
		}),
	}
}
