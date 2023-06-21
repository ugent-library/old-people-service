package repository

import (
	"context"
	"errors"
	"fmt"
	"sort"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	v1 "github.com/ugent-library/person-service/api/v1"
	"github.com/ugent-library/person-service/ent"
	"github.com/ugent-library/person-service/ent/organization"
	"github.com/ugent-library/person-service/ent/person"
	"github.com/ugent-library/person-service/ent/schema"
	"github.com/ugent-library/person-service/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type personService struct {
	client *ent.Client
	secret []byte
}

func NewPersonService(config *Config) (*personService, error) {

	return &personService{
		client: config.Client,
		secret: []byte(config.AesKey),
	}, nil
}

func (ps *personService) CreatePerson(ctx context.Context, p *models.Person) (*models.Person, error) {
	// date fields filled by schema
	tx, txErr := ps.client.Tx(ctx)
	if txErr != nil {
		return nil, fmt.Errorf("unable to start transaction: %w", txErr)
	}
	defer tx.Rollback()

	t := tx.Person.Create()

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
	t.SetRole(p.Role)
	t.SetSettings(p.Settings)

	if p.OrcidToken == "" {
		t.SetOrcidToken("")
	} else {
		eToken, eTokenErr := encryptMessage(ps.secret, p.OrcidToken)
		if eTokenErr != nil {
			return nil, fmt.Errorf("unable to encrypt orcid_token: %w", eTokenErr)
		}
		t.SetOrcidToken(eToken)
	}

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

	if p.OrganizationId != nil && len(p.OrganizationId) > 0 {
		// TODO: crashes with segmentation violation error when org does not exist
		orgs, err := tx.Organization.Query().Where(organization.PublicIDIn(p.OrganizationId...)).All(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to query organizations: %w", err)
		}
		// add unconfirmed organization id's to other_organization_id
		otherOrganizationIds := []string{}
		for _, orgId := range p.OrganizationId {
			found := false
			for _, org := range orgs {
				if org.PublicID == orgId {
					found = true
					break
				}
			}
			if !found {
				otherOrganizationIds = append(otherOrganizationIds, orgId)
			}
		}
		t.SetOtherOrganizationID(otherOrganizationIds)
		t.AddOrganizations(orgs...)
	}

	row, err := t.Save(ctx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("unable to commit transaction: %w", err)
	}

	// collect entgo managed fields
	p.DateCreated = timestamppb.New(row.DateCreated)
	p.DateUpdated = timestamppb.New(row.DateUpdated)

	return p, nil
}

func (ps *personService) SetOrcid(ctx context.Context, id string, orcid string) error {
	nUpdated, updateErr := ps.client.
		Person.
		Update().
		Where(person.PublicIDEQ(id)).
		SetOrcid(orcid).
		Save(ctx)

	if updateErr != nil {
		return updateErr
	}

	if nUpdated == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (ps *personService) SetOrcidToken(ctx context.Context, id string, orcidToken string) error {

	var uToken string
	var uTokenErr error

	if orcidToken == "" {
		uToken = ""
	} else {
		uToken, uTokenErr = encryptMessage(ps.secret, orcidToken)
	}

	if uTokenErr != nil {
		return fmt.Errorf("unable to encrypt orcid_token: %w", uTokenErr)
	}

	nUpdated, updateErr := ps.client.
		Person.
		Update().
		Where(person.PublicIDEQ(id)).
		SetOrcidToken(uToken).
		Save(ctx)

	if updateErr != nil {
		return updateErr
	}

	if nUpdated == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (ps *personService) UpdatePerson(ctx context.Context, p *models.Person) (*models.Person, error) {
	tx, txErr := ps.client.Tx(ctx)
	if txErr != nil {
		return nil, fmt.Errorf("unable to start transaction: %w", txErr)
	}
	defer tx.Rollback()

	t := tx.Person.Update().Where(person.PublicIDEQ(p.Id))

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
	t.SetRole(p.Role)
	t.SetSettings(p.Settings)

	if p.OrcidToken == "" {
		t.SetOrcidToken("")
	} else {
		eToken, eTokenErr := encryptMessage(ps.secret, p.OrcidToken)
		if eTokenErr != nil {
			return nil, fmt.Errorf("unable to encrypt orcid_token: %w", eTokenErr)
		}
		t.SetOrcidToken(eToken)
	}

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
	otherOrganizationIds := make([]string, 0)
	if p.OrganizationId != nil && len(p.OrganizationId) > 0 {
		// TODO: crashes with segmentation violation error when org does not exist
		orgs, err := tx.Organization.Query().Where(organization.PublicIDIn(p.OrganizationId...)).All(ctx)
		if err != nil {
			return nil, err
		}
		// move unconfirmed organizations to other_organization_id
		for _, orgId := range p.OrganizationId {
			found := false
			for _, org := range orgs {
				if org.PublicID == orgId {
					found = true
					break
				}
			}
			if !found {
				otherOrganizationIds = append(otherOrganizationIds, orgId)
			}
		}

		t.AddOrganizations(orgs...)
	}
	t.SetOtherOrganizationID(otherOrganizationIds)

	_, err := t.Save(ctx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return ps.GetPerson(ctx, p.Id)
}

func (ps *personService) GetPerson(ctx context.Context, id string) (*models.Person, error) {
	row, err := ps.client.Person.Query().WithOrganizations().Where(person.PublicIDEQ(id)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}

	return ps.personUnwrap(row)
}

func (ps *personService) DeletePerson(ctx context.Context, id string) error {
	_, err := ps.client.Person.Delete().Where(person.PublicIDEQ(id)).Exec(ctx)
	return err
}

func (ps *personService) EachPerson(ctx context.Context, cb func(*models.Person) bool) error {

	// TODO: find a better way to do this (no cursors possible)
	var offset int = 0
	var limit int = 500
	for {
		rows, err := ps.client.Person.Query().WithOrganizations().Offset(offset).Limit(limit).Order(ent.Asc(person.FieldDateCreated)).All(ctx)
		if err != nil {
			return err
		}
		// entgo returns no error on empty results
		if len(rows) == 0 {
			break
		}
		for _, row := range rows {
			p, err := ps.personUnwrap(row)
			if err != nil {
				return err
			}
			if !cb(p) {
				return nil
			}
		}
		offset += limit
	}

	return nil
}

func (ps *personService) SuggestPerson(ctx context.Context, query string) ([]*models.Person, error) {

	// fetch ids via a raw query ..
	tsQuery, tsQueryArgs := toTSQuery(query)
	sqlQuery := fmt.Sprintf(
		"SELECT id, ts_rank(ts, %s) as rank FROM person WHERE ts @@ %s ORDER BY rank DESC LIMIT %d",
		tsQuery,
		tsQuery,
		10,
	)
	rows, err := ps.client.QueryContext(ctx, sqlQuery, tsQueryArgs...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ids := make([]int, 0, 10)

	for rows.Next() {
		var id int
		var rank float64
		err := rows.Scan(&id, &rank)
		if err == pgx.ErrNoRows {
			return []*models.Person{}, nil
		}
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	// ... and then fetch the records via ent
	// make sure order of ids is the same
	records, err := ps.client.
		Person.
		Query().
		WithOrganizations().
		Where(person.IDIn(ids...)).
		Order(func(s *entsql.Selector) {
			orderStr := fmt.Sprintf("array_position($%d, id)", len(ids)+1)
			s.OrderExpr(entsql.ExprP(orderStr, ids))
		}).
		All(ctx)

	if err != nil {
		return nil, err
	}

	persons := make([]*models.Person, 0, len(records))

	for _, record := range records {
		if person, err := ps.personUnwrap(record); err != nil {
			return nil, err
		} else {
			persons = append(persons, person)
		}
	}

	return persons, nil
}

func (ps *personService) personUnwrap(e *ent.Person) (*models.Person, error) {
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

	var uToken string
	var uTokenErr error
	if e.OrcidToken != "" {
		uToken, uTokenErr = decryptMessage(ps.secret, e.OrcidToken)
		if uTokenErr != nil {
			return nil, fmt.Errorf("unable to decrypt orcid_token: %w", uTokenErr)
		}
	} else {
		uToken = e.OrcidToken
	}

	p := &models.Person{
		Person: &v1.Person{
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
			OrcidToken:         uToken,
			OrganizationId:     orgIds,
			PreferredLastName:  e.PreferredLastName,
			PreferredFirstName: e.PreferredFirstName,
			Title:              e.Title,
			Role:               e.Role,
			Settings:           e.Settings,
		},
	}
	return p, nil
}

func (ps *personService) SetRole(ctx context.Context, id string, roles []string) error {
	nUpdated, updateErr := ps.client.
		Person.
		Update().
		Where(person.PublicIDEQ(id)).
		SetRole(roles).
		Save(ctx)

	if updateErr != nil {
		return updateErr
	}

	if nUpdated == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (ps *personService) SetSettings(ctx context.Context, id string, settings map[string]string) error {
	nUpdated, updateErr := ps.client.
		Person.
		Update().
		Where(person.PublicIDEQ(id)).
		SetSettings(settings).
		Save(ctx)

	if updateErr != nil {
		return updateErr
	}

	if nUpdated == 0 {
		return models.ErrNotFound
	}

	return nil
}
