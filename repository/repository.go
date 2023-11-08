package repository

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v5"
	"github.com/samber/lo"
	"github.com/ugent-library/crypt"
	"github.com/ugent-library/people-service/ent"
	"github.com/ugent-library/people-service/ent/organization"
	"github.com/ugent-library/people-service/ent/organizationparent"
	"github.com/ugent-library/people-service/ent/organizationperson"
	"github.com/ugent-library/people-service/ent/person"
	"github.com/ugent-library/people-service/ent/schema"
	"github.com/ugent-library/people-service/models"
)

const (
	personPageLimit          = 200
	organizationPageLimit    = 200
	organizationSuggestLimit = 10
	personSuggestLimit       = 10
)

type repository struct {
	client *ent.Client
	secret []byte
}
type setCursor struct {
	// IMPORTANT: auto increment (of id) starts with 1, so default value 0 should never match
	LastID int `json:"l"`
}

type organizationParent struct {
	models.OrganizationParent
	organizationID       int
	parentOrganizationID int
}

type organizationMember struct {
	models.OrganizationMember
	organizationID int
	personID       int
}

func NewRepository(config *Config) (*repository, error) {
	client, err := openClient(config.DbUrl)
	if err != nil {
		return nil, err
	}
	return &repository{
		client: client,
		secret: []byte(config.AesKey),
	}, nil
}

func (repo *repository) getOrganizationMembers(ctx context.Context, ids ...int) ([]organizationMember, error) {
	query := `
SELECT "organization_id",
       "person_id",
	   "from",
	   "until",
	   "date_created",
	   "date_updated",
	   (SELECT "public_id" FROM "organization" WHERE "id" = op.organization_id) organization_public_id
FROM "organization_person" op
WHERE "person_id" = any($1)
ORDER by "from" ASC
	`
	pgIds := pgtype.Int4Array{}
	pgIds.Set(ids)
	rows, err := repo.client.QueryContext(
		ctx,
		query,
		pgIds,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	organizationMembers := []organizationMember{}

	for rows.Next() {
		om := organizationMember{}
		err := rows.Scan(
			&om.organizationID,
			&om.personID,
			&om.From,
			&om.Until,
			&om.DateCreated,
			&om.DateUpdated,
			&om.Id,
		)
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		organizationMembers = append(organizationMembers, om)
	}

	return organizationMembers, nil
}

func (repo *repository) getOrganizationParents(ctx context.Context, ids ...int) ([]organizationParent, error) {
	query := `
SELECT "organization_id",
       "parent_organization_id",
	   "from",
	   "until",
	   "date_created",
	   "date_updated",
	   (SELECT "public_id" FROM "organization" WHERE "id" = op.parent_organization_id) parent_organization_public_id
FROM "organization_parent" op
WHERE "organization_id" = any($1)
ORDER by "from" ASC
	`
	pgIds := pgtype.Int4Array{}
	pgIds.Set(ids)
	rows, err := repo.client.QueryContext(
		ctx,
		query,
		pgIds,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	organizationParents := []organizationParent{}

	for rows.Next() {
		op := organizationParent{}
		err := rows.Scan(
			&op.organizationID,
			&op.parentOrganizationID,
			&op.From,
			&op.Until,
			&op.DateCreated,
			&op.DateUpdated,
			&op.Id,
		)
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		organizationParents = append(organizationParents, op)
	}

	return organizationParents, nil
}

func (repo *repository) GetOrganization(ctx context.Context, id string) (*models.Organization, error) {
	row, err := repo.client.Organization.Query().Where(organization.PublicIDEQ(id)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}

	organizationParents, err := repo.getOrganizationParents(ctx, row.ID)
	if err != nil {
		return nil, err
	}

	return repo.orgUnwrap(row, organizationParents), nil
}

func (repo *repository) GetOrganizationByIdentifier(ctx context.Context, typ string, vals ...string) (*models.Organization, error) {
	row, err := repo.client.Organization.Query().Where(func(s *entsql.Selector) {
		s.Where(entsql.ExprP("identifier->$1 ?| $2", typ, vals))
	}).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}

	organizationParents, err := repo.getOrganizationParents(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	return repo.orgUnwrap(row, organizationParents), nil
}

func (repo *repository) GetOrganizationsByIdentifier(ctx context.Context, typ string, vals ...string) ([]*models.Organization, error) {
	rows, err := repo.client.Organization.Query().Where(func(s *entsql.Selector) {
		s.Where(entsql.ExprP("identifier->$1 ?| $2", typ, vals))
	}).All(ctx)

	if err != nil {
		return nil, err
	}

	organizationIds := []int{}
	for _, row := range rows {
		organizationIds = append(organizationIds, row.ID)
	}

	allOrganizationParents, err := repo.getOrganizationParents(ctx, organizationIds...)
	if err != nil {
		return nil, err
	}

	orgs := make([]*models.Organization, 0, len(rows))
	for _, row := range rows {
		organizationParents := []organizationParent{}
		for _, organizationParent := range allOrganizationParents {
			if organizationParent.organizationID == row.ID {
				organizationParents = append(organizationParents, organizationParent)
			}
		}
		orgs = append(orgs, repo.orgUnwrap(row, organizationParents))
	}

	// TODO: order by array_position cannot work on array itself. Find another way
	return orgs, nil
}

func (repo *repository) SaveOrganization(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	if org.IsStored() {
		return repo.UpdateOrganization(ctx, org)
	}
	return repo.CreateOrganization(ctx, org)
}

func (repo *repository) CreateOrganization(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	// date fields filled by schema
	tx, err := repo.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to start transaction: %w", err)
	}
	defer tx.Rollback()

	t := tx.Organization.Create()

	identifiers := schema.TypeVals{}
	for _, id := range org.Identifier {
		identifiers.Add(id.PropertyID, id.Value)
	}
	t.SetIdentifier(identifiers)
	t.SetNameDut(org.NameDut)
	t.SetNameEng(org.NameEng)
	var typ *string
	if org.Type != "" {
		typ = &org.Type
	}
	t.SetNillableType(typ)
	t.SetAcronym(org.Acronym)

	row, err := t.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to save organization: %w", err)
	}

	parentOrganizationPublicIds := []string{}
	for _, parent := range org.Parent {
		parentOrganizationPublicIds = append(parentOrganizationPublicIds, parent.Id)
	}
	parentOrganizationPublicIds = lo.Uniq(parentOrganizationPublicIds)
	parentOrganizationRows, err := repo.client.Organization.Query().Where(organization.PublicIDIn(parentOrganizationPublicIds...)).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(parentOrganizationRows) != len(parentOrganizationPublicIds) {
		return nil, models.ErrInvalidReference
	}

	for _, parent := range org.Parent {
		tParent := tx.OrganizationParent.Create()
		tParent.SetFrom(*parent.From)
		tParent.SetNillableUntil(parent.Until)
		tParent.SetOrganizationID(row.ID)
		for _, parentOrganizationRow := range parentOrganizationRows {
			if parentOrganizationRow.PublicID == parent.Id {
				tParent.SetParentOrganizationID(parentOrganizationRow.ID)
				break
			}
		}
		_, err := tParent.Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("unable to commit transaction: %w", err)
	}

	// collect entgo managed fields
	org.ID = row.PublicID
	org.DateCreated = &row.DateCreated
	org.DateUpdated = &row.DateUpdated

	organizationParents, err := repo.getOrganizationParents(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	for _, organizationParent := range organizationParents {
		org.Parent = append(org.Parent, models.OrganizationParent{
			Id:          organizationParent.Id,
			DateCreated: organizationParent.DateCreated,
			DateUpdated: organizationParent.DateUpdated,
			From:        organizationParent.From,
			Until:       organizationParent.Until,
		})
	}

	return org, nil
}

func (repo *repository) UpdateOrganization(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	tx, err := repo.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to start transaction: %w", err)
	}
	defer tx.Rollback()

	orgRow, err := tx.Organization.Query().Where(organization.PublicIDEQ(org.ID)).ForUpdate(entsql.WithLockAction(entsql.NoWait)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}

	t := tx.Organization.Update().Where(organization.IDEQ(orgRow.ID))

	identifiers := schema.TypeVals{}
	for _, id := range org.Identifier {
		identifiers.Add(id.PropertyID, id.Value)
	}
	t.SetIdentifier(identifiers)
	t.SetNameDut(org.NameDut)
	t.SetNameEng(org.NameEng)
	var typ *string
	if org.Type != "" {
		typ = &org.Type
	}
	t.SetNillableType(typ)
	t.SetAcronym(org.Acronym)

	if _, err := t.Save(ctx); err != nil {
		return nil, fmt.Errorf("unable to save organization: %w", err)
	}

	var newOrganizationParents []organizationParent
	if len(org.Parent) > 0 {
		parentOrganizationPublicIds := []string{}
		for _, parent := range org.Parent {
			parentOrganizationPublicIds = append(parentOrganizationPublicIds, parent.Id)
		}
		parentOrganizationPublicIds = lo.Uniq(parentOrganizationPublicIds)
		parentOrganizationRows, err := repo.client.Organization.Query().Where(organization.PublicIDIn(parentOrganizationPublicIds...)).All(ctx)
		if err != nil {
			return nil, err
		}
		if len(parentOrganizationRows) != len(parentOrganizationPublicIds) {
			return nil, models.ErrInvalidReference
		}
		for _, parent := range org.Parent {
			newOrganizationParent := organizationParent{}
			newOrganizationParent.organizationID = orgRow.ID
			newOrganizationParent.From = parent.From
			newOrganizationParent.Until = parent.Until
			for _, parentOrganizationRow := range parentOrganizationRows {
				if parentOrganizationRow.PublicID == parent.Id {
					newOrganizationParent.parentOrganizationID = parentOrganizationRow.ID
					break
				}
			}
			newOrganizationParents = append(newOrganizationParents, newOrganizationParent)
		}
	}

	updatedRelIds := []int{}
	if len(newOrganizationParents) > 0 {
		for _, newOrganizationParent := range newOrganizationParents {
			op := tx.OrganizationParent.Create()
			op.SetFrom(*newOrganizationParent.From)
			op.SetNillableUntil(newOrganizationParent.Until)
			op.SetParentOrganizationID(newOrganizationParent.parentOrganizationID)
			op.SetOrganizationID(newOrganizationParent.organizationID)
			id, err := op.OnConflictColumns("parent_organization_id", "organization_id", "from").UpdateNewValues().ID(ctx)
			if err != nil {
				return nil, err
			}
			updatedRelIds = append(updatedRelIds, id)
		}
	}
	deleteWhereClause := organizationparent.OrganizationID(orgRow.ID)
	if len(updatedRelIds) > 0 {
		deleteWhereClause = organizationparent.And(deleteWhereClause, organizationparent.IDNotIn(updatedRelIds...))
	}
	_, err = tx.OrganizationParent.Delete().Where(deleteWhereClause).Exec(ctx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("unable to commit transaction: %w", err)
	}

	return repo.GetOrganization(ctx, org.ID)
}

func (repo *repository) DeleteOrganization(ctx context.Context, id string) error {
	_, err := repo.client.Organization.Delete().Where(organization.PublicIDEQ(id)).Exec(ctx)
	return err
}

func (repo *repository) EachOrganization(ctx context.Context, cb func(*models.Organization) bool) error {

	cursor := setCursor{}

	for {
		organizations, newCursor, err := repo.getOrganizations(ctx, cursor)
		if err != nil {
			return err
		}

		for _, organization := range organizations {
			if !cb(organization) {
				return nil
			}
		}

		if len(organizations) == 0 {
			break
		}
		if newCursor.LastID <= 0 {
			break
		}
		cursor = newCursor
	}

	return nil
}

func (repo *repository) SuggestOrganizations(ctx context.Context, query string) ([]*models.Organization, error) {
	tsQuery, tsQueryArgs := toTSQuery(query)
	tsQuery = "ts @@ " + tsQuery
	rows, err := repo.client.Organization.Query().Where(func(s *entsql.Selector) {
		s.Where(
			entsql.ExprP(tsQuery, tsQueryArgs...),
		)
	}).Limit(organizationSuggestLimit).All(ctx)

	if err != nil {
		return nil, err
	}

	organizationIds := []int{}
	for _, row := range rows {
		organizationIds = append(organizationIds, row.ID)
	}
	organizationIds = lo.Uniq(organizationIds)

	allOrganizationParents, err := repo.getOrganizationParents(ctx, organizationIds...)
	if err != nil {
		return nil, err
	}

	orgs := make([]*models.Organization, 0, len(rows))
	for _, row := range rows {
		organizationParents := []organizationParent{}
		for _, organizationParent := range allOrganizationParents {
			if row.ID == organizationParent.organizationID {
				organizationParents = append(organizationParents, organizationParent)
			}
		}
		orgs = append(orgs, repo.orgUnwrap(row, organizationParents))
	}

	return orgs, nil
}

func (repo *repository) GetOrganizations(ctx context.Context) ([]*models.Organization, string, error) {
	organizations, newCursor, err := repo.getOrganizations(ctx, setCursor{})
	if err != nil {
		return nil, "", err
	}

	var encodedCursor string
	if newCursor.LastID > 0 {
		encodedCursor, err = repo.encodeCursor(newCursor)
		if err != nil {
			return nil, "", err
		}
	}
	return organizations, encodedCursor, nil
}

func (repo *repository) GetMoreOrganizations(ctx context.Context, tokenValue string) ([]*models.Organization, string, error) {
	cursor := setCursor{}
	if err := repo.decodeCursor(tokenValue, &cursor); err != nil {
		return nil, "", err
	}
	organizations, newCursor, err := repo.getOrganizations(ctx, cursor)
	if err != nil {
		return nil, "", err
	}

	var encodedCursor string
	if newCursor.LastID > 0 {
		encodedCursor, err = repo.encodeCursor(newCursor)
		if err != nil {
			return nil, "", err
		}
	}

	return organizations, encodedCursor, nil
}

func (repo *repository) getOrganizations(ctx context.Context, cursor setCursor) ([]*models.Organization, setCursor, error) {
	newCursor := setCursor{}
	rows, err := repo.client.Organization.Query().Where(organization.IDGT(cursor.LastID)).Order(ent.Asc(organization.FieldID)).Limit(organizationPageLimit).All(ctx)
	if err != nil {
		return nil, newCursor, err
	}
	if len(rows) == 0 {
		return []*models.Organization{}, setCursor{}, nil
	}

	total, err := repo.client.Organization.Query().Count(ctx)
	if err != nil {
		return nil, newCursor, err
	}

	if total > len(rows) {
		newCursor = setCursor{
			LastID: rows[len(rows)-1].ID,
		}
	}

	organizationIds := []int{}
	for _, row := range rows {
		organizationIds = append(organizationIds, row.ID)
	}
	organizationIds = lo.Uniq(organizationIds)

	allOrganizationParents, err := repo.getOrganizationParents(ctx, organizationIds...)
	if err != nil {
		return nil, newCursor, err
	}

	orgs := make([]*models.Organization, 0, len(rows))
	for _, row := range rows {
		organizationParents := []organizationParent{}
		for _, organizationParent := range allOrganizationParents {
			if row.ID == organizationParent.organizationID {
				organizationParents = append(organizationParents, organizationParent)
			}
		}
		orgs = append(orgs, repo.orgUnwrap(row, organizationParents))
	}
	return orgs, newCursor, nil
}

func (repo *repository) orgUnwrap(organizationRow *ent.Organization, organizationParents []organizationParent) *models.Organization {
	org := &models.Organization{
		ID:          organizationRow.PublicID,
		DateCreated: &organizationRow.DateCreated,
		DateUpdated: &organizationRow.DateUpdated,
		Type:        organizationRow.Type,
		NameDut:     organizationRow.NameDut,
		NameEng:     organizationRow.NameEng,
		Acronym:     organizationRow.Acronym,
	}

	for key, vals := range organizationRow.Identifier {
		for _, val := range vals {
			org.AddIdentifier(key, val)
		}
	}

	for _, organizationParent := range organizationParents {
		org.Parent = append(org.Parent, models.OrganizationParent{
			Id:          organizationParent.Id,
			DateCreated: organizationParent.DateCreated,
			DateUpdated: organizationParent.DateUpdated,
			From:        organizationParent.From,
			Until:       organizationParent.Until,
		})
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
	tx, err := repo.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to start transaction: %w", err)
	}
	defer tx.Rollback()

	t := tx.Person.Create()

	// keep in order; copy to Update if it changes
	t.SetActive(p.Active)
	t.SetBirthDate(p.BirthDate)
	t.SetJobCategory(p.JobCategory)
	t.SetEmail(p.Email)
	t.SetGivenName(p.GivenName)
	t.SetName(p.Name)
	t.SetHonorificPrefix(p.HonorificPrefix)
	t.SetFamilyName(p.FamilyName)
	t.SetRole(p.Role)
	t.SetSettings(p.Settings)
	if len(p.ObjectClass) > 0 {
		t.SetObjectClass(p.ObjectClass)
	} else {
		t.SetObjectClass(nil)
	}
	t.SetExpirationDate(p.ExpirationDate)

	tokens := schema.TypeVals{}
	for _, token := range p.Token {
		eToken, err := encryptMessage(repo.secret, token.Value)
		if err != nil {
			return nil, fmt.Errorf("unable to encrypt %s: %w", token.PropertyID, err)
		}
		tokens.Add(token.PropertyID, eToken)
	}
	t.SetToken(tokens)

	identifiers := schema.TypeVals{}
	for _, id := range p.Identifier {
		identifiers.Add(id.PropertyID, id.Value)
	}
	t.SetIdentifier(identifiers)
	t.SetPreferredGivenName(p.PreferredGivenName)
	t.SetPreferredFamilyName(p.PreferredFamilyName)

	row, err := t.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(p.Organization) > 0 {
		var organizationPublicIds []string
		for _, orgMember := range p.Organization {
			organizationPublicIds = append(organizationPublicIds, orgMember.Id)
		}
		organizationPublicIds = lo.Uniq(organizationPublicIds)
		orgRows, err := tx.Organization.Query().Where(organization.PublicIDIn(organizationPublicIds...)).All(ctx)
		if err != nil {
			return nil, err
		}
		if len(organizationPublicIds) != len(orgRows) {
			return nil, fmt.Errorf("%w: person.organization_id contains invalid organization id's", models.ErrInvalidReference)
		}
		for _, organizationMember := range p.Organization {
			top := tx.OrganizationPerson.Create()
			top.SetFrom(*organizationMember.From)
			top.SetNillableUntil(organizationMember.Until)
			top.SetPersonID(row.ID)
			for _, orgRow := range orgRows {
				if orgRow.PublicID == organizationMember.Id {
					top.SetOrganizationID(orgRow.ID)
					break
				}
			}
			if _, err = top.Save(ctx); err != nil {
				return nil, err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("unable to commit transaction: %w", err)
	}

	return repo.GetPerson(ctx, row.PublicID)
}

func (repo *repository) SetPersonOrcid(ctx context.Context, id string, orcid string) error {
	var sqlRes sql.Result
	var err error

	if orcid == "" {
		sqlRes, err = repo.client.Person.ExecContext(
			ctx,
			"UPDATE person SET date_updated = $1, identifier = identifier - 'orcid'  WHERE public_id = $2",
			time.Now(),
			id,
		)
	} else {
		jsonb, _ := json.Marshal(schema.TypeVals{}.Add("orcid", orcid))
		sqlRes, err = repo.client.Person.ExecContext(
			ctx,
			"UPDATE person SET date_updated = $1, identifier = identifier || $2::jsonb WHERE public_id = $3",
			time.Now(),
			string(jsonb),
			id,
		)
	}
	if err != nil {
		return err
	}

	nUpdated, err := sqlRes.RowsAffected()
	if err != nil {
		return err
	}
	if nUpdated == 0 {
		return models.ErrNotFound
	}
	return nil
}

func (repo *repository) SetPersonOrcidToken(ctx context.Context, id string, orcidToken string) error {
	var uToken string
	var err error
	var sqlRes sql.Result

	if orcidToken == "" {
		sqlRes, err = repo.client.Person.ExecContext(
			ctx,
			"UPDATE person SET date_updated = $1, token = token - 'orcid' WHERE public_id = $2",
			time.Now(),
			id,
		)
	} else {
		uToken, err = encryptMessage(repo.secret, orcidToken)
		if err != nil {
			return fmt.Errorf("unable to encrypt orcid_token: %w", err)
		}
		jsonb, _ := json.Marshal(schema.TypeVals{}.Add("orcid", uToken))
		sqlRes, err = repo.client.Person.ExecContext(
			ctx,
			"UPDATE person SET date_updated = $1, token = token || $2::jsonb WHERE public_id = $3",
			time.Now(),
			string(jsonb),
			id,
		)
	}

	if err != nil {
		return err
	}
	nUpdated, err := sqlRes.RowsAffected()
	if err != nil {
		return err
	}
	if nUpdated == 0 {
		return models.ErrNotFound
	}
	return nil
}

func (repo *repository) UpdatePerson(ctx context.Context, p *models.Person) (*models.Person, error) {
	tx, err := repo.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to start transaction: %w", err)
	}
	defer tx.Rollback()

	// lock row for update
	personRow, err := tx.Person.Query().Where(person.PublicIDEQ(p.ID)).ForUpdate(entsql.WithLockAction(entsql.NoWait)).First(ctx)
	if err != nil {
		return nil, err
	}

	t := tx.Person.Update().Where(person.IDEQ(personRow.ID))

	t.SetActive(p.Active)
	t.SetBirthDate(p.BirthDate)
	t.SetJobCategory(p.JobCategory)
	t.SetEmail(p.Email)
	t.SetGivenName(p.GivenName)
	t.SetName(p.Name)
	t.SetHonorificPrefix(p.HonorificPrefix)
	t.SetFamilyName(p.FamilyName)
	t.SetRole(p.Role)
	t.SetSettings(p.Settings)
	t.SetObjectClass(p.ObjectClass)
	t.SetExpirationDate(p.ExpirationDate)

	tokens := schema.TypeVals{}
	for _, token := range p.Token {
		eToken, err := encryptMessage(repo.secret, token.Value)
		if err != nil {
			return nil, fmt.Errorf("unable to encrypt %s: %w", token.PropertyID, err)
		}
		tokens.Add(token.PropertyID, eToken)
	}
	t.SetToken(tokens)

	identifiers := schema.TypeVals{}
	for _, id := range p.Identifier {
		identifiers.Add(id.PropertyID, id.Value)
	}
	t.SetIdentifier(identifiers)
	t.SetPreferredGivenName(p.PreferredGivenName)
	t.SetPreferredFamilyName(p.PreferredFamilyName)

	var newOrganizationMembers []organizationMember
	if len(p.Organization) > 0 {
		var orgRows []*ent.Organization
		var organizationPublicIds []string
		for _, orgMember := range p.Organization {
			organizationPublicIds = append(organizationPublicIds, orgMember.Id)
		}
		organizationPublicIds = lo.Uniq(organizationPublicIds)
		orgRows, err = tx.Organization.Query().Where(organization.PublicIDIn(organizationPublicIds...)).All(ctx)
		if err != nil {
			return nil, err
		}
		if len(organizationPublicIds) != len(orgRows) {
			return nil, fmt.Errorf("%w: person.organization_id contains invalid organization id's", models.ErrInvalidReference)
		}
		for _, orgMember := range p.Organization {
			newOrganizationMember := organizationMember{
				OrganizationMember: *orgMember,
				personID:           personRow.ID,
			}
			for _, orgRow := range orgRows {
				if orgRow.PublicID == orgMember.Id {
					newOrganizationMember.organizationID = orgRow.ID
					break
				}
			}
			newOrganizationMembers = append(newOrganizationMembers, newOrganizationMember)
		}
	}

	updatedRelIds := []int{}
	if len(newOrganizationMembers) > 0 {
		for _, newOrganizationMember := range newOrganizationMembers {
			top := tx.OrganizationPerson.Create()
			top.SetFrom(*newOrganizationMember.From)
			top.SetNillableUntil(newOrganizationMember.Until)
			top.SetPersonID(newOrganizationMember.personID)
			top.SetOrganizationID(newOrganizationMember.organizationID)
			id, err := top.OnConflictColumns("organization_id", "person_id", "from").UpdateNewValues().ID(ctx)
			if err != nil {
				return nil, err
			}
			updatedRelIds = append(updatedRelIds, id)
		}
	}
	deleteWhereClause := organizationperson.PersonID(personRow.ID)
	if len(updatedRelIds) > 0 {
		deleteWhereClause = organizationperson.And(deleteWhereClause, organizationperson.IDNotIn(updatedRelIds...))
	}
	_, err = tx.OrganizationPerson.Delete().Where(deleteWhereClause).Exec(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := t.Save(ctx); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return repo.GetPerson(ctx, p.ID)
}

func (repo *repository) GetPerson(ctx context.Context, id string) (*models.Person, error) {
	row, err := repo.client.Person.Query().Where(person.PublicIDEQ(id)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}

	personOrganizations, err := repo.getOrganizationMembers(ctx, row.ID)
	if err != nil {
		return nil, err
	}

	return repo.personUnwrap(row, personOrganizations)
}

func (repo *repository) GetPersonByIdentifier(ctx context.Context, typ string, vals ...string) (*models.Person, error) {
	row, err := repo.client.Person.Query().Where(func(s *entsql.Selector) {
		s.Where(entsql.ExprP("identifier->$1 ?| $2", typ, vals))
	}).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}

	personOrganizations, err := repo.getOrganizationMembers(ctx, row.ID)
	if err != nil {
		return nil, err
	}

	return repo.personUnwrap(row, personOrganizations)
}

func (repo *repository) GetPeopleByIdentifier(ctx context.Context, typ string, vals ...string) ([]*models.Person, error) {
	rows, err := repo.client.Person.Query().Where(func(s *entsql.Selector) {
		s.Where(entsql.ExprP("identifier->$1 ?| $2", typ, vals))
	}).All(ctx)
	if err != nil {
		return nil, err
	}

	allPersonIds := make([]int, 0, len(rows))
	for _, row := range rows {
		allPersonIds = append(allPersonIds, row.ID)
	}

	allPersonOrganizationMember, err := repo.getOrganizationMembers(ctx, allPersonIds...)
	if err != nil {
		return nil, err
	}

	people := make([]*models.Person, 0, len(rows))
	for _, row := range rows {
		currentPersonOrganizationMembers := []organizationMember{}
		for _, organizationMember := range allPersonOrganizationMember {
			if organizationMember.personID == row.ID {
				currentPersonOrganizationMembers = append(currentPersonOrganizationMembers, organizationMember)
			}
		}
		person, err := repo.personUnwrap(row, currentPersonOrganizationMembers)
		if err != nil {
			return nil, err
		}
		people = append(people, person)
	}

	return people, nil
}

func (repo *repository) DeletePerson(ctx context.Context, id string) error {
	_, err := repo.client.Person.Delete().Where(person.PublicIDEQ(id)).Exec(ctx)
	return err
}

func (repo *repository) EachPerson(ctx context.Context, cb func(*models.Person) bool) error {
	cursor := setCursor{}

	for {
		people, newCursor, err := repo.getPeople(ctx, cursor)
		if err != nil {
			return err
		}

		for _, person := range people {
			if !cb(person) {
				return nil
			}
		}

		if len(people) == 0 {
			break
		}
		if newCursor.LastID <= 0 {
			break
		}
		cursor = newCursor
	}

	return nil

}

func (repo *repository) SuggestPeople(ctx context.Context, query string) ([]*models.Person, error) {
	// fetch ids via a raw query ..
	tsQuery, tsQueryArgs := toTSQuery(query)
	sqlQuery := fmt.Sprintf(
		"SELECT id, ts_rank(ts, %s) as rank FROM person WHERE ts @@ %s ORDER BY rank DESC LIMIT %d",
		tsQuery,
		tsQuery,
		personSuggestLimit,
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
		Where(person.IDIn(ids...)).
		Order(func(s *entsql.Selector) {
			orderStr := fmt.Sprintf("array_position($%d, id)", len(ids)+1)
			s.OrderExpr(entsql.ExprP(orderStr, ids))
		}).
		All(ctx)

	if err != nil {
		return nil, err
	}

	allPersonIds := make([]int, 0, len(records))
	for _, record := range records {
		allPersonIds = append(allPersonIds, record.ID)
	}
	allPersonOrganizationMembers, err := repo.getOrganizationMembers(ctx, allPersonIds...)
	if err != nil {
		return nil, err
	}

	people := make([]*models.Person, 0, len(records))

	for _, record := range records {
		currentOrganizationMembers := []organizationMember{}
		for _, organizationMember := range allPersonOrganizationMembers {
			if organizationMember.personID == record.ID {
				currentOrganizationMembers = append(currentOrganizationMembers, organizationMember)
			}
		}
		if person, err := repo.personUnwrap(record, currentOrganizationMembers); err != nil {
			return nil, err
		} else {
			people = append(people, person)
		}
	}

	return people, nil
}

func (repo *repository) personUnwrap(entPerson *ent.Person, internalOrganizationMembers []organizationMember) (*models.Person, error) {
	var organizationMembers []*models.OrganizationMember
	for _, iorgMember := range internalOrganizationMembers {
		orgMember := models.NewOrganizationMember(iorgMember.Id)
		orgMember.DateCreated = iorgMember.DateCreated
		orgMember.DateUpdated = iorgMember.DateUpdated
		orgMember.From = iorgMember.From
		orgMember.Until = iorgMember.Until
		organizationMembers = append(organizationMembers, orgMember)
	}

	var tokens []models.Token
	for typ, eTokenVals := range entPerson.Token {
		for _, eTokenVal := range eTokenVals {
			rawTokenVal, err := decryptMessage(repo.secret, eTokenVal)
			if err != nil {
				return nil, fmt.Errorf("unable to decrypt %s token: %w", typ, err)
			}
			tokens = append(tokens, models.Token{
				PropertyID: typ,
				Value:      rawTokenVal,
			})
		}
	}

	p := &models.Person{
		Active:              entPerson.Active,
		BirthDate:           entPerson.BirthDate,
		DateCreated:         &entPerson.DateCreated,
		DateUpdated:         &entPerson.DateUpdated,
		Email:               entPerson.Email,
		ExpirationDate:      entPerson.ExpirationDate,
		GivenName:           entPerson.GivenName,
		Name:                entPerson.Name,
		ID:                  entPerson.PublicID,
		FamilyName:          entPerson.FamilyName,
		JobCategory:         entPerson.JobCategory,
		Token:               tokens,
		Organization:        organizationMembers,
		PreferredFamilyName: entPerson.PreferredFamilyName,
		PreferredGivenName:  entPerson.PreferredGivenName,
		HonorificPrefix:     entPerson.HonorificPrefix,
		Role:                entPerson.Role,
		Settings:            entPerson.Settings,
		ObjectClass:         entPerson.ObjectClass,
	}

	for key, vals := range entPerson.Identifier {
		for _, val := range vals {
			p.AddIdentifier(key, val)
		}
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

func (repo *repository) AutoExpirePeople(ctx context.Context) (int64, error) {
	updateQuery := "UPDATE person SET active = false WHERE expiration_date <= $1 AND active = true"

	res, err := repo.client.ExecContext(ctx, updateQuery, time.Now().Local().Format("20060101"))
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (repo *repository) GetPeople(ctx context.Context) ([]*models.Person, string, error) {
	people, newCursor, err := repo.getPeople(ctx, setCursor{})
	if err != nil {
		return nil, "", err
	}

	var encodedCursor string
	if newCursor.LastID > 0 {
		encodedCursor, err = repo.encodeCursor(newCursor)
		if err != nil {
			return nil, "", err
		}
	}
	return people, encodedCursor, nil
}

func (repo *repository) GetMorePeople(ctx context.Context, tokenValue string) ([]*models.Person, string, error) {
	cursor := setCursor{}
	if err := repo.decodeCursor(tokenValue, &cursor); err != nil {
		return nil, "", err
	}

	people, newCursor, err := repo.getPeople(ctx, cursor)
	if err != nil {
		return nil, "", err
	}

	var encodedCursor string
	if newCursor.LastID > 0 {
		encodedCursor, err = repo.encodeCursor(newCursor)
		if err != nil {
			return nil, "", err
		}
	}
	return people, encodedCursor, nil
}

func (repo *repository) getPeople(ctx context.Context, cursor setCursor) ([]*models.Person, setCursor, error) {
	newCursor := setCursor{}
	rows, err := repo.client.Person.Query().Where(person.IDGT(cursor.LastID)).Order(ent.Asc(person.FieldID)).Limit(personPageLimit).All(ctx)
	if err != nil {
		return nil, newCursor, err
	}
	if len(rows) == 0 {
		return []*models.Person{}, setCursor{}, nil
	}

	allPersonIds := make([]int, 0, len(rows))
	for _, row := range rows {
		allPersonIds = append(allPersonIds, row.ID)
	}

	allPersonOrganizationMembers, err := repo.getOrganizationMembers(ctx, allPersonIds...)
	if err != nil {
		return nil, newCursor, err
	}

	people := make([]*models.Person, 0, len(rows))
	for _, row := range rows {
		currentPersonOrganizationMembers := []organizationMember{}
		for _, personOrganizationMember := range allPersonOrganizationMembers {
			if personOrganizationMember.personID == row.ID {
				currentPersonOrganizationMembers = append(currentPersonOrganizationMembers, personOrganizationMember)
			}
		}
		person, e := repo.personUnwrap(row, currentPersonOrganizationMembers)
		if e != nil {
			return nil, newCursor, e
		}
		people = append(people, person)
	}

	total, err := repo.client.Person.Query().Count(ctx)
	if err != nil {
		return nil, newCursor, err
	}

	if total > len(rows) {
		newCursor = setCursor{
			LastID: rows[len(rows)-1].ID,
		}
	}

	return people, newCursor, nil
}

func (repo *repository) encodeCursor(c any) (string, error) {
	plaintext, _ := json.Marshal(c)
	ciphertext, err := crypt.Encrypt(repo.secret, plaintext)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func (repo *repository) decodeCursor(encryptedCursor string, c any) error {
	ciphertext, err := base64.URLEncoding.DecodeString(encryptedCursor)
	if err != nil {
		return err
	}
	plaintext, err := crypt.Decrypt(repo.secret, ciphertext)
	if err != nil {
		return err
	}
	return json.Unmarshal(plaintext, c)
}
