package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"sort"

	entdialect "entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/ent"
	entmigrate "github.com/ugent-library/people/ent/migrate"
	"github.com/ugent-library/people/ent/organization"
	"github.com/ugent-library/people/ent/person"
	"github.com/ugent-library/people/ent/schema"
	"github.com/ugent-library/people/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PersonConfig struct {
	DB string
}

type personService struct {
	db *ent.Client
}

func NewPersonService(cfg *PersonConfig) (models.PersonService, error) {
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

func (ps *personService) Create(ctx context.Context, p *models.Person) (*models.Person, error) {
	// date fields filled by schema
	t := ps.db.Person.Create()

	// keep in order; copy to Update if it changes
	t.SetActive(p.Active)
	t.SetBirthDate(p.BirthDate)
	t.SetJobCategory(p.JobCategory)
	t.SetEmail(p.Email)
	t.SetFirstName(p.FirstName)
	t.SetFullName(p.FullName)
	t.SetPublicID(p.Id) // TODO: nil value overriden by entgo default function?
	t.SetTitle(p.Title)
	t.SetLastName(p.LastName)
	t.SetOrcid(p.Orcid)
	t.SetOrcidToken(p.OrcidToken)
	schemaOtherIds := make([]schema.IdRef, 0, len(p.OtherId))
	for _, refId := range p.OtherId {
		schemaOtherIds = append(schemaOtherIds, schema.IdRef{
			ID:   refId.Id,
			Type: refId.Type,
		})
	}
	t.SetOtherID(schemaOtherIds)
	t.SetPreferredFirstName(p.PreferredFirstName)
	t.SetPreferredLastName(p.PreferredLastName)

	// TODO: test
	if p.OrganizationId != nil && len(p.OrganizationId) > 0 {
		// TODO: crashes with segmentation violation error when org does not exist
		orgs, err := ps.db.Organization.Query().Where(organization.PublicIDIn(p.OrganizationId...)).All(ctx)
		if err != nil {
			return nil, err
		}
		// add missing organizations
		for _, orgId := range p.OrganizationId {
			found := false
			for _, org := range orgs {
				if org.PublicID == orgId {
					found = true
					break
				}
			}
			if !found {
				oc := ps.db.Organization.Create()
				oc.SetPublicID(orgId)
				oc.SetName(orgId)
				newOrg, err := oc.Save(ctx)
				if err != nil {
					return nil, err
				}
				orgs = append(orgs, newOrg)
			}
		}
		t.AddOrganizations(orgs...)
	}

	row, err := t.Save(ctx)
	if err != nil {
		return nil, err
	}

	// collect entgo managed fields
	p.DateCreated = timestamppb.New(row.DateCreated)
	p.DateUpdated = timestamppb.New(row.DateUpdated)
	p.Id = row.PublicID

	return p, nil
}

func (ps *personService) Update(ctx context.Context, p *models.Person) (*models.Person, error) {
	t := ps.db.Person.Update().Where(person.PublicIDEQ(p.Id))

	// keep in order; copy to Update if it changes
	t.SetActive(p.Active)
	t.SetBirthDate(p.BirthDate)
	t.SetJobCategory(p.JobCategory)
	t.SetEmail(p.Email)
	t.SetFirstName(p.FirstName)
	t.SetFullName(p.FullName)
	t.SetTitle(p.Title)
	t.SetLastName(p.LastName)
	t.SetOrcid(p.Orcid)
	t.SetOrcidToken(p.OrcidToken)
	schemaOtherIds := make([]schema.IdRef, 0, len(p.OtherId))
	for _, refId := range p.OtherId {
		schemaOtherIds = append(schemaOtherIds, schema.IdRef{
			ID:   refId.Id,
			Type: refId.Type,
		})
	}
	t.SetOtherID(schemaOtherIds)
	t.SetPreferredFirstName(p.PreferredFirstName)
	t.SetPreferredLastName(p.PreferredLastName)

	t.ClearOrganizations()
	if p.OrganizationId != nil && len(p.OrganizationId) > 0 {
		// TODO: crashes with segmentation violation error when org does not exist
		orgs, err := ps.db.Organization.Query().Where(organization.PublicIDIn(p.OrganizationId...)).All(ctx)
		if err != nil {
			return nil, err
		}
		// add missing organizations
		for _, orgId := range p.OrganizationId {
			found := false
			for _, org := range orgs {
				if org.PublicID == orgId {
					found = true
					break
				}
			}
			if !found {
				oc := ps.db.Organization.Create()
				oc.SetPublicID(orgId)
				oc.SetName(orgId)
				newOrg, err := oc.Save(ctx)
				if err != nil {
					return nil, err
				}
				orgs = append(orgs, newOrg)
			}
		}
		t.AddOrganizations(orgs...)
	}

	_, err := t.Save(ctx)
	if err != nil {
		return nil, err
	}

	return ps.Get(ctx, p.Id)
}

func (ps *personService) Get(ctx context.Context, id string) (*models.Person, error) {
	row, err := ps.db.Person.Query().WithOrganizations().Where(person.PublicIDEQ(id)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		fmt.Fprintf(os.Stderr, "return other error!!!!\n")
		return nil, err
	}
	return personUnwrap(row), nil
}

func (ps *personService) Delete(ctx context.Context, id string) error {
	_, err := ps.db.Person.Delete().Where(person.PublicIDEQ(id)).Exec(ctx)
	return err
}

func (ps *personService) Each(ctx context.Context, cb func(*models.Person) bool) error {

	// TODO: find a better way to do this (no cursors possible)
	var offset int = 0
	var limit int = 500
	for {
		rows, err := ps.db.Person.Query().WithOrganizations().Offset(offset).Limit(limit).Order(ent.Asc(person.FieldDateCreated)).All(ctx)
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

func personUnwrap(e *ent.Person) *models.Person {
	refIds := make([]*v1.IdRef, 0, len(e.OtherID))
	for _, schemaOtherId := range e.OtherID {
		refIds = append(refIds, &v1.IdRef{
			Id:   schemaOtherId.ID,
			Type: schemaOtherId.Type,
		})
	}
	orgIds := make([]string, 0)
	orgRows := e.Edges.Organizations
	sort.SliceStable(orgRows, func(i, j int) bool {
		return orgRows[i].DateCreated.Before(orgRows[j].DateCreated)
	})
	for _, org := range e.Edges.Organizations {
		orgIds = append(orgIds, org.PublicID)
	}
	p := &models.Person{
		Person: v1.Person{
			Active:             e.Active,
			BirthDate:          e.BirthDate,
			DateCreated:        timestamppb.New(e.DateCreated),
			DateUpdated:        timestamppb.New(e.DateUpdated),
			Email:              e.Email,
			OtherId:            refIds,
			FirstName:          e.FirstName,
			FullName:           e.FullName,
			Id:                 e.PublicID,
			LastName:           e.LastName,
			JobCategory:        e.JobCategory,
			Orcid:              e.Orcid,
			OrcidToken:         e.OrcidToken,
			OrganizationId:     orgIds,
			PreferredLastName:  e.PreferredLastName,
			PreferredFirstName: e.PreferredFirstName,
			Title:              e.Title,
		},
	}
	return p
}
