// Code generated by ent, DO NOT EDIT.

package person

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/old-people-service/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Person {
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

// PublicID applies equality check predicate on the "public_id" field. It's identical to PublicIDEQ.
func PublicID(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldPublicID, v))
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

// GivenName applies equality check predicate on the "given_name" field. It's identical to GivenNameEQ.
func GivenName(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldGivenName, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldName, v))
}

// FamilyName applies equality check predicate on the "family_name" field. It's identical to FamilyNameEQ.
func FamilyName(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldFamilyName, v))
}

// PreferredGivenName applies equality check predicate on the "preferred_given_name" field. It's identical to PreferredGivenNameEQ.
func PreferredGivenName(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldPreferredGivenName, v))
}

// PreferredFamilyName applies equality check predicate on the "preferred_family_name" field. It's identical to PreferredFamilyNameEQ.
func PreferredFamilyName(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldPreferredFamilyName, v))
}

// HonorificPrefix applies equality check predicate on the "honorific_prefix" field. It's identical to HonorificPrefixEQ.
func HonorificPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldHonorificPrefix, v))
}

// ExpirationDate applies equality check predicate on the "expiration_date" field. It's identical to ExpirationDateEQ.
func ExpirationDate(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldExpirationDate, v))
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

// PublicIDEQ applies the EQ predicate on the "public_id" field.
func PublicIDEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldPublicID, v))
}

// PublicIDNEQ applies the NEQ predicate on the "public_id" field.
func PublicIDNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldPublicID, v))
}

// PublicIDIn applies the In predicate on the "public_id" field.
func PublicIDIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldPublicID, vs...))
}

// PublicIDNotIn applies the NotIn predicate on the "public_id" field.
func PublicIDNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldPublicID, vs...))
}

// PublicIDGT applies the GT predicate on the "public_id" field.
func PublicIDGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldPublicID, v))
}

// PublicIDGTE applies the GTE predicate on the "public_id" field.
func PublicIDGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldPublicID, v))
}

// PublicIDLT applies the LT predicate on the "public_id" field.
func PublicIDLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldPublicID, v))
}

// PublicIDLTE applies the LTE predicate on the "public_id" field.
func PublicIDLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldPublicID, v))
}

// PublicIDContains applies the Contains predicate on the "public_id" field.
func PublicIDContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldPublicID, v))
}

// PublicIDHasPrefix applies the HasPrefix predicate on the "public_id" field.
func PublicIDHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldPublicID, v))
}

// PublicIDHasSuffix applies the HasSuffix predicate on the "public_id" field.
func PublicIDHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldPublicID, v))
}

// PublicIDEqualFold applies the EqualFold predicate on the "public_id" field.
func PublicIDEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldPublicID, v))
}

// PublicIDContainsFold applies the ContainsFold predicate on the "public_id" field.
func PublicIDContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldPublicID, v))
}

// IdentifierIsNil applies the IsNil predicate on the "identifier" field.
func IdentifierIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldIdentifier))
}

// IdentifierNotNil applies the NotNil predicate on the "identifier" field.
func IdentifierNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldIdentifier))
}

// IdentifierValuesIsNil applies the IsNil predicate on the "identifier_values" field.
func IdentifierValuesIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldIdentifierValues))
}

// IdentifierValuesNotNil applies the NotNil predicate on the "identifier_values" field.
func IdentifierValuesNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldIdentifierValues))
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

// GivenNameEQ applies the EQ predicate on the "given_name" field.
func GivenNameEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldGivenName, v))
}

// GivenNameNEQ applies the NEQ predicate on the "given_name" field.
func GivenNameNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldGivenName, v))
}

// GivenNameIn applies the In predicate on the "given_name" field.
func GivenNameIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldGivenName, vs...))
}

// GivenNameNotIn applies the NotIn predicate on the "given_name" field.
func GivenNameNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldGivenName, vs...))
}

// GivenNameGT applies the GT predicate on the "given_name" field.
func GivenNameGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldGivenName, v))
}

// GivenNameGTE applies the GTE predicate on the "given_name" field.
func GivenNameGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldGivenName, v))
}

// GivenNameLT applies the LT predicate on the "given_name" field.
func GivenNameLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldGivenName, v))
}

// GivenNameLTE applies the LTE predicate on the "given_name" field.
func GivenNameLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldGivenName, v))
}

// GivenNameContains applies the Contains predicate on the "given_name" field.
func GivenNameContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldGivenName, v))
}

// GivenNameHasPrefix applies the HasPrefix predicate on the "given_name" field.
func GivenNameHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldGivenName, v))
}

// GivenNameHasSuffix applies the HasSuffix predicate on the "given_name" field.
func GivenNameHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldGivenName, v))
}

// GivenNameIsNil applies the IsNil predicate on the "given_name" field.
func GivenNameIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldGivenName))
}

// GivenNameNotNil applies the NotNil predicate on the "given_name" field.
func GivenNameNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldGivenName))
}

// GivenNameEqualFold applies the EqualFold predicate on the "given_name" field.
func GivenNameEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldGivenName, v))
}

// GivenNameContainsFold applies the ContainsFold predicate on the "given_name" field.
func GivenNameContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldGivenName, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldName, v))
}

// FamilyNameEQ applies the EQ predicate on the "family_name" field.
func FamilyNameEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldFamilyName, v))
}

// FamilyNameNEQ applies the NEQ predicate on the "family_name" field.
func FamilyNameNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldFamilyName, v))
}

// FamilyNameIn applies the In predicate on the "family_name" field.
func FamilyNameIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldFamilyName, vs...))
}

// FamilyNameNotIn applies the NotIn predicate on the "family_name" field.
func FamilyNameNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldFamilyName, vs...))
}

// FamilyNameGT applies the GT predicate on the "family_name" field.
func FamilyNameGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldFamilyName, v))
}

// FamilyNameGTE applies the GTE predicate on the "family_name" field.
func FamilyNameGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldFamilyName, v))
}

// FamilyNameLT applies the LT predicate on the "family_name" field.
func FamilyNameLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldFamilyName, v))
}

// FamilyNameLTE applies the LTE predicate on the "family_name" field.
func FamilyNameLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldFamilyName, v))
}

// FamilyNameContains applies the Contains predicate on the "family_name" field.
func FamilyNameContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldFamilyName, v))
}

// FamilyNameHasPrefix applies the HasPrefix predicate on the "family_name" field.
func FamilyNameHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldFamilyName, v))
}

// FamilyNameHasSuffix applies the HasSuffix predicate on the "family_name" field.
func FamilyNameHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldFamilyName, v))
}

// FamilyNameIsNil applies the IsNil predicate on the "family_name" field.
func FamilyNameIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldFamilyName))
}

// FamilyNameNotNil applies the NotNil predicate on the "family_name" field.
func FamilyNameNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldFamilyName))
}

// FamilyNameEqualFold applies the EqualFold predicate on the "family_name" field.
func FamilyNameEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldFamilyName, v))
}

// FamilyNameContainsFold applies the ContainsFold predicate on the "family_name" field.
func FamilyNameContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldFamilyName, v))
}

// JobCategoryIsNil applies the IsNil predicate on the "job_category" field.
func JobCategoryIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldJobCategory))
}

// JobCategoryNotNil applies the NotNil predicate on the "job_category" field.
func JobCategoryNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldJobCategory))
}

// PreferredGivenNameEQ applies the EQ predicate on the "preferred_given_name" field.
func PreferredGivenNameEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldPreferredGivenName, v))
}

// PreferredGivenNameNEQ applies the NEQ predicate on the "preferred_given_name" field.
func PreferredGivenNameNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldPreferredGivenName, v))
}

// PreferredGivenNameIn applies the In predicate on the "preferred_given_name" field.
func PreferredGivenNameIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldPreferredGivenName, vs...))
}

// PreferredGivenNameNotIn applies the NotIn predicate on the "preferred_given_name" field.
func PreferredGivenNameNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldPreferredGivenName, vs...))
}

// PreferredGivenNameGT applies the GT predicate on the "preferred_given_name" field.
func PreferredGivenNameGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldPreferredGivenName, v))
}

// PreferredGivenNameGTE applies the GTE predicate on the "preferred_given_name" field.
func PreferredGivenNameGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldPreferredGivenName, v))
}

// PreferredGivenNameLT applies the LT predicate on the "preferred_given_name" field.
func PreferredGivenNameLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldPreferredGivenName, v))
}

// PreferredGivenNameLTE applies the LTE predicate on the "preferred_given_name" field.
func PreferredGivenNameLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldPreferredGivenName, v))
}

// PreferredGivenNameContains applies the Contains predicate on the "preferred_given_name" field.
func PreferredGivenNameContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldPreferredGivenName, v))
}

// PreferredGivenNameHasPrefix applies the HasPrefix predicate on the "preferred_given_name" field.
func PreferredGivenNameHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldPreferredGivenName, v))
}

// PreferredGivenNameHasSuffix applies the HasSuffix predicate on the "preferred_given_name" field.
func PreferredGivenNameHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldPreferredGivenName, v))
}

// PreferredGivenNameIsNil applies the IsNil predicate on the "preferred_given_name" field.
func PreferredGivenNameIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldPreferredGivenName))
}

// PreferredGivenNameNotNil applies the NotNil predicate on the "preferred_given_name" field.
func PreferredGivenNameNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldPreferredGivenName))
}

// PreferredGivenNameEqualFold applies the EqualFold predicate on the "preferred_given_name" field.
func PreferredGivenNameEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldPreferredGivenName, v))
}

// PreferredGivenNameContainsFold applies the ContainsFold predicate on the "preferred_given_name" field.
func PreferredGivenNameContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldPreferredGivenName, v))
}

// PreferredFamilyNameEQ applies the EQ predicate on the "preferred_family_name" field.
func PreferredFamilyNameEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldPreferredFamilyName, v))
}

// PreferredFamilyNameNEQ applies the NEQ predicate on the "preferred_family_name" field.
func PreferredFamilyNameNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldPreferredFamilyName, v))
}

// PreferredFamilyNameIn applies the In predicate on the "preferred_family_name" field.
func PreferredFamilyNameIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldPreferredFamilyName, vs...))
}

// PreferredFamilyNameNotIn applies the NotIn predicate on the "preferred_family_name" field.
func PreferredFamilyNameNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldPreferredFamilyName, vs...))
}

// PreferredFamilyNameGT applies the GT predicate on the "preferred_family_name" field.
func PreferredFamilyNameGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldPreferredFamilyName, v))
}

// PreferredFamilyNameGTE applies the GTE predicate on the "preferred_family_name" field.
func PreferredFamilyNameGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldPreferredFamilyName, v))
}

// PreferredFamilyNameLT applies the LT predicate on the "preferred_family_name" field.
func PreferredFamilyNameLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldPreferredFamilyName, v))
}

// PreferredFamilyNameLTE applies the LTE predicate on the "preferred_family_name" field.
func PreferredFamilyNameLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldPreferredFamilyName, v))
}

// PreferredFamilyNameContains applies the Contains predicate on the "preferred_family_name" field.
func PreferredFamilyNameContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldPreferredFamilyName, v))
}

// PreferredFamilyNameHasPrefix applies the HasPrefix predicate on the "preferred_family_name" field.
func PreferredFamilyNameHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldPreferredFamilyName, v))
}

// PreferredFamilyNameHasSuffix applies the HasSuffix predicate on the "preferred_family_name" field.
func PreferredFamilyNameHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldPreferredFamilyName, v))
}

// PreferredFamilyNameIsNil applies the IsNil predicate on the "preferred_family_name" field.
func PreferredFamilyNameIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldPreferredFamilyName))
}

// PreferredFamilyNameNotNil applies the NotNil predicate on the "preferred_family_name" field.
func PreferredFamilyNameNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldPreferredFamilyName))
}

// PreferredFamilyNameEqualFold applies the EqualFold predicate on the "preferred_family_name" field.
func PreferredFamilyNameEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldPreferredFamilyName, v))
}

// PreferredFamilyNameContainsFold applies the ContainsFold predicate on the "preferred_family_name" field.
func PreferredFamilyNameContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldPreferredFamilyName, v))
}

// HonorificPrefixEQ applies the EQ predicate on the "honorific_prefix" field.
func HonorificPrefixEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldHonorificPrefix, v))
}

// HonorificPrefixNEQ applies the NEQ predicate on the "honorific_prefix" field.
func HonorificPrefixNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldHonorificPrefix, v))
}

// HonorificPrefixIn applies the In predicate on the "honorific_prefix" field.
func HonorificPrefixIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldHonorificPrefix, vs...))
}

// HonorificPrefixNotIn applies the NotIn predicate on the "honorific_prefix" field.
func HonorificPrefixNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldHonorificPrefix, vs...))
}

// HonorificPrefixGT applies the GT predicate on the "honorific_prefix" field.
func HonorificPrefixGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldHonorificPrefix, v))
}

// HonorificPrefixGTE applies the GTE predicate on the "honorific_prefix" field.
func HonorificPrefixGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldHonorificPrefix, v))
}

// HonorificPrefixLT applies the LT predicate on the "honorific_prefix" field.
func HonorificPrefixLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldHonorificPrefix, v))
}

// HonorificPrefixLTE applies the LTE predicate on the "honorific_prefix" field.
func HonorificPrefixLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldHonorificPrefix, v))
}

// HonorificPrefixContains applies the Contains predicate on the "honorific_prefix" field.
func HonorificPrefixContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldHonorificPrefix, v))
}

// HonorificPrefixHasPrefix applies the HasPrefix predicate on the "honorific_prefix" field.
func HonorificPrefixHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldHonorificPrefix, v))
}

// HonorificPrefixHasSuffix applies the HasSuffix predicate on the "honorific_prefix" field.
func HonorificPrefixHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldHonorificPrefix, v))
}

// HonorificPrefixIsNil applies the IsNil predicate on the "honorific_prefix" field.
func HonorificPrefixIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldHonorificPrefix))
}

// HonorificPrefixNotNil applies the NotNil predicate on the "honorific_prefix" field.
func HonorificPrefixNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldHonorificPrefix))
}

// HonorificPrefixEqualFold applies the EqualFold predicate on the "honorific_prefix" field.
func HonorificPrefixEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldHonorificPrefix, v))
}

// HonorificPrefixContainsFold applies the ContainsFold predicate on the "honorific_prefix" field.
func HonorificPrefixContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldHonorificPrefix, v))
}

// RoleIsNil applies the IsNil predicate on the "role" field.
func RoleIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldRole))
}

// RoleNotNil applies the NotNil predicate on the "role" field.
func RoleNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldRole))
}

// SettingsIsNil applies the IsNil predicate on the "settings" field.
func SettingsIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldSettings))
}

// SettingsNotNil applies the NotNil predicate on the "settings" field.
func SettingsNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldSettings))
}

// ObjectClassIsNil applies the IsNil predicate on the "object_class" field.
func ObjectClassIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldObjectClass))
}

// ObjectClassNotNil applies the NotNil predicate on the "object_class" field.
func ObjectClassNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldObjectClass))
}

// ExpirationDateEQ applies the EQ predicate on the "expiration_date" field.
func ExpirationDateEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldExpirationDate, v))
}

// ExpirationDateNEQ applies the NEQ predicate on the "expiration_date" field.
func ExpirationDateNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldExpirationDate, v))
}

// ExpirationDateIn applies the In predicate on the "expiration_date" field.
func ExpirationDateIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldExpirationDate, vs...))
}

// ExpirationDateNotIn applies the NotIn predicate on the "expiration_date" field.
func ExpirationDateNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldExpirationDate, vs...))
}

// ExpirationDateGT applies the GT predicate on the "expiration_date" field.
func ExpirationDateGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldExpirationDate, v))
}

// ExpirationDateGTE applies the GTE predicate on the "expiration_date" field.
func ExpirationDateGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldExpirationDate, v))
}

// ExpirationDateLT applies the LT predicate on the "expiration_date" field.
func ExpirationDateLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldExpirationDate, v))
}

// ExpirationDateLTE applies the LTE predicate on the "expiration_date" field.
func ExpirationDateLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldExpirationDate, v))
}

// ExpirationDateContains applies the Contains predicate on the "expiration_date" field.
func ExpirationDateContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldExpirationDate, v))
}

// ExpirationDateHasPrefix applies the HasPrefix predicate on the "expiration_date" field.
func ExpirationDateHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldExpirationDate, v))
}

// ExpirationDateHasSuffix applies the HasSuffix predicate on the "expiration_date" field.
func ExpirationDateHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldExpirationDate, v))
}

// ExpirationDateIsNil applies the IsNil predicate on the "expiration_date" field.
func ExpirationDateIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldExpirationDate))
}

// ExpirationDateNotNil applies the NotNil predicate on the "expiration_date" field.
func ExpirationDateNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldExpirationDate))
}

// ExpirationDateEqualFold applies the EqualFold predicate on the "expiration_date" field.
func ExpirationDateEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldExpirationDate, v))
}

// ExpirationDateContainsFold applies the ContainsFold predicate on the "expiration_date" field.
func ExpirationDateContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldExpirationDate, v))
}

// TokenIsNil applies the IsNil predicate on the "token" field.
func TokenIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldToken))
}

// TokenNotNil applies the NotNil predicate on the "token" field.
func TokenNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldToken))
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
