package repository

import (
	"context"
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

func (repo *repository) GetOrganizationByGismoId(ctx context.Context, gismoId string) (*models.Organization, error) {
	row, err := repo.client.Organization.Query().WithParent().Where(organization.GismoIDEQ(gismoId)).First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return repo.orgUnwrap(row), nil
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
		orgs = append(orgs, repo.orgUnwrap(row))
	}

	return orgs, nil
}

func (repo *repository) GetOrganizationByAnyOtherId(ctx context.Context, typ string, values ...string) (*models.Organization, error) {
	row, err := repo.client.Organization.Query().WithParent().Where(func(s *entsql.Selector) {
		s.Where(entsql.ExprP("other_id->$1 ?| $2", typ, values))
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
	t.SetOtherID(schema.IdRefs(org.OtherId))
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
	org.DateCreated = &row.DateCreated
	org.DateUpdated = &row.DateUpdated

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
	t.SetOtherID(schema.IdRefs(org.OtherId))
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

	encodedCursor, err := repo.encodeCursor(newCursor)
	if err != nil {
		return nil, "", err
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

	encodedCursor, err := repo.encodeCursor(newCursor)
	if err != nil {
		return nil, "", err
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
	var gismoId string = ""
	if e.GismoID != nil {
		gismoId = *e.GismoID
	}
	org := &models.Organization{
		Id:          e.PublicID,
		GismoId:     gismoId,
		DateCreated: &e.DateCreated,
		DateUpdated: &e.DateUpdated,
		Type:        e.Type,
		NameDut:     e.NameDut,
		NameEng:     e.NameEng,
		OtherId:     models.IdRefs(e.OtherID),
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

	t.SetOtherID(schema.IdRefs(p.OtherId))
	t.SetPreferredFirstName(p.PreferredFirstName)
	t.SetPreferredLastName(p.PreferredLastName)

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
	p.DateCreated = &row.DateCreated
	p.DateUpdated = &row.DateUpdated
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

	t.SetOtherID(schema.IdRefs(p.OtherId))
	t.SetPreferredFirstName(p.PreferredFirstName)
	t.SetPreferredLastName(p.PreferredLastName)
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

func (repo *repository) GetPersonByAnyOtherId(ctx context.Context, typ string, values ...string) (*models.Person, error) {
	row, err := repo.client.Person.Query().WithOrganizations().WithOrganizationPerson().Where(func(s *entsql.Selector) {
		s.Where(entsql.ExprP("other_id->$1 ?| $2", typ, values))
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
	row, err := repo.client.Person.Query().WithOrganizations().WithOrganizationPerson().Where(person.GismoID(gismoId)).First(ctx)
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
		Active:             e.Active,
		BirthDate:          e.BirthDate,
		DateCreated:        &e.DateCreated,
		DateUpdated:        &e.DateUpdated,
		Email:              e.Email,
		OtherId:            models.IdRefs(e.OtherID),
		FirstName:          e.FirstName,
		FullName:           e.FullName,
		Id:                 e.PublicID,
		GismoId:            gismoId,
		LastName:           e.LastName,
		JobCategory:        e.JobCategory,
		Orcid:              e.Orcid,
		OrcidToken:         uToken,
		Organization:       orgRefs,
		PreferredLastName:  e.PreferredLastName,
		PreferredFirstName: e.PreferredFirstName,
		Title:              e.Title,
		Role:               e.Role,
		Settings:           e.Settings,
		ObjectClass:        e.ObjectClass,
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

	encodedCursor, err := repo.encodeCursor(newCursor)
	if err != nil {
		return nil, "", err
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

	encodedCursor, err := repo.encodeCursor(newCursor)
	if err != nil {
		return nil, "", err
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
