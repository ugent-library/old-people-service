package models

import (
	"context"
	"database/sql"
	"errors"

	entdialect "entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/ugent-library/people/ent"
	entmigrate "github.com/ugent-library/people/ent/migrate"
	"github.com/ugent-library/people/ent/person"
)

var ErrNotFound = errors.New("not found")

type PersonConfig struct {
	DB string
}

type PersonService interface {
	Create(context.Context, *Person) (*Person, error)
	Update(context.Context, *Person) (*Person, error)
	Get(context.Context, string) (*Person, error)
	Delete(context.Context, string) error
}

type personService struct {
	db *ent.Client
}

func NewPersonService(cfg *PersonConfig) (PersonService, error) {
	db, err := sql.Open("pgx", cfg.DB)
	if err != nil {
		return nil, err
	}

	driver := entsql.OpenDB(entdialect.Postgres, db)
	client := ent.NewClient(ent.Driver(driver))

	err = client.Schema.Create(context.Background(),
		entmigrate.WithDropIndex(true),
	)
	if err != nil {
		return nil, err
	}

	return &personService{
		db: client,
	}, nil
}

func (ps *personService) Create(ctx context.Context, p *Person) (*Person, error) {
	// date fields filled by schema
	t := ps.db.Person.Create()

	// keep in order; copy to Update if it changes
	t.SetActive(p.Active)
	t.SetBirthDate(p.BirthDate)
	if p.DateLastLogin != nil {
		t.SetDateLastLogin(*p.DateLastLogin)
	}
	t.SetDeleted(p.Deleted)
	t.SetDormCountry(p.DormCountry)
	t.SetDormLocality(p.DormLocality)
	t.SetDormPostalCode(p.DormPostalCode)
	t.SetDormStreetAddress(p.DormStreetAddress)
	t.SetEmail(p.Email)
	t.SetFirstName(p.FirstName)
	t.SetHomeCountry(p.HomeCountry)
	t.SetHomeLocality(p.HomeLocality)
	t.SetHomePostalCode(p.HomePostalCode)
	t.SetHomeStreetAddress(p.HomeStreetAddress)
	t.SetHomeTel(p.HomeTel)
	t.SetID(p.ID) // TODO: nil value overriden by entgo default function?
	t.SetLastName(p.LastName)
	t.SetMiddleName(p.MiddleName)
	t.SetNationality(p.Nationality)
	t.SetObjectClass(p.ObjectClass)
	t.SetOrcidBio(p.OrcidBio)
	t.SetOrcidID(p.OrcidID)
	if p.OrcidSettings != nil {
		t.SetOrcidSettings(*p.OrcidSettings)
	}
	t.SetOrcidToken(p.OrcidToken)
	if p.OrcidVerify != nil {
		t.SetOrcidVerify(*p.OrcidVerify)
	}
	t.SetPreferredFirstName(p.PreferredFirstName)
	t.SetPreferredLastName(p.PreferedLastName)
	t.SetPublicationCount(p.PublicationCount)
	t.SetReplacedBy(fromPersonRefs(p.ReplacedBy))
	t.SetReplaces(fromPersonRefs(p.Replaces))
	t.SetResearchDiscipline(p.ResearchDiscipline)
	t.SetResearchDisciplineCode(p.ResearchDisciplineCode)
	t.SetRoles(p.Roles)
	if p.Settings != nil {
		t.SetSettings(*p.Settings)
	}
	t.SetTitle(p.Title)
	t.SetUgentAppointmentDate(p.UgentAppointmentDate)
	t.SetUgentBarcode(p.UgentBarcode)
	t.SetUgentCampus(p.UgentCampus)
	t.SetUgentDepartmentID(p.UgentDepartmentID)
	t.SetUgentDepartmentName(p.UgentDepartmentName)
	t.SetUgentExpirationDate(p.UgentExpirationDate)
	t.SetUgentExtCategory(p.UgentExtCategory)
	t.SetUgentFacultyID(p.UgentFacultyID)
	t.SetUgentID(p.UgentID)
	t.SetUgentJobCategory(p.UgentJobCategory)
	t.SetUgentJobTitle(p.UgentJobTitle)
	t.SetUgentLastEnrolled(p.UgentLastEnrolled)
	t.SetUgentLocality(p.UgentLocality)
	t.SetUgentMemorialisID(p.UgentMemorialisID)
	t.SetUgentPostalCode(p.UgentPostalCode)
	t.SetUgentStreetAddress(p.UgentStreetAddress)
	t.SetUgentTel(p.UgentTel)
	t.SetUgentUsername(p.UgentUsername)
	t.SetUzgentDepartmentName(p.UzgentDepartmentName)
	t.SetUzgentID(p.UzgentID)
	t.SetUzgentJobTitle(p.UzgentJobTitle)

	row, err := t.Save(ctx)
	if err != nil {
		return nil, err
	}

	// collect entgo managed fields
	p.DateCreated = &row.DateCreated
	p.DateUpdated = &row.DateUpdated
	p.ID = row.ID

	return p, nil
}

func (ps *personService) Update(ctx context.Context, p *Person) (*Person, error) {
	t := ps.db.Person.UpdateOneID(p.ID)

	// keep in order; copy to Update if it changes
	t.SetActive(p.Active)
	t.SetBirthDate(p.BirthDate)
	if p.DateLastLogin != nil {
		t.SetDateLastLogin(*p.DateLastLogin)
	}
	t.SetDeleted(p.Deleted)
	t.SetDormCountry(p.DormCountry)
	t.SetDormLocality(p.DormLocality)
	t.SetDormPostalCode(p.DormPostalCode)
	t.SetDormStreetAddress(p.DormStreetAddress)
	t.SetEmail(p.Email)
	t.SetFirstName(p.FirstName)
	t.SetHomeCountry(p.HomeCountry)
	t.SetHomeLocality(p.HomeLocality)
	t.SetHomePostalCode(p.HomePostalCode)
	t.SetHomeStreetAddress(p.HomeStreetAddress)
	t.SetHomeTel(p.HomeTel)
	//t.SetID(p.ID) // TODO: nil value overriden by entgo default function?
	t.SetLastName(p.LastName)
	t.SetMiddleName(p.MiddleName)
	t.SetNationality(p.Nationality)
	t.SetObjectClass(p.ObjectClass)
	t.SetOrcidBio(p.OrcidBio)
	t.SetOrcidID(p.OrcidID)
	if p.OrcidSettings != nil {
		t.SetOrcidSettings(*p.OrcidSettings)
	}
	t.SetOrcidToken(p.OrcidToken)
	if p.OrcidVerify != nil {
		t.SetOrcidVerify(*p.OrcidVerify)
	}
	t.SetPreferredFirstName(p.PreferredFirstName)
	t.SetPreferredLastName(p.PreferedLastName)
	t.SetPublicationCount(p.PublicationCount)
	t.SetReplacedBy(fromPersonRefs(p.ReplacedBy))
	t.SetReplaces(fromPersonRefs(p.Replaces))
	t.SetResearchDiscipline(p.ResearchDiscipline)
	t.SetResearchDisciplineCode(p.ResearchDisciplineCode)
	t.SetRoles(p.Roles)
	if p.Settings != nil {
		t.SetSettings(*p.Settings)
	}
	t.SetTitle(p.Title)
	t.SetUgentAppointmentDate(p.UgentAppointmentDate)
	t.SetUgentBarcode(p.UgentBarcode)
	t.SetUgentCampus(p.UgentCampus)
	t.SetUgentDepartmentID(p.UgentDepartmentID)
	t.SetUgentDepartmentName(p.UgentDepartmentName)
	t.SetUgentExpirationDate(p.UgentExpirationDate)
	t.SetUgentExtCategory(p.UgentExtCategory)
	t.SetUgentFacultyID(p.UgentFacultyID)
	t.SetUgentID(p.UgentID)
	t.SetUgentJobCategory(p.UgentJobCategory)
	t.SetUgentJobTitle(p.UgentJobTitle)
	t.SetUgentLastEnrolled(p.UgentLastEnrolled)
	t.SetUgentLocality(p.UgentLocality)
	t.SetUgentMemorialisID(p.UgentMemorialisID)
	t.SetUgentPostalCode(p.UgentPostalCode)
	t.SetUgentStreetAddress(p.UgentStreetAddress)
	t.SetUgentTel(p.UgentTel)
	t.SetUgentUsername(p.UgentUsername)
	t.SetUzgentDepartmentName(p.UzgentDepartmentName)
	t.SetUzgentID(p.UzgentID)
	t.SetUzgentJobTitle(p.UzgentJobTitle)

	row, err := t.Save(ctx)
	if err != nil {
		return nil, err
	}

	// collect entgo managed fields (ID and DateCreated are supposed to preexist)
	p.DateUpdated = &row.DateUpdated

	return p, nil
}

func (ps *personService) Get(ctx context.Context, id string) (*Person, error) {
	row, err := ps.db.Person.Query().Where(person.IDEQ(id)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return personUnwrap(row), nil
}

func (ps *personService) Delete(ctx context.Context, id string) error {
	return ps.db.Person.DeleteOneID(id).Exec(ctx)
}

func personUnwrap(e *ent.Person) *Person {
	p := &Person{
		Active:                 e.Active,
		BirthDate:              e.BirthDate,
		DateCreated:            &e.DateCreated,
		DateLastLogin:          &e.DateLastLogin,
		DateUpdated:            &e.DateUpdated,
		Deleted:                e.Deleted,
		DormCountry:            e.DormCountry,
		DormLocality:           e.DormLocality,
		DormPostalCode:         e.DormPostalCode,
		DormStreetAddress:      e.DormStreetAddress,
		Email:                  e.Email,
		FirstName:              e.FirstName,
		HomeCountry:            e.HomeCountry,
		HomeLocality:           e.HomeLocality,
		HomePostalCode:         e.HomePostalCode,
		HomeStreetAddress:      e.HomeStreetAddress,
		HomeTel:                e.HomeTel,
		ID:                     e.ID,
		LastName:               e.LastName,
		MiddleName:             e.MiddleName,
		Nationality:            e.Nationality,
		ObjectClass:            e.ObjectClass,
		OrcidBio:               e.OrcidBio,
		OrcidID:                e.OrcidID,
		OrcidSettings:          &e.OrcidSettings,
		OrcidToken:             e.OrcidToken,
		OrcidVerify:            &e.OrcidVerify,
		PreferedLastName:       e.PreferredLastName,
		PreferredFirstName:     e.PreferredFirstName,
		PublicationCount:       e.PublicationCount,
		ReplacedBy:             toPersonRefs(e.ReplacedBy),
		Replaces:               toPersonRefs(e.Replaces),
		ResearchDiscipline:     e.ResearchDiscipline,
		ResearchDisciplineCode: e.ResearchDisciplineCode,
		Roles:                  e.Roles,
		Settings:               &e.Settings,
		Title:                  e.Title,
		UgentAppointmentDate:   e.UgentAppointmentDate,
		UgentBarcode:           e.UgentBarcode,
		UgentCampus:            e.UgentCampus,
		UgentDepartmentID:      e.UgentDepartmentID,
		UgentDepartmentName:    e.UgentDepartmentName,
		UgentExpirationDate:    e.UgentExpirationDate,
		UgentExtCategory:       e.UgentExtCategory,
		UgentFacultyID:         e.UgentFacultyID,
		UgentID:                e.UgentID,
		UgentJobCategory:       e.UgentJobCategory,
		UgentJobTitle:          e.UgentJobTitle,
		UgentLastEnrolled:      e.UgentLastEnrolled,
		UgentLocality:          e.UgentLocality,
		UgentMemorialisID:      e.UgentMemorialisID,
		UgentPostalCode:        e.UgentPostalCode,
		UgentStreetAddress:     e.UgentStreetAddress,
		UgentTel:               e.UgentTel,
		UgentUsername:          e.UgentUsername,
		UzgentDepartmentName:   e.UzgentDepartmentName,
		UzgentID:               e.UzgentID,
		UzgentJobTitle:         e.UzgentJobTitle,
	}
	return p
}

func toPersonRefs(m []map[string]string) []*PersonRef {
	refs := make([]*PersonRef, 0, len(m))
	for _, mr := range m {
		refs = append(refs, &PersonRef{ID: mr["id"]})
	}
	return refs
}

func fromPersonRefs(refs []*PersonRef) []map[string]string {
	m := make([]map[string]string, 0, len(refs))
	for _, ref := range refs {
		m = append(m, map[string]string{"id": ref.ID})
	}
	return m
}
