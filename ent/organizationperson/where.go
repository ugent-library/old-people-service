// Code generated by ent, DO NOT EDIT.

package organizationperson

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/people-service/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldLTE(FieldID, id))
}

// DateCreated applies equality check predicate on the "date_created" field. It's identical to DateCreatedEQ.
func DateCreated(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldEQ(FieldDateCreated, v))
}

// DateUpdated applies equality check predicate on the "date_updated" field. It's identical to DateUpdatedEQ.
func DateUpdated(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldEQ(FieldDateUpdated, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldEQ(FieldOrganizationID, v))
}

// PersonID applies equality check predicate on the "person_id" field. It's identical to PersonIDEQ.
func PersonID(v int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldEQ(FieldPersonID, v))
}

// From applies equality check predicate on the "from" field. It's identical to FromEQ.
func From(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldEQ(FieldFrom, v))
}

// Until applies equality check predicate on the "until" field. It's identical to UntilEQ.
func Until(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldEQ(FieldUntil, v))
}

// DateCreatedEQ applies the EQ predicate on the "date_created" field.
func DateCreatedEQ(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldEQ(FieldDateCreated, v))
}

// DateCreatedNEQ applies the NEQ predicate on the "date_created" field.
func DateCreatedNEQ(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNEQ(FieldDateCreated, v))
}

// DateCreatedIn applies the In predicate on the "date_created" field.
func DateCreatedIn(vs ...time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldIn(FieldDateCreated, vs...))
}

// DateCreatedNotIn applies the NotIn predicate on the "date_created" field.
func DateCreatedNotIn(vs ...time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNotIn(FieldDateCreated, vs...))
}

// DateCreatedGT applies the GT predicate on the "date_created" field.
func DateCreatedGT(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldGT(FieldDateCreated, v))
}

// DateCreatedGTE applies the GTE predicate on the "date_created" field.
func DateCreatedGTE(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldGTE(FieldDateCreated, v))
}

// DateCreatedLT applies the LT predicate on the "date_created" field.
func DateCreatedLT(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldLT(FieldDateCreated, v))
}

// DateCreatedLTE applies the LTE predicate on the "date_created" field.
func DateCreatedLTE(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldLTE(FieldDateCreated, v))
}

// DateUpdatedEQ applies the EQ predicate on the "date_updated" field.
func DateUpdatedEQ(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldEQ(FieldDateUpdated, v))
}

// DateUpdatedNEQ applies the NEQ predicate on the "date_updated" field.
func DateUpdatedNEQ(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNEQ(FieldDateUpdated, v))
}

// DateUpdatedIn applies the In predicate on the "date_updated" field.
func DateUpdatedIn(vs ...time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldIn(FieldDateUpdated, vs...))
}

// DateUpdatedNotIn applies the NotIn predicate on the "date_updated" field.
func DateUpdatedNotIn(vs ...time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNotIn(FieldDateUpdated, vs...))
}

// DateUpdatedGT applies the GT predicate on the "date_updated" field.
func DateUpdatedGT(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldGT(FieldDateUpdated, v))
}

// DateUpdatedGTE applies the GTE predicate on the "date_updated" field.
func DateUpdatedGTE(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldGTE(FieldDateUpdated, v))
}

// DateUpdatedLT applies the LT predicate on the "date_updated" field.
func DateUpdatedLT(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldLT(FieldDateUpdated, v))
}

// DateUpdatedLTE applies the LTE predicate on the "date_updated" field.
func DateUpdatedLTE(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldLTE(FieldDateUpdated, v))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// OrganizationIDGT applies the GT predicate on the "organization_id" field.
func OrganizationIDGT(v int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldGT(FieldOrganizationID, v))
}

// OrganizationIDGTE applies the GTE predicate on the "organization_id" field.
func OrganizationIDGTE(v int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldGTE(FieldOrganizationID, v))
}

// OrganizationIDLT applies the LT predicate on the "organization_id" field.
func OrganizationIDLT(v int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldLT(FieldOrganizationID, v))
}

// OrganizationIDLTE applies the LTE predicate on the "organization_id" field.
func OrganizationIDLTE(v int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldLTE(FieldOrganizationID, v))
}

// PersonIDEQ applies the EQ predicate on the "person_id" field.
func PersonIDEQ(v int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldEQ(FieldPersonID, v))
}

// PersonIDNEQ applies the NEQ predicate on the "person_id" field.
func PersonIDNEQ(v int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNEQ(FieldPersonID, v))
}

// PersonIDIn applies the In predicate on the "person_id" field.
func PersonIDIn(vs ...int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldIn(FieldPersonID, vs...))
}

// PersonIDNotIn applies the NotIn predicate on the "person_id" field.
func PersonIDNotIn(vs ...int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNotIn(FieldPersonID, vs...))
}

// PersonIDGT applies the GT predicate on the "person_id" field.
func PersonIDGT(v int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldGT(FieldPersonID, v))
}

// PersonIDGTE applies the GTE predicate on the "person_id" field.
func PersonIDGTE(v int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldGTE(FieldPersonID, v))
}

// PersonIDLT applies the LT predicate on the "person_id" field.
func PersonIDLT(v int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldLT(FieldPersonID, v))
}

// PersonIDLTE applies the LTE predicate on the "person_id" field.
func PersonIDLTE(v int) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldLTE(FieldPersonID, v))
}

// FromEQ applies the EQ predicate on the "from" field.
func FromEQ(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldEQ(FieldFrom, v))
}

// FromNEQ applies the NEQ predicate on the "from" field.
func FromNEQ(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNEQ(FieldFrom, v))
}

// FromIn applies the In predicate on the "from" field.
func FromIn(vs ...time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldIn(FieldFrom, vs...))
}

// FromNotIn applies the NotIn predicate on the "from" field.
func FromNotIn(vs ...time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNotIn(FieldFrom, vs...))
}

// FromGT applies the GT predicate on the "from" field.
func FromGT(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldGT(FieldFrom, v))
}

// FromGTE applies the GTE predicate on the "from" field.
func FromGTE(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldGTE(FieldFrom, v))
}

// FromLT applies the LT predicate on the "from" field.
func FromLT(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldLT(FieldFrom, v))
}

// FromLTE applies the LTE predicate on the "from" field.
func FromLTE(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldLTE(FieldFrom, v))
}

// UntilEQ applies the EQ predicate on the "until" field.
func UntilEQ(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldEQ(FieldUntil, v))
}

// UntilNEQ applies the NEQ predicate on the "until" field.
func UntilNEQ(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNEQ(FieldUntil, v))
}

// UntilIn applies the In predicate on the "until" field.
func UntilIn(vs ...time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldIn(FieldUntil, vs...))
}

// UntilNotIn applies the NotIn predicate on the "until" field.
func UntilNotIn(vs ...time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNotIn(FieldUntil, vs...))
}

// UntilGT applies the GT predicate on the "until" field.
func UntilGT(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldGT(FieldUntil, v))
}

// UntilGTE applies the GTE predicate on the "until" field.
func UntilGTE(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldGTE(FieldUntil, v))
}

// UntilLT applies the LT predicate on the "until" field.
func UntilLT(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldLT(FieldUntil, v))
}

// UntilLTE applies the LTE predicate on the "until" field.
func UntilLTE(v time.Time) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldLTE(FieldUntil, v))
}

// UntilIsNil applies the IsNil predicate on the "until" field.
func UntilIsNil() predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldIsNull(FieldUntil))
}

// UntilNotNil applies the NotNil predicate on the "until" field.
func UntilNotNil() predicate.OrganizationPerson {
	return predicate.OrganizationPerson(sql.FieldNotNull(FieldUntil))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrganizationPerson) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrganizationPerson) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrganizationPerson) predicate.OrganizationPerson {
	return predicate.OrganizationPerson(func(s *sql.Selector) {
		p(s.Not())
	})
}
