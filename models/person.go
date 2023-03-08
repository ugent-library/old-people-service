package models

import (
	"time"

	"github.com/ugent-library/people/ent/schema"
)

type PersonRef struct {
	ID string `json:"id"`
}

type Person struct {
	Active                 bool                  `json:"active"`
	BirthDate              string                `json:"birth_date,omitempty"`
	DateCreated            *time.Time            `json:"date_created"`
	DateLastLogin          *time.Time            `json:"date_last_login,omitempty"`
	DateUpdated            *time.Time            `json:"date_updated"`
	Deleted                bool                  `json:"deleted"`
	DormCountry            string                `json:"dorm_country,omitempty"`
	DormLocality           string                `json:"dorm_locality,omitempty"`
	DormPostalCode         string                `json:"dorm_postal_code,omitempty"`
	DormStreetAddress      string                `json:"dorm_street_address,omitempty"`
	Email                  string                `json:"email,omitempty"`
	FirstName              string                `json:"first_name,omitempty"`
	HomeCountry            string                `json:"home_country,omitempty"`
	HomeLocality           string                `json:"home_locality,omitempty"`
	HomePostalCode         string                `json:"home_postal_code,omitempty"`
	HomeStreetAddress      string                `json:"home_street_address,omitempty"`
	HomeTel                string                `json:"home_tel,omitempty"`
	ID                     string                `json:"id,omitempty"`
	LastName               string                `json:"last_name,omitempty"`
	MiddleName             string                `json:"middle_name,omitempty"`
	Nationality            string                `json:"nationality,omitempty"`
	ObjectClass            []string              `json:"object_class,omitempty"`
	OrcidBio               string                `json:"orcid_bio,omitempty"`
	OrcidID                string                `json:"orcid_id,omitempty"`
	OrcidSettings          *schema.OrcidSettings `json:"orcid_settings,omitempty"`
	OrcidToken             string                `json:"orcid_token,omitempty"`
	OrcidVerify            *schema.OrcidVerify   `json:"orcid_verify,omitempty"`
	PreferedLastName       string                `json:"preferred_last_name,omitempty"`
	PreferredFirstName     string                `json:"preferred_first_name,omitempty"`
	PublicationCount       int                   `json:"publication_count"`
	ReplacedBy             []*PersonRef          `json:"replaced_by,omitempty"`
	Replaces               []*PersonRef          `json:"replaces,omitempty"`
	ResearchDiscipline     []string              `json:"research_discipline,omitempty"`
	ResearchDisciplineCode []string              `json:"research_discipline_code,omitempty"`
	Roles                  []string              `json:"roles,omitempty"`
	Settings               *schema.Settings      `json:"settings,omitempty"`
	Title                  string                `json:"title,omitempty"`
	UgentAppointmentDate   string                `json:"ugent_appointment_date,omitempty"`
	UgentBarcode           []string              `json:"ugent_barcode,omitempty"`
	UgentCampus            []string              `json:"ugent_campus,omitempty"`
	UgentDepartmentID      []string              `json:"ugent_department_id,omitempty"`
	UgentDepartmentName    []string              `json:"ugent_department_name,omitempty"`
	UgentExpirationDate    string                `json:"ugent_expiration_date,omitempty"`
	UgentExtCategory       []string              `json:"ugent_ext_category,omitempty"`
	UgentFacultyID         []string              `json:"ugent_faculty_id,omitempty"`
	UgentID                []string              `json:"ugent_id,omitempty"`
	UgentJobCategory       []string              `json:"ugent_job_category,omitempty"`
	UgentJobTitle          []string              `json:"ugent_job_title,omitempty"`
	UgentLastEnrolled      string                `json:"ugent_last_enrolled,omitempty"`
	UgentLocality          string                `json:"ugent_locality,omitempty"`
	UgentMemorialisID      string                `json:"ugent_memorialis_id,omitempty"`
	UgentPostalCode        string                `json:"ugent_postal_code,omitempty"`
	UgentStreetAddress     string                `json:"ugent_street_address,omitempty"`
	UgentTel               string                `json:"ugent_tel,omitempty"`
	UgentUsername          string                `json:"ugent_username,omitempty"`
	UzgentDepartmentName   []string              `json:"uzgent_department_name,omitempty"`
	UzgentID               []string              `json:"uzgent_id,omitempty"`
	UzgentJobTitle         []string              `json:"uzgent_job_title,omitempty"`
}
