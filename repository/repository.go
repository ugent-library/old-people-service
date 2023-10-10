package repository

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/jackc/pgx/v5"
	"github.com/ugent-library/crypt"
	"github.com/ugent-library/people-service/ent"
	"github.com/ugent-library/people-service/ent/organization"
	"github.com/ugent-library/people-service/ent/person"
	"github.com/ugent-library/people-service/ent/schema"
	"github.com/ugent-library/people-service/models"
)

type repository struct {
	client *ent.Client
	secret []byte
}
type setCursor struct {
	// IMPORTANT: auto increment (of id) starts with 1, so default value 0 should never match
	LastID int `json:"l"`
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

func (repo *repository) GetOrganization(ctx context.Context, id string) (*models.Organization, error) {
	row, err := repo.client.Organization.Query().WithParent().Where(organization.PublicIDEQ(id)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return repo.orgUnwrap(row), nil
}

func (repo *repository) GetOrganizationByIdentifier(ctx context.Context, typ string, vals ...string) (*models.Organization, error) {
	row, err := repo.client.Organization.Query().WithParent().Where(func(s *entsql.Selector) {
		s.Where(entsql.ExprP("identifier->$1 ?| $2", typ, vals))
	}).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return repo.orgUnwrap(row), nil
}

func (repo *repository) GetOrganizationsByIdentifier(ctx context.Context, typ string, vals ...string) ([]*models.Organization, error) {
	rows, err := repo.client.Organization.Query().WithParent().Where(func(s *entsql.Selector) {
		s.Where(entsql.ExprP("identifier->$1 ?| $2", typ, vals))
	}).All(ctx)

	if err != nil {
		return nil, err
	}
	// TODO: order by array_position cannot work on array itself. Find another way
	orgs := make([]*models.Organization, 0, len(rows))
	for _, row := range rows {
		orgs = append(orgs, repo.orgUnwrap(row))
	}

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
	t.SetType(org.Type)
	if org.ParentID != "" {
		parentOrgRow, err := tx.Organization.Query().Where(organization.PublicIDEQ(org.ParentID)).First(ctx)
		if err != nil {
			var e *ent.NotFoundError
			if errors.As(err, &e) {
				return nil, fmt.Errorf("%w: parent organization with public_id %s not found", models.ErrInvalidReference, org.ParentID)
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
	org.ID = row.PublicID
	org.DateCreated = &row.DateCreated
	org.DateUpdated = &row.DateUpdated

	return org, nil
}

func (repo *repository) UpdateOrganization(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	tx, err := repo.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to start transaction: %w", err)
	}
	defer tx.Rollback()

	t := tx.Organization.Update().Where(organization.PublicIDEQ(org.ID))

	identifiers := schema.TypeVals{}
	for _, id := range org.Identifier {
		identifiers.Add(id.PropertyID, id.Value)
	}
	t.SetIdentifier(identifiers)
	t.SetNameDut(org.NameDut)
	t.SetNameEng(org.NameEng)
	t.SetType(org.Type)
	t.ClearParent()
	if org.ParentID != "" {
		parentOrg, err := tx.Organization.Query().Where(organization.PublicIDEQ(org.ParentID)).First(ctx)
		if err != nil {
			var e *ent.NotFoundError
			if errors.As(err, &e) {
				return nil, fmt.Errorf("%w: parent organization with public_id %s not found", models.ErrInvalidReference, org.ParentID)
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
	row, err := tx.Organization.Query().Where(organization.PublicIDEQ(org.ID)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to query organizations: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("unable to commit transaction: %w", err)
	}

	return repo.orgUnwrap(row), nil
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
		orgs = append(orgs, repo.orgUnwrap(row))
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
	limit := 100
	newCursor := setCursor{}
	rows, err := repo.client.Organization.Query().Where(organization.IDGT(cursor.LastID)).Order(ent.Asc(organization.FieldID)).WithParent().Limit(limit).All(ctx)
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

	orgs := make([]*models.Organization, 0, len(rows))
	for _, row := range rows {
		orgs = append(orgs, repo.orgUnwrap(row))
	}
	return orgs, newCursor, nil
}

func (repo *repository) orgUnwrap(e *ent.Organization) *models.Organization {
	org := &models.Organization{
		ID:          e.PublicID,
		DateCreated: &e.DateCreated,
		DateUpdated: &e.DateUpdated,
		Type:        e.Type,
		NameDut:     e.NameDut,
		NameEng:     e.NameEng,
	}

	for key, vals := range e.Identifier {
		for _, val := range vals {
			org.AddIdentifier(key, val)
		}
	}

	if parentOrg := e.Edges.Parent; parentOrg != nil {
		org.ParentID = parentOrg.PublicID
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

	if len(p.Organization) > 0 {
		var organizationId []string
		for _, orgRef := range p.Organization {
			organizationId = append(organizationId, orgRef.Id)
		}
		// TODO: crashes with segmentation violation error when org does not exist
		orgs, err := tx.Organization.Query().Where(organization.PublicIDIn(organizationId...)).All(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to query organizations: %w", err)
		}
		if len(p.Organization) != len(orgs) {
			return nil, fmt.Errorf("%w: 's", models.ErrInvalidReference)
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

	// reload everything because you also added new organization references
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

	t := tx.Person.Update().Where(person.PublicIDEQ(p.ID))

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
	t.ClearOrganizations()
	if len(p.Organization) > 0 {
		var organizationId []string
		for _, orgRef := range p.Organization {
			organizationId = append(organizationId, orgRef.Id)
		}
		// TODO: crashes with segmentation violation error when org does not exist
		orgs, err := tx.Organization.Query().Where(organization.PublicIDIn(organizationId...)).All(ctx)
		if err != nil {
			return nil, err
		}
		if len(p.Organization) != len(orgs) {
			return nil, fmt.Errorf("%w: person.organization_id contains invalid organization id's", models.ErrInvalidReference)
		}
		t.AddOrganizations(orgs...)
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
	row, err := repo.client.Person.Query().WithOrganizations().WithOrganizationPerson().Where(person.PublicIDEQ(id)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}

	return repo.personUnwrap(row)
}

func (repo *repository) GetPersonByIdentifier(ctx context.Context, typ string, vals ...string) (*models.Person, error) {
	row, err := repo.client.Person.Query().WithOrganizations().WithOrganizationPerson().Where(func(s *entsql.Selector) {
		s.Where(entsql.ExprP("identifier->$1 ?| $2", typ, vals))
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
		WithOrganizationPerson().
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
	var orgRefs []*models.OrganizationRef
	for _, orgRow := range e.Edges.Organizations {
		var thisOrgPersonRow *ent.OrganizationPerson
		for _, orgPersonRow := range e.Edges.OrganizationPerson {
			if orgPersonRow.OrganizationID == orgRow.ID {
				thisOrgPersonRow = orgPersonRow
				break
			}
		}
		orgRef := models.NewOrganizationRef(orgRow.PublicID)
		orgRef.DateCreated = &thisOrgPersonRow.DateCreated
		orgRef.DateUpdated = &thisOrgPersonRow.DateUpdated
		orgRef.From = &thisOrgPersonRow.From
		orgRef.Until = &thisOrgPersonRow.Until
		orgRefs = append(orgRefs, orgRef)
	}
	sort.SliceStable(orgRefs, func(i, j int) bool {
		return orgRefs[i].DateCreated.Before(*orgRefs[j].DateCreated)
	})

	var tokens []models.Token
	for typ, eTokenVals := range e.Token {
		for _, eTokenVal := range eTokenVals {
			rawTokenVal, err := decryptMessage(repo.secret, eTokenVal)
			if err != nil {
				return nil, fmt.Errorf("unable to decrypt % token: %w", typ, err)
			}
			tokens = append(tokens, models.Token{
				PropertyID: typ,
				Value:      rawTokenVal,
			})
		}
	}

	p := &models.Person{
		Active:              e.Active,
		BirthDate:           e.BirthDate,
		DateCreated:         &e.DateCreated,
		DateUpdated:         &e.DateUpdated,
		Email:               e.Email,
		GivenName:           e.GivenName,
		Name:                e.Name,
		ID:                  e.PublicID,
		FamilyName:          e.FamilyName,
		JobCategory:         e.JobCategory,
		Token:               tokens,
		Organization:        orgRefs,
		PreferredFamilyName: e.PreferredFamilyName,
		PreferredGivenName:  e.PreferredGivenName,
		HonorificPrefix:     e.HonorificPrefix,
		Role:                e.Role,
		Settings:            e.Settings,
		ObjectClass:         e.ObjectClass,
	}

	for key, vals := range e.Identifier {
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
	limit := 100
	newCursor := setCursor{}
	rows, err := repo.client.Person.Query().Where(person.IDGT(cursor.LastID)).Order(ent.Asc(person.FieldID)).WithOrganizations().WithOrganizationPerson().Limit(limit).All(ctx)
	if err != nil {
		return nil, newCursor, err
	}
	if len(rows) == 0 {
		return []*models.Person{}, setCursor{}, nil
	}

	people := make([]*models.Person, 0, len(rows))
	for _, row := range rows {
		person, e := repo.personUnwrap(row)
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
