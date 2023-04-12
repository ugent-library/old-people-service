package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	entdialect "entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/ent"
	entmigrate "github.com/ugent-library/people/ent/migrate"
	"github.com/ugent-library/people/ent/organization"
	"github.com/ugent-library/people/ent/schema"
	"github.com/ugent-library/people/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrganizationConfig struct {
	DB string
}

type organizationService struct {
	db *ent.Client
}

func NewOrganizationService(cfg *OrganizationConfig) (models.OrganizationService, error) {
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

	return &organizationService{
		db: client,
	}, nil
}

func (os *organizationService) GetOrganization(ctx context.Context, id string) (*models.Organization, error) {
	row, err := os.db.Organization.Query().WithParent().Where(organization.PublicIDEQ(id)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return orgUnwrap(row), nil
}

func (os *organizationService) CreateOrganization(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	// date fields filled by schema
	tx, txErr := os.db.Tx(ctx)
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
				pt := tx.Organization.Create()
				pt.SetPublicID(org.ParentId)
				pt.SetNameDut(org.ParentId)
				pt.SetNameEng(org.ParentId)
				parentOrg, err := pt.Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("unable to save parent organization: %w", err)
				}
				t.SetParent(parentOrg)
			} else {
				return nil, fmt.Errorf("unable to query organizations: %w", err)
			}
		}
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

func (os *organizationService) UpdateOrganization(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	tx, txErr := os.db.Tx(ctx)
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

	if org.ParentId != "" {
		_, err := tx.Organization.Query().Where(organization.PublicIDEQ(org.ParentId)).First(ctx)
		if err != nil {
			var e *ent.NotFoundError
			if errors.As(err, &e) {
				pt := tx.Organization.Create()
				pt.SetPublicID(org.ParentId)
				pt.SetNameDut(org.ParentId)
				pt.SetNameEng(org.ParentId)
				parentOrg, err := pt.Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("unable to save parent organization: %w", err)
				}
				t.SetParent(parentOrg)
			} else {
				return nil, fmt.Errorf("unable to query organizations: %w", err)
			}
		}
	}

	_, err := t.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to save organization: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("unable to commit transaction: %w", err)
	}

	return os.GetOrganization(ctx, org.Id)
}

func (os *organizationService) DeleteOrganization(ctx context.Context, id string) error {
	_, err := os.db.Organization.Delete().Where(organization.PublicIDEQ(id)).Exec(ctx)
	return err
}

func (os *organizationService) EachOrganization(ctx context.Context, cb func(*models.Organization) bool) error {

	// TODO: find a better way to do this (no cursors possible)
	var offset int = 0
	var limit int = 500
	for {
		rows, err := os.db.Organization.Query().WithParent().Offset(offset).Limit(limit).Order(ent.Asc(organization.FieldDateCreated)).All(ctx)
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
	}
	if parentOrg := e.Edges.Parent; parentOrg != nil {
		org.ParentId = parentOrg.PublicID
	}
	return org
}
