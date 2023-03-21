// Code generated by ent, DO NOT EDIT.

package person

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/people/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldID, id))
}

// DateCreated applies equality check predicate on the "date_created" field. It's identical to DateCreatedEQ.
func DateCreated(v time.Time) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldDateCreated, v))
}

// DateUpdated applies equality check predicate on the "date_updated" field. It's identical to DateUpdatedEQ.
func DateUpdated(v time.Time) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldDateUpdated, v))
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldActive, v))
}

// BirthDate applies equality check predicate on the "birth_date" field. It's identical to BirthDateEQ.
func BirthDate(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldBirthDate, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldEmail, v))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldFirstName, v))
}

// FullName applies equality check predicate on the "full_name" field. It's identical to FullNameEQ.
func FullName(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldFullName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldLastName, v))
}

// Orcid applies equality check predicate on the "orcid" field. It's identical to OrcidEQ.
func Orcid(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldOrcid, v))
}

// OrcidToken applies equality check predicate on the "orcid_token" field. It's identical to OrcidTokenEQ.
func OrcidToken(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldOrcidToken, v))
}

// PreferredFirstName applies equality check predicate on the "preferred_first_name" field. It's identical to PreferredFirstNameEQ.
func PreferredFirstName(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldPreferredFirstName, v))
}

// PreferredLastName applies equality check predicate on the "preferred_last_name" field. It's identical to PreferredLastNameEQ.
func PreferredLastName(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldPreferredLastName, v))
}

// JobTitle applies equality check predicate on the "job_title" field. It's identical to JobTitleEQ.
func JobTitle(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldJobTitle, v))
}

// DateCreatedEQ applies the EQ predicate on the "date_created" field.
func DateCreatedEQ(v time.Time) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldDateCreated, v))
}

// DateCreatedNEQ applies the NEQ predicate on the "date_created" field.
func DateCreatedNEQ(v time.Time) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldDateCreated, v))
}

// DateCreatedIn applies the In predicate on the "date_created" field.
func DateCreatedIn(vs ...time.Time) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldDateCreated, vs...))
}

// DateCreatedNotIn applies the NotIn predicate on the "date_created" field.
func DateCreatedNotIn(vs ...time.Time) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldDateCreated, vs...))
}

// DateCreatedGT applies the GT predicate on the "date_created" field.
func DateCreatedGT(v time.Time) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldDateCreated, v))
}

// DateCreatedGTE applies the GTE predicate on the "date_created" field.
func DateCreatedGTE(v time.Time) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldDateCreated, v))
}

// DateCreatedLT applies the LT predicate on the "date_created" field.
func DateCreatedLT(v time.Time) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldDateCreated, v))
}

// DateCreatedLTE applies the LTE predicate on the "date_created" field.
func DateCreatedLTE(v time.Time) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldDateCreated, v))
}

// DateUpdatedEQ applies the EQ predicate on the "date_updated" field.
func DateUpdatedEQ(v time.Time) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldDateUpdated, v))
}

// DateUpdatedNEQ applies the NEQ predicate on the "date_updated" field.
func DateUpdatedNEQ(v time.Time) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldDateUpdated, v))
}

// DateUpdatedIn applies the In predicate on the "date_updated" field.
func DateUpdatedIn(vs ...time.Time) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldDateUpdated, vs...))
}

// DateUpdatedNotIn applies the NotIn predicate on the "date_updated" field.
func DateUpdatedNotIn(vs ...time.Time) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldDateUpdated, vs...))
}

// DateUpdatedGT applies the GT predicate on the "date_updated" field.
func DateUpdatedGT(v time.Time) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldDateUpdated, v))
}

// DateUpdatedGTE applies the GTE predicate on the "date_updated" field.
func DateUpdatedGTE(v time.Time) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldDateUpdated, v))
}

// DateUpdatedLT applies the LT predicate on the "date_updated" field.
func DateUpdatedLT(v time.Time) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldDateUpdated, v))
}

// DateUpdatedLTE applies the LTE predicate on the "date_updated" field.
func DateUpdatedLTE(v time.Time) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldDateUpdated, v))
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldActive, v))
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldActive, v))
}

// BirthDateEQ applies the EQ predicate on the "birth_date" field.
func BirthDateEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldBirthDate, v))
}

// BirthDateNEQ applies the NEQ predicate on the "birth_date" field.
func BirthDateNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldBirthDate, v))
}

// BirthDateIn applies the In predicate on the "birth_date" field.
func BirthDateIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldBirthDate, vs...))
}

// BirthDateNotIn applies the NotIn predicate on the "birth_date" field.
func BirthDateNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldBirthDate, vs...))
}

// BirthDateGT applies the GT predicate on the "birth_date" field.
func BirthDateGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldBirthDate, v))
}

// BirthDateGTE applies the GTE predicate on the "birth_date" field.
func BirthDateGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldBirthDate, v))
}

// BirthDateLT applies the LT predicate on the "birth_date" field.
func BirthDateLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldBirthDate, v))
}

// BirthDateLTE applies the LTE predicate on the "birth_date" field.
func BirthDateLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldBirthDate, v))
}

// BirthDateContains applies the Contains predicate on the "birth_date" field.
func BirthDateContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldBirthDate, v))
}

// BirthDateHasPrefix applies the HasPrefix predicate on the "birth_date" field.
func BirthDateHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldBirthDate, v))
}

// BirthDateHasSuffix applies the HasSuffix predicate on the "birth_date" field.
func BirthDateHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldBirthDate, v))
}

// BirthDateIsNil applies the IsNil predicate on the "birth_date" field.
func BirthDateIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldBirthDate))
}

// BirthDateNotNil applies the NotNil predicate on the "birth_date" field.
func BirthDateNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldBirthDate))
}

// BirthDateEqualFold applies the EqualFold predicate on the "birth_date" field.
func BirthDateEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldBirthDate, v))
}

// BirthDateContainsFold applies the ContainsFold predicate on the "birth_date" field.
func BirthDateContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldBirthDate, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldEmail, v))
}

// OtherIDIsNil applies the IsNil predicate on the "other_id" field.
func OtherIDIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldOtherID))
}

// OtherIDNotNil applies the NotNil predicate on the "other_id" field.
func OtherIDNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldOtherID))
}

// OrganizationIDIsNil applies the IsNil predicate on the "organization_id" field.
func OrganizationIDIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldOrganizationID))
}

// OrganizationIDNotNil applies the NotNil predicate on the "organization_id" field.
func OrganizationIDNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldOrganizationID))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameIsNil applies the IsNil predicate on the "first_name" field.
func FirstNameIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldFirstName))
}

// FirstNameNotNil applies the NotNil predicate on the "first_name" field.
func FirstNameNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldFirstName))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldFirstName, v))
}

// FullNameEQ applies the EQ predicate on the "full_name" field.
func FullNameEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldFullName, v))
}

// FullNameNEQ applies the NEQ predicate on the "full_name" field.
func FullNameNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldFullName, v))
}

// FullNameIn applies the In predicate on the "full_name" field.
func FullNameIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldFullName, vs...))
}

// FullNameNotIn applies the NotIn predicate on the "full_name" field.
func FullNameNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldFullName, vs...))
}

// FullNameGT applies the GT predicate on the "full_name" field.
func FullNameGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldFullName, v))
}

// FullNameGTE applies the GTE predicate on the "full_name" field.
func FullNameGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldFullName, v))
}

// FullNameLT applies the LT predicate on the "full_name" field.
func FullNameLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldFullName, v))
}

// FullNameLTE applies the LTE predicate on the "full_name" field.
func FullNameLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldFullName, v))
}

// FullNameContains applies the Contains predicate on the "full_name" field.
func FullNameContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldFullName, v))
}

// FullNameHasPrefix applies the HasPrefix predicate on the "full_name" field.
func FullNameHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldFullName, v))
}

// FullNameHasSuffix applies the HasSuffix predicate on the "full_name" field.
func FullNameHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldFullName, v))
}

// FullNameIsNil applies the IsNil predicate on the "full_name" field.
func FullNameIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldFullName))
}

// FullNameNotNil applies the NotNil predicate on the "full_name" field.
func FullNameNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldFullName))
}

// FullNameEqualFold applies the EqualFold predicate on the "full_name" field.
func FullNameEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldFullName, v))
}

// FullNameContainsFold applies the ContainsFold predicate on the "full_name" field.
func FullNameContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldFullName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameIsNil applies the IsNil predicate on the "last_name" field.
func LastNameIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldLastName))
}

// LastNameNotNil applies the NotNil predicate on the "last_name" field.
func LastNameNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldLastName))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldLastName, v))
}

// CategoryIsNil applies the IsNil predicate on the "category" field.
func CategoryIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldCategory))
}

// CategoryNotNil applies the NotNil predicate on the "category" field.
func CategoryNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldCategory))
}

// OrcidEQ applies the EQ predicate on the "orcid" field.
func OrcidEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldOrcid, v))
}

// OrcidNEQ applies the NEQ predicate on the "orcid" field.
func OrcidNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldOrcid, v))
}

// OrcidIn applies the In predicate on the "orcid" field.
func OrcidIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldOrcid, vs...))
}

// OrcidNotIn applies the NotIn predicate on the "orcid" field.
func OrcidNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldOrcid, vs...))
}

// OrcidGT applies the GT predicate on the "orcid" field.
func OrcidGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldOrcid, v))
}

// OrcidGTE applies the GTE predicate on the "orcid" field.
func OrcidGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldOrcid, v))
}

// OrcidLT applies the LT predicate on the "orcid" field.
func OrcidLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldOrcid, v))
}

// OrcidLTE applies the LTE predicate on the "orcid" field.
func OrcidLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldOrcid, v))
}

// OrcidContains applies the Contains predicate on the "orcid" field.
func OrcidContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldOrcid, v))
}

// OrcidHasPrefix applies the HasPrefix predicate on the "orcid" field.
func OrcidHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldOrcid, v))
}

// OrcidHasSuffix applies the HasSuffix predicate on the "orcid" field.
func OrcidHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldOrcid, v))
}

// OrcidIsNil applies the IsNil predicate on the "orcid" field.
func OrcidIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldOrcid))
}

// OrcidNotNil applies the NotNil predicate on the "orcid" field.
func OrcidNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldOrcid))
}

// OrcidEqualFold applies the EqualFold predicate on the "orcid" field.
func OrcidEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldOrcid, v))
}

// OrcidContainsFold applies the ContainsFold predicate on the "orcid" field.
func OrcidContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldOrcid, v))
}

// OrcidTokenEQ applies the EQ predicate on the "orcid_token" field.
func OrcidTokenEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldOrcidToken, v))
}

// OrcidTokenNEQ applies the NEQ predicate on the "orcid_token" field.
func OrcidTokenNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldOrcidToken, v))
}

// OrcidTokenIn applies the In predicate on the "orcid_token" field.
func OrcidTokenIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldOrcidToken, vs...))
}

// OrcidTokenNotIn applies the NotIn predicate on the "orcid_token" field.
func OrcidTokenNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldOrcidToken, vs...))
}

// OrcidTokenGT applies the GT predicate on the "orcid_token" field.
func OrcidTokenGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldOrcidToken, v))
}

// OrcidTokenGTE applies the GTE predicate on the "orcid_token" field.
func OrcidTokenGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldOrcidToken, v))
}

// OrcidTokenLT applies the LT predicate on the "orcid_token" field.
func OrcidTokenLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldOrcidToken, v))
}

// OrcidTokenLTE applies the LTE predicate on the "orcid_token" field.
func OrcidTokenLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldOrcidToken, v))
}

// OrcidTokenContains applies the Contains predicate on the "orcid_token" field.
func OrcidTokenContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldOrcidToken, v))
}

// OrcidTokenHasPrefix applies the HasPrefix predicate on the "orcid_token" field.
func OrcidTokenHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldOrcidToken, v))
}

// OrcidTokenHasSuffix applies the HasSuffix predicate on the "orcid_token" field.
func OrcidTokenHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldOrcidToken, v))
}

// OrcidTokenIsNil applies the IsNil predicate on the "orcid_token" field.
func OrcidTokenIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldOrcidToken))
}

// OrcidTokenNotNil applies the NotNil predicate on the "orcid_token" field.
func OrcidTokenNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldOrcidToken))
}

// OrcidTokenEqualFold applies the EqualFold predicate on the "orcid_token" field.
func OrcidTokenEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldOrcidToken, v))
}

// OrcidTokenContainsFold applies the ContainsFold predicate on the "orcid_token" field.
func OrcidTokenContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldOrcidToken, v))
}

// PreferredFirstNameEQ applies the EQ predicate on the "preferred_first_name" field.
func PreferredFirstNameEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldPreferredFirstName, v))
}

// PreferredFirstNameNEQ applies the NEQ predicate on the "preferred_first_name" field.
func PreferredFirstNameNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldPreferredFirstName, v))
}

// PreferredFirstNameIn applies the In predicate on the "preferred_first_name" field.
func PreferredFirstNameIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldPreferredFirstName, vs...))
}

// PreferredFirstNameNotIn applies the NotIn predicate on the "preferred_first_name" field.
func PreferredFirstNameNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldPreferredFirstName, vs...))
}

// PreferredFirstNameGT applies the GT predicate on the "preferred_first_name" field.
func PreferredFirstNameGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldPreferredFirstName, v))
}

// PreferredFirstNameGTE applies the GTE predicate on the "preferred_first_name" field.
func PreferredFirstNameGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldPreferredFirstName, v))
}

// PreferredFirstNameLT applies the LT predicate on the "preferred_first_name" field.
func PreferredFirstNameLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldPreferredFirstName, v))
}

// PreferredFirstNameLTE applies the LTE predicate on the "preferred_first_name" field.
func PreferredFirstNameLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldPreferredFirstName, v))
}

// PreferredFirstNameContains applies the Contains predicate on the "preferred_first_name" field.
func PreferredFirstNameContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldPreferredFirstName, v))
}

// PreferredFirstNameHasPrefix applies the HasPrefix predicate on the "preferred_first_name" field.
func PreferredFirstNameHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldPreferredFirstName, v))
}

// PreferredFirstNameHasSuffix applies the HasSuffix predicate on the "preferred_first_name" field.
func PreferredFirstNameHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldPreferredFirstName, v))
}

// PreferredFirstNameIsNil applies the IsNil predicate on the "preferred_first_name" field.
func PreferredFirstNameIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldPreferredFirstName))
}

// PreferredFirstNameNotNil applies the NotNil predicate on the "preferred_first_name" field.
func PreferredFirstNameNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldPreferredFirstName))
}

// PreferredFirstNameEqualFold applies the EqualFold predicate on the "preferred_first_name" field.
func PreferredFirstNameEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldPreferredFirstName, v))
}

// PreferredFirstNameContainsFold applies the ContainsFold predicate on the "preferred_first_name" field.
func PreferredFirstNameContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldPreferredFirstName, v))
}

// PreferredLastNameEQ applies the EQ predicate on the "preferred_last_name" field.
func PreferredLastNameEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldPreferredLastName, v))
}

// PreferredLastNameNEQ applies the NEQ predicate on the "preferred_last_name" field.
func PreferredLastNameNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldPreferredLastName, v))
}

// PreferredLastNameIn applies the In predicate on the "preferred_last_name" field.
func PreferredLastNameIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldPreferredLastName, vs...))
}

// PreferredLastNameNotIn applies the NotIn predicate on the "preferred_last_name" field.
func PreferredLastNameNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldPreferredLastName, vs...))
}

// PreferredLastNameGT applies the GT predicate on the "preferred_last_name" field.
func PreferredLastNameGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldPreferredLastName, v))
}

// PreferredLastNameGTE applies the GTE predicate on the "preferred_last_name" field.
func PreferredLastNameGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldPreferredLastName, v))
}

// PreferredLastNameLT applies the LT predicate on the "preferred_last_name" field.
func PreferredLastNameLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldPreferredLastName, v))
}

// PreferredLastNameLTE applies the LTE predicate on the "preferred_last_name" field.
func PreferredLastNameLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldPreferredLastName, v))
}

// PreferredLastNameContains applies the Contains predicate on the "preferred_last_name" field.
func PreferredLastNameContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldPreferredLastName, v))
}

// PreferredLastNameHasPrefix applies the HasPrefix predicate on the "preferred_last_name" field.
func PreferredLastNameHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldPreferredLastName, v))
}

// PreferredLastNameHasSuffix applies the HasSuffix predicate on the "preferred_last_name" field.
func PreferredLastNameHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldPreferredLastName, v))
}

// PreferredLastNameIsNil applies the IsNil predicate on the "preferred_last_name" field.
func PreferredLastNameIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldPreferredLastName))
}

// PreferredLastNameNotNil applies the NotNil predicate on the "preferred_last_name" field.
func PreferredLastNameNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldPreferredLastName))
}

// PreferredLastNameEqualFold applies the EqualFold predicate on the "preferred_last_name" field.
func PreferredLastNameEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldPreferredLastName, v))
}

// PreferredLastNameContainsFold applies the ContainsFold predicate on the "preferred_last_name" field.
func PreferredLastNameContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldPreferredLastName, v))
}

// JobTitleEQ applies the EQ predicate on the "job_title" field.
func JobTitleEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldJobTitle, v))
}

// JobTitleNEQ applies the NEQ predicate on the "job_title" field.
func JobTitleNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldJobTitle, v))
}

// JobTitleIn applies the In predicate on the "job_title" field.
func JobTitleIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldJobTitle, vs...))
}

// JobTitleNotIn applies the NotIn predicate on the "job_title" field.
func JobTitleNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldJobTitle, vs...))
}

// JobTitleGT applies the GT predicate on the "job_title" field.
func JobTitleGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldJobTitle, v))
}

// JobTitleGTE applies the GTE predicate on the "job_title" field.
func JobTitleGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldJobTitle, v))
}

// JobTitleLT applies the LT predicate on the "job_title" field.
func JobTitleLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldJobTitle, v))
}

// JobTitleLTE applies the LTE predicate on the "job_title" field.
func JobTitleLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldJobTitle, v))
}

// JobTitleContains applies the Contains predicate on the "job_title" field.
func JobTitleContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldJobTitle, v))
}

// JobTitleHasPrefix applies the HasPrefix predicate on the "job_title" field.
func JobTitleHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldJobTitle, v))
}

// JobTitleHasSuffix applies the HasSuffix predicate on the "job_title" field.
func JobTitleHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldJobTitle, v))
}

// JobTitleIsNil applies the IsNil predicate on the "job_title" field.
func JobTitleIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldJobTitle))
}

// JobTitleNotNil applies the NotNil predicate on the "job_title" field.
func JobTitleNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldJobTitle))
}

// JobTitleEqualFold applies the EqualFold predicate on the "job_title" field.
func JobTitleEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldJobTitle, v))
}

// JobTitleContainsFold applies the ContainsFold predicate on the "job_title" field.
func JobTitleContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldJobTitle, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Person) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Person) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
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
func Not(p predicate.Person) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		p(s.Not())
	})
}
