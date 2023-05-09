package db

import (
	"context"
	"errors"
	"fmt"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/ent"
	"github.com/ugent-library/people/ent/organization"
	"github.com/ugent-library/people/ent/person"
	"github.com/ugent-library/people/ent/schema"
	"github.com/ugent-library/people/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type organizationService struct {
	client *ent.Client
}

func NewOrganizationService(client *ent.Client) (*organizationService, error) {

	execQuery := `
	BEGIN;
	ALTER TABLE organization 
	ADD COLUMN IF NOT EXISTS ts tsvector GENERATED ALWAYS AS 
	(
		to_tsvector('simple', jsonb_path_query_array(other_id, '$[*].id')) || 
		to_tsvector('simple', public_id) || 
		to_tsvector('simple',name_dut) || 
		to_tsvector('simple', name_eng)
	) STORED;
	CREATE INDEX IF NOT EXISTS ts_idx ON organization USING GIN(ts);
	COMMIT;
	`

	if _, err := client.ExecContext(context.Background(), execQuery); err != nil {
		return nil, err
	}

	return &organizationService{
		client: client,
	}, nil
}

func (orgSvc *organizationService) GetOrganization(ctx context.Context, id string) (*models.Organization, error) {
	row, err := orgSvc.client.Organization.Query().WithParent().Where(organization.PublicIDEQ(id)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return orgUnwrap(row), nil
}

func (orgSvc *organizationService) CreateOrganization(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	// date fields filled by schema
	tx, txErr := orgSvc.client.Tx(ctx)
	if txErr != nil {
		return nil, fmt.Errorf("unable to start transaction: %w", txErr)
	}
	defer tx.Rollback()

	t := tx.Organization.Create()

	t.SetPublicID(org.Id)
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

	if org.ParentId != "" {
		_, err := tx.Organization.Query().Where(organization.PublicIDEQ(org.ParentId)).First(ctx)
		if err != nil {
			var e *ent.NotFoundError
			if errors.As(err, &e) {
				t.SetOtherParentID(org.ParentId)
			} else {
				return nil, fmt.Errorf("unable to query organizations: %w", err)
			}
		}
	}

	row, err := t.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to save organization: %w", err)
	}

	// check if other_parent_id is using this new organization (upgrade)
	childOrgUpdate := tx.Organization.Update().
		Where(organization.OtherParentIDEQ(org.Id)).
		SetOtherParentID("").
		SetParent(row)
	if _, err := childOrgUpdate.Save(ctx); err != nil {
		return nil, fmt.Errorf("unable to update child organizations: %w", err)
	}

	// move organization id from person.other_organization_id into real relation
	// TODO: also do this "migration" in UpdateOrganization?

	// create relations for person records that are referring to that new org
	_, sqlResErr := tx.OrganizationPerson.ExecContext(ctx, fmt.Sprintf(`
		INSERT INTO organization_person(person_id, organization_id, date_created, date_updated)
		SELECT id person_id, %d organization_id, now() date_created, now() date_updated
		FROM person WHERE person.other_organization_id @> '"%s"'
	`, row.ID, org.Id))

	if sqlResErr != nil {
		return nil, fmt.Errorf("unable to insert into table organization_person: %w", sqlResErr)
	}

	/*
		Remove value from person.other_organization_id:

		sql:
			update person set other_organization_id = other_organization_id - 'CA20' where other_organization_id @> '"CA20"';
	*/
	personUpdateErr := tx.Person.Update().Where(func(s *entsql.Selector) {
		s.Where(sqljson.ValueContains("other_organization_id", org.Id))
	}).Modify(func(u *entsql.UpdateBuilder) {
		u.Set(
			person.FieldOtherOrganizationID,
			entsql.ExprFunc(func(b *entsql.Builder) {
				b.Ident(person.FieldOtherOrganizationID).WriteOp(entsql.OpSub).Arg(org.Id)
			}),
		)
	}).Exec(ctx)

	if personUpdateErr != nil {
		return nil, fmt.Errorf("unable to update person.other_organization_id: %w", personUpdateErr)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("unable to commit transaction: %w", err)
	}

	// collect entgo managed fields
	org.DateCreated = timestamppb.New(row.DateCreated)
	org.DateUpdated = timestamppb.New(row.DateUpdated)

	return org, nil
}

func (orgSvc *organizationService) UpdateOrganization(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	tx, txErr := orgSvc.client.Tx(ctx)
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

	t.ClearParent()
	t.SetOtherParentID("")
	if org.ParentId != "" {
		parentOrg, err := tx.Organization.Query().Where(organization.PublicIDEQ(org.ParentId)).First(ctx)
		if err != nil {
			var e *ent.NotFoundError
			if errors.As(err, &e) {
				t.SetOtherParentID(org.ParentId)
			} else {
				return nil, fmt.Errorf("unable to query organizations: %w", err)
			}
		} else {
			t.SetParent(parentOrg)
		}
	}

	nAffected, err := t.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to save organization: %w", err)
	}

	// load new row (must be found)
	row, err := tx.Organization.Query().Where(organization.PublicIDEQ(org.Id)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to query organizations: %w", err)
	}

	// check if other_parent_id is using this new organization (upgrade)
	if nAffected > 0 {
		childOrgUpdate := tx.Organization.Update().
			Where(organization.OtherParentIDEQ(org.Id)).
			SetOtherParentID("").
			SetParent(row)
		if _, err := childOrgUpdate.Save(ctx); err != nil {
			return nil, fmt.Errorf("unable to update child organizations: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("unable to commit transaction: %w", err)
	}

	return orgUnwrap(row), nil
}

func (orgSvc *organizationService) DeleteOrganization(ctx context.Context, id string) error {
	_, err := orgSvc.client.Organization.Delete().Where(organization.PublicIDEQ(id)).Exec(ctx)
	return err
}

func (orgSvc *organizationService) EachOrganization(ctx context.Context, cb func(*models.Organization) bool) error {

	// TODO: find a better way to do this (no cursors possible)
	var offset int = 0
	var limit int = 500
	for {
		rows, err := orgSvc.client.Organization.Query().WithParent().Offset(offset).Limit(limit).Order(ent.Asc(organization.FieldDateCreated)).All(ctx)
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

func (orgSvc *organizationService) SuggestOrganization(ctx context.Context, query string) ([]*models.Organization, error) {
	rows, err := orgSvc.client.Organization.Query().WithParent().Where(func(s *entsql.Selector) {
		s.Where(
			toTSQuery("ts", query),
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
	org := &models.Organization{
		Organization: v1.Organization{
			Id:          e.PublicID,
			DateCreated: timestamppb.New(e.DateCreated),
			DateUpdated: timestamppb.New(e.DateUpdated),
			Type:        e.Type,
			NameDut:     e.NameDut,
			NameEng:     e.NameEng,
			OtherId:     otherIds,
		},
		OtherParentId: e.OtherParentID,
	}
	if parentOrg := e.Edges.Parent; parentOrg != nil {
		org.ParentId = parentOrg.PublicID
	}
	return org
}
