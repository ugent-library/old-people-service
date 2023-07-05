package repository

import (
	"context"
	"errors"
	"fmt"
	"sort"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/jackc/pgx/v5"
	v1 "github.com/ugent-library/person-service/api/v1"
	"github.com/ugent-library/person-service/ent"
	"github.com/ugent-library/person-service/ent/organization"
	"github.com/ugent-library/person-service/ent/person"
	"github.com/ugent-library/person-service/ent/schema"
	"github.com/ugent-library/person-service/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type repository struct {
	client *ent.Client
	secret []byte
}

func NewRepository(config *Config) (*repository, error) {
	return &repository{
		client: config.Client,
		secret: []byte(config.AesKey),
	}, nil
}

func (repo *repository) GetOrganization(ctx context.Context, id string) (*models.Organization, error) {
	row, err := repo.client.Organization.Query().WithParent().Where(organization.PublicIDEQ(id)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return orgUnwrap(row), nil
}

func (repo *repository) GetOrganizationByGismoId(ctx context.Context, gismoId string) (*models.Organization, error) {
	row, err := repo.client.Organization.Query().WithParent().Where(organization.GismoIDEQ(gismoId)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return orgUnwrap(row), nil
}

func (repo *repository) GetOrganizationsByGismoId(ctx context.Context, gismoIds ...string) ([]*models.Organization, error) {
	rows, err := repo.client.Organization.Query().WithParent().Where(
		organization.GismoIDIn(gismoIds...),
	).Order(func(s *entsql.Selector) {
		orderStr := fmt.Sprintf("array_position($%d, gismo_id)", len(gismoIds)+1)
		s.OrderExpr(entsql.ExprP(orderStr, gismoIds))
	}).All(ctx)

	if err != nil {
		return nil, err
	}

	orgs := make([]*models.Organization, 0, len(rows))
	for _, row := range rows {
		orgs = append(orgs, orgUnwrap(row))
	}

	return orgs, nil
}

func (repo *repository) GetOrganizationByOtherId(ctx context.Context, typ string, val string) (*models.Organization, error) {
	row, err := repo.client.Organization.Query().WithParent().Where(func(s *entsql.Selector) {
		exprVal := fmt.Sprintf(`[{"id":"%s","type":"%s"}]`, val, typ)
		s.Where(entsql.ExprP("other_id::jsonb @> $1", exprVal))
	}).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return orgUnwrap(row), nil
}

func (repo *repository) SaveOrganization(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	if org.IsStored() {
		return repo.UpdateOrganization(ctx, org)
	}
	return repo.CreateOrganization(ctx, org)
}

func (repo *repository) CreateOrganization(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	// date fields filled by schema
	tx, txErr := repo.client.Tx(ctx)
	if txErr != nil {
		return nil, fmt.Errorf("unable to start transaction: %w", txErr)
	}
	defer tx.Rollback()

	t := tx.Organization.Create()

	t.SetNameDut(org.NameDut)
	t.SetNameEng(org.NameEng)
	t.SetType(org.Type)
	schemaOtherIds := make([]schema.IdRef, 0, len(org.OtherId))
	for _, refId := range org.OtherId {
		schemaOtherIds = append(schemaOtherIds, schema.IdRef{
			ID:   refId.Id,
			Type: refId.Type,
		})
	}
	t.SetOtherID(schemaOtherIds)
	t.SetGismoID(org.GismoId)
	if org.ParentId != "" {
		parentOrgRow, err := tx.Organization.Query().Where(organization.PublicIDEQ(org.ParentId)).First(ctx)
		if err != nil {
			var e *ent.NotFoundError
			if errors.As(err, &e) {
				return nil, fmt.Errorf("parent organization with public_id %s not found", org.ParentId)
			} else {
				return nil, fmt.Errorf("unable to query organizations: %w", err)
			}
		}
		t.SetParent(parentOrgRow)
	}

	row, err := t.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to save organization: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("unable to commit transaction: %w", err)
	}

	// collect entgo managed fields
	org.DateCreated = timestamppb.New(row.DateCreated)
	org.DateUpdated = timestamppb.New(row.DateUpdated)

	return org, nil
}

func (repo *repository) UpdateOrganization(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	tx, txErr := repo.client.Tx(ctx)
	if txErr != nil {
		return nil, fmt.Errorf("unable to start transaction: %w", txErr)
	}
	defer tx.Rollback()

	t := tx.Organization.Update().Where(organization.PublicIDEQ(org.Id))

	t.SetNameDut(org.NameDut)
	t.SetNameEng(org.NameEng)
	t.SetType(org.Type)
	schemaOtherIds := make([]schema.IdRef, 0, len(org.OtherId))
	for _, refId := range org.OtherId {
		schemaOtherIds = append(schemaOtherIds, schema.IdRef{
			ID:   refId.Id,
			Type: refId.Type,
		})
	}
	t.SetOtherID(schemaOtherIds)
	t.SetGismoID(org.GismoId)
	t.ClearParent()
	if org.ParentId != "" {
		parentOrg, err := tx.Organization.Query().Where(organization.PublicIDEQ(org.ParentId)).First(ctx)
		if err != nil {
			var e *ent.NotFoundError
			if errors.As(err, &e) {
				return nil, fmt.Errorf("parent organization with public_id %s not found", org.ParentId)
			} else {
				return nil, fmt.Errorf("unable to query organizations: %w", err)
			}
		} else {
			t.SetParent(parentOrg)
		}
	}

	if _, err := t.Save(ctx); err != nil {
		return nil, fmt.Errorf("unable to save organization: %w", err)
	}

	// load new row (must be found)
	row, err := tx.Organization.Query().Where(organization.PublicIDEQ(org.Id)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to query organizations: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("unable to commit transaction: %w", err)
	}

	return orgUnwrap(row), nil
}

func (repo *repository) DeleteOrganization(ctx context.Context, id string) error {
	_, err := repo.client.Organization.Delete().Where(organization.PublicIDEQ(id)).Exec(ctx)
	return err
}

func (repo *repository) EachOrganization(ctx context.Context, cb func(*models.Organization) bool) error {

	// TODO: find a better way to do this (no cursors possible)
	var offset int = 0
	var limit int = 500
	for {
		rows, err := repo.client.Organization.Query().WithParent().Offset(offset).Limit(limit).Order(ent.Asc(organization.FieldDateCreated)).All(ctx)
		if err != nil {
			return err
		}
		// entgo returns no error on empty results
		if len(rows) == 0 {
			break
		}
		for _, row := range rows {
			if !cb(orgUnwrap(row)) {
				return nil
			}
		}
		offset += limit
	}

	return nil
}

func (repo *repository) SuggestOrganization(ctx context.Context, query string) ([]*models.Organization, error) {
	tsQuery, tsQueryArgs := toTSQuery(query)
	tsQuery = "ts @@ " + tsQuery
	rows, err := repo.client.Organization.Query().WithParent().Where(func(s *entsql.Selector) {
		s.Where(
			entsql.ExprP(tsQuery, tsQueryArgs...),
		)
	}).Limit(10).All(ctx)

	if err != nil {
		return nil, err
	}

	orgs := make([]*models.Organization, 0, len(rows))
	for _, row := range rows {
		orgs = append(orgs, orgUnwrap(row))
	}

	return orgs, nil
}

func orgUnwrap(e *ent.Organization) *models.Organization {
	otherIds := make([]*v1.IdRef, 0, len(e.OtherID))
	for _, schemaOtherId := range e.OtherID {
		otherIds = append(otherIds, &v1.IdRef{
			Id:   schemaOtherId.ID,
			Type: schemaOtherId.Type,
		})
	}
	var gismoId string = ""
	if e.GismoID != nil {
		gismoId = *e.GismoID
	}
	org := &models.Organization{
		Organization: &v1.Organization{
			Id:          e.PublicID,
			GismoId:     gismoId,
			DateCreated: timestamppb.New(e.DateCreated),
			DateUpdated: timestamppb.New(e.DateUpdated),
			Type:        e.Type,
			NameDut:     e.NameDut,
			NameEng:     e.NameEng,
			OtherId:     otherIds,
		},
	}
	if parentOrg := e.Edges.Parent; parentOrg != nil {
		org.ParentId = parentOrg.PublicID
	}
	return org
}

func (repo *repository) SavePerson(ctx context.Context, p *models.Person) (*models.Person, error) {
	if p.IsStored() {
		return repo.UpdatePerson(ctx, p)
	}
	return repo.CreatePerson(ctx, p)
}

func (repo *repository) CreatePerson(ctx context.Context, p *models.Person) (*models.Person, error) {
	// date fields filled by schema
	tx, txErr := repo.client.Tx(ctx)
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
	t.SetTitle(p.Title)
	t.SetLastName(p.LastName)
	t.SetOrcid(p.Orcid)
	t.SetRole(p.Role)
	t.SetSettings(p.Settings)
	var gismoId *string = nil
	if p.GismoId != "" {
		gismoId = &p.GismoId
	}
	t.SetNillableGismoID(gismoId)
	if len(p.ObjectClass) > 0 {
		t.SetObjectClass(p.ObjectClass)
	} else {
		t.SetObjectClass(nil)
	}
	t.SetExpirationDate(p.ExpirationDate)

	if p.OrcidToken == "" {
		t.SetOrcidToken("")
	} else {
		eToken, eTokenErr := encryptMessage(repo.secret, p.OrcidToken)
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

	if len(p.OrganizationId) > 0 {
		// TODO: crashes with segmentation violation error when org does not exist
		orgs, err := tx.Organization.Query().Where(organization.PublicIDIn(p.OrganizationId...)).All(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to query organizations: %w", err)
		}
		if len(p.OrganizationId) != len(orgs) {
			return nil, fmt.Errorf("person.organization_id contains invalid organization id's")
		}
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
	p.Id = row.PublicID

	return p, nil
}

func (repo *repository) SetPersonOrcid(ctx context.Context, id string, orcid string) error {
	nUpdated, updateErr := repo.client.
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

func (repo *repository) SetPersonOrcidToken(ctx context.Context, id string, orcidToken string) error {
	var uToken string
	var uTokenErr error

	if orcidToken == "" {
		uToken = ""
	} else {
		uToken, uTokenErr = encryptMessage(repo.secret, orcidToken)
	}

	if uTokenErr != nil {
		return fmt.Errorf("unable to encrypt orcid_token: %w", uTokenErr)
	}

	nUpdated, updateErr := repo.client.
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

func (repo *repository) UpdatePerson(ctx context.Context, p *models.Person) (*models.Person, error) {
	tx, txErr := repo.client.Tx(ctx)
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
	var gismoId *string = nil
	if p.GismoId != "" {
		gismoId = &p.GismoId
	}
	t.SetNillableGismoID(gismoId)
	t.SetObjectClass(p.ObjectClass)
	t.SetExpirationDate(p.ExpirationDate)

	if p.OrcidToken == "" {
		t.SetOrcidToken("")
	} else {
		eToken, eTokenErr := encryptMessage(repo.secret, p.OrcidToken)
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
	if len(p.OrganizationId) > 0 {
		// TODO: crashes with segmentation violation error when org does not exist
		orgs, err := tx.Organization.Query().Where(organization.PublicIDIn(p.OrganizationId...)).All(ctx)
		if err != nil {
			return nil, err
		}
		if len(p.OrganizationId) != len(orgs) {
			return nil, fmt.Errorf("person.organization_id contains invalid organization id's")
		}
		t.AddOrganizations(orgs...)
	}

	if _, err := t.Save(ctx); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return repo.GetPerson(ctx, p.Id)
}

func (repo *repository) GetPerson(ctx context.Context, id string) (*models.Person, error) {
	row, err := repo.client.Person.Query().WithOrganizations().Where(person.PublicIDEQ(id)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}

	return repo.personUnwrap(row)
}

func (repo *repository) GetPersonByOtherId(ctx context.Context, typ string, val string) (*models.Person, error) {
	row, err := repo.client.Person.Query().WithOrganizations().Where(func(s *entsql.Selector) {
		exprVal := fmt.Sprintf(`[{"id":"%s","type":"%s"}]`, val, typ)
		s.Where(entsql.ExprP("other_id::jsonb @> $1", exprVal))
	}).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return repo.personUnwrap(row)
}

func (repo *repository) GetPersonByGismoId(ctx context.Context, gismoId string) (*models.Person, error) {
	row, err := repo.client.Person.Query().WithOrganizations().Where(person.GismoID(gismoId)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return repo.personUnwrap(row)
}

func (repo *repository) DeletePerson(ctx context.Context, id string) error {
	_, err := repo.client.Person.Delete().Where(person.PublicIDEQ(id)).Exec(ctx)
	return err
}

func (repo *repository) EachPerson(ctx context.Context, cb func(*models.Person) bool) error {
	// TODO: find a better way to do this (no cursors possible)
	var offset int = 0
	var limit int = 500
	for {
		rows, err := repo.client.Person.Query().WithOrganizations().Offset(offset).Limit(limit).Order(ent.Asc(person.FieldDateCreated)).All(ctx)
		if err != nil {
			return err
		}
		// entgo returns no error on empty results
		if len(rows) == 0 {
			break
		}
		for _, row := range rows {
			p, err := repo.personUnwrap(row)
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

func (repo *repository) SuggestPerson(ctx context.Context, query string) ([]*models.Person, error) {
	// fetch ids via a raw query ..
	tsQuery, tsQueryArgs := toTSQuery(query)
	sqlQuery := fmt.Sprintf(
		"SELECT id, ts_rank(ts, %s) as rank FROM person WHERE ts @@ %s ORDER BY rank DESC LIMIT %d",
		tsQuery,
		tsQuery,
		10,
	)
	rows, err := repo.client.QueryContext(ctx, sqlQuery, tsQueryArgs...)

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
	records, err := repo.client.
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
		if person, err := repo.personUnwrap(record); err != nil {
			return nil, err
		} else {
			persons = append(persons, person)
		}
	}

	return persons, nil
}

func (repo *repository) personUnwrap(e *ent.Person) (*models.Person, error) {
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
		uToken, uTokenErr = decryptMessage(repo.secret, e.OrcidToken)
		if uTokenErr != nil {
			return nil, fmt.Errorf("unable to decrypt orcid_token: %w", uTokenErr)
		}
	} else {
		uToken = e.OrcidToken
	}

	var gismoId string = ""
	if e.GismoID != nil {
		gismoId = *e.GismoID
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
			GismoId:            gismoId,
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
			ObjectClass:        e.ObjectClass,
		},
	}
	return p, nil
}

func (repo *repository) SetPersonRole(ctx context.Context, id string, roles []string) error {
	nUpdated, updateErr := repo.client.
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

func (repo *repository) SetPersonSettings(ctx context.Context, id string, settings map[string]string) error {
	nUpdated, updateErr := repo.client.
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
