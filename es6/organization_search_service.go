package es6

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/elastic/go-elasticsearch/v6/esapi"
	"github.com/pkg/errors"
	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrganizationSearchService struct {
	Client
}

func (os *OrganizationSearchService) Search(searchArgs *models.SearchArgs) (*models.OrganizationHits, error) {
	req := es6SearchReq{
		From:  searchArgs.Start,
		Size:  searchArgs.Limit,
		Query: searchArgs.Query,
		Sort:  searchArgs.Sort,
	}
	res := es6SearchRes{}
	if err := os.Client.Search(&req, &res); err != nil {
		return nil, fmt.Errorf("es6 search error: %w", err)
	}

	hits := &models.OrganizationHits{}
	hits.Start = searchArgs.Start
	hits.Limit = searchArgs.Limit
	hits.Total = res.Hits.Total

	for _, h := range res.Hits.Hits {
		var io indexedOrganization = indexedOrganization{}
		if err := json.Unmarshal(h.Source, &io); err != nil {
			return nil, fmt.Errorf("unable to decode es6 hit: %w", err)
		}
		hits.Hits = append(hits.Hits, indexedOrganizationToModel(&io))
	}

	return hits, nil
}

func (os *OrganizationSearchService) Index(org *models.Organization) error {
	payload, err := json.Marshal(newIndexedOrganization(org))
	if err != nil {
		return err
	}
	ctx := context.Background()
	res, err := esapi.IndexRequest{
		Index:      os.Client.Index,
		DocumentID: org.Id,
		Body:       bytes.NewReader(payload),
	}.Do(ctx, os.Client.es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		buf := &bytes.Buffer{}
		if _, err := io.Copy(buf, res.Body); err != nil {
			return err
		}
		return errors.New("Es6 error response: " + buf.String())
	}

	return nil
}

func (os *OrganizationSearchService) Delete(id string) error {
	ctx := context.Background()
	res, err := esapi.DeleteRequest{
		Index:      os.Client.Index,
		DocumentID: id,
	}.Do(ctx, os.Client.es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		buf := &bytes.Buffer{}
		if _, err := io.Copy(buf, res.Body); err != nil {
			return err
		}
		return errors.New("Es6 error response: " + buf.String())
	}

	return nil
}

func (os *OrganizationSearchService) DeleteAll() error {
	ctx := context.Background()
	req := esapi.DeleteByQueryRequest{
		Index: []string{os.Client.Index},
		Body: strings.NewReader(`{
            "query" : { 
                "match_all" : {}
            }
        }`),
	}
	res, err := req.Do(ctx, os.Client.es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		buf := &bytes.Buffer{}
		if _, err := io.Copy(buf, res.Body); err != nil {
			return err
		}
		return errors.New("Es6 error response: " + buf.String())
	}

	return nil
}

func (os *OrganizationSearchService) NewBulkIndexer(config models.BulkIndexerConfig) (models.BulkIndexer[*models.Organization], error) {
	docFn := func(org *models.Organization) (string, []byte, error) {
		doc, err := json.Marshal(newIndexedOrganization(org))
		return org.Id, doc, err
	}
	return newBulkIndexer(os.Client.es, os.Client.Index, docFn, config)
}

func (os *OrganizationSearchService) NewIndexSwitcher(config models.BulkIndexerConfig) (models.IndexSwitcher[*models.Organization], error) {
	docFn := func(org *models.Organization) (string, []byte, error) {
		doc, err := json.Marshal(newIndexedOrganization(org))
		return org.Id, doc, err
	}
	return newIndexSwitcher(os.Client.es, os.Client.Index,
		os.Client.Settings, os.Client.IndexRetention, docFn, config)
}

func (os *OrganizationSearchService) Suggest(query string) ([]*models.Organization, error) {

	limit := 500

	// remove duplicate spaces
	query = regexMultipleSpaces.ReplaceAllString(query, " ")

	// trim
	query = strings.TrimSpace(query)

	qParts := strings.Split(query, " ")
	queryMust := make([]models.M, 0, len(qParts))

	for _, qp := range qParts {

		// remove terms that contain brackets
		if regexNoBrackets.MatchString(qp) {
			continue
		}

		queryMust = append(queryMust, models.M{
			"query_string": models.M{
				"default_operator": "AND",
				"query":            fmt.Sprintf("%s*", qp),
				"default_field":    "all",
				"analyze_wildcard": "true",
			},
		})
	}

	searchArgs := &models.SearchArgs{
		Start: 0,
		Limit: limit,
		Query: models.M{
			"bool": models.M{
				"must": queryMust,
			},
		},
	}

	hits, err := os.Search(searchArgs)
	if err != nil {
		return nil, err
	}

	return hits.Hits, nil
}

func indexedOrganizationToModel(io *indexedOrganization) *models.Organization {
	dc, _ := ParseTimeUTC(io.DateCreated)
	du, _ := ParseTimeUTC(io.DateUpdated)
	org := &models.Organization{
		Organization: v1.Organization{
			Id:          io.Id,
			Type:        io.Type,
			DateCreated: timestamppb.New(*dc),
			DateUpdated: timestamppb.New(*du),
			NameDut:     io.NameDut,
			NameEng:     io.NameEng,
			ParentId:    io.ParentId,
			OtherId:     make([]*v1.IdRef, 0),
		},
	}
	for idType, idValues := range io.OtherId {
		for _, idVal := range idValues {
			org.OtherId = append(org.OtherId, &v1.IdRef{
				Type: idType,
				Id:   idVal,
			})
		}
	}
	return org
}

func newIndexedOrganization(org *models.Organization) *indexedOrganization {

	dateCreated := org.DateCreated.AsTime()
	dateUpdated := org.DateUpdated.AsTime()

	// TODO: fill "tree"
	io := &indexedOrganization{
		Id:          org.Id,
		Type:        org.Type,
		ParentId:    org.ParentId,
		NameDut:     org.NameDut,
		NameEng:     org.NameEng,
		DateCreated: formatTimeUTC(&dateCreated),
		DateUpdated: formatTimeUTC(&dateUpdated),
		OtherId:     map[string][]string{},
	}

	for _, otherId := range org.OtherId {
		idList, ok := io.OtherId[otherId.Type]
		if !ok {
			idList = make([]string, 0)
		}
		idList = append(idList, otherId.Id)
		io.OtherId[otherId.Type] = idList
	}

	return io
}
