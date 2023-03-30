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
	Upsert(context.Context, *Person) (*Person, error)
	Get(context.Context, string) (*Person, error)
	Delete(context.Context, string) error
	Each(context.Context, func(*Person) bool) error
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

func (ps *personService) Upsert(ctx context.Context, p *Person) (*Person, error) {
	_, err := ps.db.Person.Query().Where(person.IDEQ(p.ID)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return ps.Create(ctx, p)
		}
		return nil, err
	}
	return ps.Update(ctx, p)
}

func (ps *personService) Create(ctx context.Context, p *Person) (*Person, error) {
	// date fields filled by schema
	t := ps.db.Person.Create()

	// keep in order; copy to Update if it changes
	t.SetActive(p.Active)
	t.SetBirthDate(p.BirthDate)
	t.SetJobCategory(p.JobCategory)
	t.SetEmail(p.Email)
	t.SetFirstName(p.FirstName)
	t.SetFullName(p.FullName)
	t.SetID(p.ID) // TODO: nil value overriden by entgo default function?
	t.SetJobTitle(p.JobTitle)
	t.SetLastName(p.LastName)
	t.SetOrcid(p.Orcid)
	t.SetOrganizationID(p.OrganizationID)
	t.SetOrcidToken(p.OrcidToken)
	t.SetOtherID(p.OtherID)
	t.SetPreferredFirstName(p.PreferredFirstName)
	t.SetPreferredLastName(p.PreferedLastName)

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
	t.SetJobCategory(p.JobCategory)
	t.SetEmail(p.Email)
	t.SetFirstName(p.FirstName)
	t.SetFullName(p.FullName)
	t.SetJobTitle(p.JobTitle)
	t.SetLastName(p.LastName)
	t.SetOrcid(p.Orcid)
	t.SetOrcidToken(p.OrcidToken)
	t.SetOrganizationID(p.OrganizationID)
	t.SetOtherID(p.OtherID)
	t.SetPreferredFirstName(p.PreferredFirstName)
	t.SetPreferredLastName(p.PreferedLastName)

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

func (ps *personService) Each(ctx context.Context, cb func(*Person) bool) error {

	// TODO: find a better way to do this (no cursors possible)
	var offset int = 0
	var limit int = 500
	for {
		rows, err := ps.db.Person.Query().Offset(offset).Limit(limit).Order(ent.Asc(person.FieldDateCreated)).All(ctx)
		if err != nil {
			return err
		}
		// entgo returns no error on empty results
		if len(rows) == 0 {
			break
		}
		for _, row := range rows {
			if !cb(personUnwrap(row)) {
				return nil
			}
		}
		offset += limit
	}

	return nil
}

func personUnwrap(e *ent.Person) *Person {
	p := &Person{
		Active:             e.Active,
		BirthDate:          e.BirthDate,
		DateCreated:        &e.DateCreated,
		DateUpdated:        &e.DateUpdated,
		Email:              e.Email,
		OtherID:            e.OtherID,
		FirstName:          e.FirstName,
		FullName:           e.FullName,
		ID:                 e.ID,
		LastName:           e.LastName,
		JobCategory:        e.JobCategory,
		Orcid:              e.Orcid,
		OrcidToken:         e.OrcidToken,
		OrganizationID:     e.OrganizationID,
		PreferedLastName:   e.PreferredLastName,
		PreferredFirstName: e.PreferredFirstName,
		JobTitle:           e.JobTitle,
	}
	return p
}
