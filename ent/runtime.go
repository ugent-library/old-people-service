// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/ugent-library/people/ent/person"
	"github.com/ugent-library/people/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	personMixin := schema.Person{}.Mixin()
	personMixinFields0 := personMixin[0].Fields()
	_ = personMixinFields0
	personMixinFields1 := personMixin[1].Fields()
	_ = personMixinFields1
	personFields := schema.Person{}.Fields()
	_ = personFields
	// personDescDateCreated is the schema descriptor for date_created field.
	personDescDateCreated := personMixinFields1[0].Descriptor()
	// person.DefaultDateCreated holds the default value on creation for the date_created field.
	person.DefaultDateCreated = personDescDateCreated.Default.(func() time.Time)
	// personDescDateUpdated is the schema descriptor for date_updated field.
	personDescDateUpdated := personMixinFields1[1].Descriptor()
	// person.DefaultDateUpdated holds the default value on creation for the date_updated field.
	person.DefaultDateUpdated = personDescDateUpdated.Default.(func() time.Time)
	// person.UpdateDefaultDateUpdated holds the default value on update for the date_updated field.
	person.UpdateDefaultDateUpdated = personDescDateUpdated.UpdateDefault.(func() time.Time)
	// personDescActive is the schema descriptor for active field.
	personDescActive := personFields[0].Descriptor()
	// person.DefaultActive holds the default value on creation for the active field.
	person.DefaultActive = personDescActive.Default.(bool)
	// personDescID is the schema descriptor for id field.
	personDescID := personMixinFields0[0].Descriptor()
	// person.DefaultID holds the default value on creation for the id field.
	person.DefaultID = personDescID.Default.(func() string)
}
