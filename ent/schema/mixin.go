package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/oklog/ulid/v2"
)

type TimeMixin struct {
	mixin.Schema
}

var timeUTC = func() time.Time {
	return time.Now().UTC()
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date_created").Default(timeUTC).Immutable(),
		field.Time("date_updated").Default(timeUTC).UpdateDefault(timeUTC),
	}
}

type UUIDMixin struct {
	mixin.Schema
}

func (UUIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable().Unique().DefaultFunc(func() string {
			return ulid.Make().String()
		}),
	}
}
