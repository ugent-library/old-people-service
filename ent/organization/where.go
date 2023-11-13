// Code generated by ent, DO NOT EDIT.

package organization

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/people-service/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldID, id))
}

// DateCreated applies equality check predicate on the "date_created" field. It's identical to DateCreatedEQ.
func DateCreated(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDateCreated, v))
}

// DateUpdated applies equality check predicate on the "date_updated" field. It's identical to DateUpdatedEQ.
func DateUpdated(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDateUpdated, v))
}

// PublicID applies equality check predicate on the "public_id" field. It's identical to PublicIDEQ.
func PublicID(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldPublicID, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldType, v))
}

// Acronym applies equality check predicate on the "acronym" field. It's identical to AcronymEQ.
func Acronym(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldAcronym, v))
}

// NameDut applies equality check predicate on the "name_dut" field. It's identical to NameDutEQ.
func NameDut(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldNameDut, v))
}

// NameEng applies equality check predicate on the "name_eng" field. It's identical to NameEngEQ.
func NameEng(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldNameEng, v))
}

// DateCreatedEQ applies the EQ predicate on the "date_created" field.
func DateCreatedEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDateCreated, v))
}

// DateCreatedNEQ applies the NEQ predicate on the "date_created" field.
func DateCreatedNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldDateCreated, v))
}

// DateCreatedIn applies the In predicate on the "date_created" field.
func DateCreatedIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldDateCreated, vs...))
}

// DateCreatedNotIn applies the NotIn predicate on the "date_created" field.
func DateCreatedNotIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldDateCreated, vs...))
}

// DateCreatedGT applies the GT predicate on the "date_created" field.
func DateCreatedGT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldDateCreated, v))
}

// DateCreatedGTE applies the GTE predicate on the "date_created" field.
func DateCreatedGTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldDateCreated, v))
}

// DateCreatedLT applies the LT predicate on the "date_created" field.
func DateCreatedLT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldDateCreated, v))
}

// DateCreatedLTE applies the LTE predicate on the "date_created" field.
func DateCreatedLTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldDateCreated, v))
}

// DateUpdatedEQ applies the EQ predicate on the "date_updated" field.
func DateUpdatedEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDateUpdated, v))
}

// DateUpdatedNEQ applies the NEQ predicate on the "date_updated" field.
func DateUpdatedNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldDateUpdated, v))
}

// DateUpdatedIn applies the In predicate on the "date_updated" field.
func DateUpdatedIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldDateUpdated, vs...))
}

// DateUpdatedNotIn applies the NotIn predicate on the "date_updated" field.
func DateUpdatedNotIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldDateUpdated, vs...))
}

// DateUpdatedGT applies the GT predicate on the "date_updated" field.
func DateUpdatedGT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldDateUpdated, v))
}

// DateUpdatedGTE applies the GTE predicate on the "date_updated" field.
func DateUpdatedGTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldDateUpdated, v))
}

// DateUpdatedLT applies the LT predicate on the "date_updated" field.
func DateUpdatedLT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldDateUpdated, v))
}

// DateUpdatedLTE applies the LTE predicate on the "date_updated" field.
func DateUpdatedLTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldDateUpdated, v))
}

// PublicIDEQ applies the EQ predicate on the "public_id" field.
func PublicIDEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldPublicID, v))
}

// PublicIDNEQ applies the NEQ predicate on the "public_id" field.
func PublicIDNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldPublicID, v))
}

// PublicIDIn applies the In predicate on the "public_id" field.
func PublicIDIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldPublicID, vs...))
}

// PublicIDNotIn applies the NotIn predicate on the "public_id" field.
func PublicIDNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldPublicID, vs...))
}

// PublicIDGT applies the GT predicate on the "public_id" field.
func PublicIDGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldPublicID, v))
}

// PublicIDGTE applies the GTE predicate on the "public_id" field.
func PublicIDGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldPublicID, v))
}

// PublicIDLT applies the LT predicate on the "public_id" field.
func PublicIDLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldPublicID, v))
}

// PublicIDLTE applies the LTE predicate on the "public_id" field.
func PublicIDLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldPublicID, v))
}

// PublicIDContains applies the Contains predicate on the "public_id" field.
func PublicIDContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldPublicID, v))
}

// PublicIDHasPrefix applies the HasPrefix predicate on the "public_id" field.
func PublicIDHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldPublicID, v))
}

// PublicIDHasSuffix applies the HasSuffix predicate on the "public_id" field.
func PublicIDHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldPublicID, v))
}

// PublicIDEqualFold applies the EqualFold predicate on the "public_id" field.
func PublicIDEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldPublicID, v))
}

// PublicIDContainsFold applies the ContainsFold predicate on the "public_id" field.
func PublicIDContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldPublicID, v))
}

// IdentifierIsNil applies the IsNil predicate on the "identifier" field.
func IdentifierIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldIdentifier))
}

// IdentifierNotNil applies the NotNil predicate on the "identifier" field.
func IdentifierNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldIdentifier))
}

// IdentifierValuesIsNil applies the IsNil predicate on the "identifier_values" field.
func IdentifierValuesIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldIdentifierValues))
}

// IdentifierValuesNotNil applies the NotNil predicate on the "identifier_values" field.
func IdentifierValuesNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldIdentifierValues))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldType, v))
}

// AcronymEQ applies the EQ predicate on the "acronym" field.
func AcronymEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldAcronym, v))
}

// AcronymNEQ applies the NEQ predicate on the "acronym" field.
func AcronymNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldAcronym, v))
}

// AcronymIn applies the In predicate on the "acronym" field.
func AcronymIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldAcronym, vs...))
}

// AcronymNotIn applies the NotIn predicate on the "acronym" field.
func AcronymNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldAcronym, vs...))
}

// AcronymGT applies the GT predicate on the "acronym" field.
func AcronymGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldAcronym, v))
}

// AcronymGTE applies the GTE predicate on the "acronym" field.
func AcronymGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldAcronym, v))
}

// AcronymLT applies the LT predicate on the "acronym" field.
func AcronymLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldAcronym, v))
}

// AcronymLTE applies the LTE predicate on the "acronym" field.
func AcronymLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldAcronym, v))
}

// AcronymContains applies the Contains predicate on the "acronym" field.
func AcronymContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldAcronym, v))
}

// AcronymHasPrefix applies the HasPrefix predicate on the "acronym" field.
func AcronymHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldAcronym, v))
}

// AcronymHasSuffix applies the HasSuffix predicate on the "acronym" field.
func AcronymHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldAcronym, v))
}

// AcronymIsNil applies the IsNil predicate on the "acronym" field.
func AcronymIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldAcronym))
}

// AcronymNotNil applies the NotNil predicate on the "acronym" field.
func AcronymNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldAcronym))
}

// AcronymEqualFold applies the EqualFold predicate on the "acronym" field.
func AcronymEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldAcronym, v))
}

// AcronymContainsFold applies the ContainsFold predicate on the "acronym" field.
func AcronymContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldAcronym, v))
}

// NameDutEQ applies the EQ predicate on the "name_dut" field.
func NameDutEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldNameDut, v))
}

// NameDutNEQ applies the NEQ predicate on the "name_dut" field.
func NameDutNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldNameDut, v))
}

// NameDutIn applies the In predicate on the "name_dut" field.
func NameDutIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldNameDut, vs...))
}

// NameDutNotIn applies the NotIn predicate on the "name_dut" field.
func NameDutNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldNameDut, vs...))
}

// NameDutGT applies the GT predicate on the "name_dut" field.
func NameDutGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldNameDut, v))
}

// NameDutGTE applies the GTE predicate on the "name_dut" field.
func NameDutGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldNameDut, v))
}

// NameDutLT applies the LT predicate on the "name_dut" field.
func NameDutLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldNameDut, v))
}

// NameDutLTE applies the LTE predicate on the "name_dut" field.
func NameDutLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldNameDut, v))
}

// NameDutContains applies the Contains predicate on the "name_dut" field.
func NameDutContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldNameDut, v))
}

// NameDutHasPrefix applies the HasPrefix predicate on the "name_dut" field.
func NameDutHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldNameDut, v))
}

// NameDutHasSuffix applies the HasSuffix predicate on the "name_dut" field.
func NameDutHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldNameDut, v))
}

// NameDutIsNil applies the IsNil predicate on the "name_dut" field.
func NameDutIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldNameDut))
}

// NameDutNotNil applies the NotNil predicate on the "name_dut" field.
func NameDutNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldNameDut))
}

// NameDutEqualFold applies the EqualFold predicate on the "name_dut" field.
func NameDutEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldNameDut, v))
}

// NameDutContainsFold applies the ContainsFold predicate on the "name_dut" field.
func NameDutContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldNameDut, v))
}

// NameEngEQ applies the EQ predicate on the "name_eng" field.
func NameEngEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldNameEng, v))
}

// NameEngNEQ applies the NEQ predicate on the "name_eng" field.
func NameEngNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldNameEng, v))
}

// NameEngIn applies the In predicate on the "name_eng" field.
func NameEngIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldNameEng, vs...))
}

// NameEngNotIn applies the NotIn predicate on the "name_eng" field.
func NameEngNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldNameEng, vs...))
}

// NameEngGT applies the GT predicate on the "name_eng" field.
func NameEngGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldNameEng, v))
}

// NameEngGTE applies the GTE predicate on the "name_eng" field.
func NameEngGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldNameEng, v))
}

// NameEngLT applies the LT predicate on the "name_eng" field.
func NameEngLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldNameEng, v))
}

// NameEngLTE applies the LTE predicate on the "name_eng" field.
func NameEngLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldNameEng, v))
}

// NameEngContains applies the Contains predicate on the "name_eng" field.
func NameEngContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldNameEng, v))
}

// NameEngHasPrefix applies the HasPrefix predicate on the "name_eng" field.
func NameEngHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldNameEng, v))
}

// NameEngHasSuffix applies the HasSuffix predicate on the "name_eng" field.
func NameEngHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldNameEng, v))
}

// NameEngIsNil applies the IsNil predicate on the "name_eng" field.
func NameEngIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldNameEng))
}

// NameEngNotNil applies the NotNil predicate on the "name_eng" field.
func NameEngNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldNameEng))
}

// NameEngEqualFold applies the EqualFold predicate on the "name_eng" field.
func NameEngEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldNameEng, v))
}

// NameEngContainsFold applies the ContainsFold predicate on the "name_eng" field.
func NameEngContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldNameEng, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
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
func Not(p predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		p(s.Not())
	})
}
