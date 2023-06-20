package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/ent"
	"github.com/ugent-library/people/ent/organization"
	"github.com/ugent-library/people/ent/person"
	"github.com/ugent-library/people/ent/schema"
	"github.com/ugent-library/people/models"
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

	// TODO: update field list if ent schema changes
	fields := []string{
		"id", "date_created", "date_updated", "public_id", "active",
		"birth_date", "email", "other_id", "other_organization_id",
		"first_name", "full_name", "last_name", "job_category",
		"orcid", "orcid_token", "preferred_first_name", "preferred_last_name",
		"title", "role", "settings",
	}

	tsQuery, tsQueryArgs := toTSQuery(query)
	sqlQuery := "SELECT " + strings.Join(fields, ", ")
	sqlQuery += fmt.Sprintf(", ts_rank(ts, %s) as rank", tsQuery)
	sqlQuery += " FROM person WHERE ts @@ " + tsQuery
	sqlQuery += " ORDER BY rank DESC"
	rows, err := ps.client.QueryContext(ctx, sqlQuery, tsQueryArgs...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	persons := make([]*models.Person, 0)

	for rows.Next() {
		ep := ent.Person{}

		var otherId sql.NullString
		var otherOrganizationID sql.NullString
		var jobCategory sql.NullString
		var role sql.NullString
		var settings sql.NullString
		rank := ""
		var active sql.NullBool
		var orcid sql.NullString
		var orcidToken sql.NullString
		var birthDate sql.NullString
		var email sql.NullString
		var firstName sql.NullString
		var lastName sql.NullString
		var fullName sql.NullString
		var preferredFirstName sql.NullString
		var preferredLastName sql.NullString
		var title sql.NullString

		err := rows.Scan(
			&ep.ID,
			&ep.DateCreated,
			&ep.DateUpdated,
			&ep.PublicID,
			&active,
			&birthDate,
			&email,
			&otherId,
			&otherOrganizationID,
			&firstName,
			&fullName,
			&lastName,
			&jobCategory,
			&orcid,
			&orcidToken,
			&preferredFirstName,
			&preferredLastName,
			&title,
			&role,
			&settings,
			&rank,
		)

		if err == pgx.ErrNoRows {
			return persons, nil
		}

		if err != nil {
			return nil, err
		}

		if active.Valid {
			ep.Active = active.Bool
		}
		if title.Valid {
			ep.Title = title.String
		}
		if preferredFirstName.Valid {
			ep.PreferredFirstName = preferredFirstName.String
		}
		if preferredLastName.Valid {
			ep.PreferredLastName = preferredLastName.String
		}
		if firstName.Valid {
			ep.FirstName = firstName.String
		}
		if lastName.Valid {
			ep.LastName = lastName.String
		}
		if fullName.Valid {
			ep.FullName = fullName.String
		}
		if birthDate.Valid {
			ep.BirthDate = birthDate.String
		}
		if email.Valid {
			ep.Email = email.String
		}
		if orcid.Valid {
			ep.Orcid = orcid.String
		}
		if orcidToken.Valid {
			ep.OrcidToken = orcidToken.String
		}

		if len(otherId.String) == 0 {

		} else if err := json.Unmarshal([]byte(otherId.String), &ep.OtherID); err != nil {
			return nil, fmt.Errorf("unable to decode other_id: %w", err)
		}
		if len(otherOrganizationID.String) == 0 {

		} else if err := json.Unmarshal([]byte(otherOrganizationID.String), &ep.OtherOrganizationID); err != nil {
			return nil, fmt.Errorf("unable to decode other_organization_id: %w", err)
		}
		if len(jobCategory.String) == 0 {

		} else if err := json.Unmarshal([]byte(jobCategory.String), &ep.JobCategory); err != nil {
			return nil, fmt.Errorf("unable to decode job_category: %w", err)
		}
		if len(role.String) == 0 {

		} else if err := json.Unmarshal([]byte(role.String), &ep.Role); err != nil {
			return nil, fmt.Errorf("unable to decode role: %w", err)
		}
		if len(settings.String) == 0 {

		} else if err := json.Unmarshal([]byte(settings.String), &ep.Settings); err != nil {
			return nil, fmt.Errorf("unable to decode settings: %w", err)
		}

		person, err := ps.personUnwrap(&ep)
		if err != nil {
			return nil, err
		}

		persons = append(persons, person)
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
