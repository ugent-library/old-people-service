package es6

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/elastic/go-elasticsearch/v6/esapi"
	"github.com/pkg/errors"
	"github.com/ugent-library/people/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var regexMultipleSpaces = regexp.MustCompile(`\s+`)
var regexNoBrackets = regexp.MustCompile(`[\[\]()\{\}]`)

type PersonSearchService struct {
	Client
}

func (ps *PersonSearchService) Search(searchArgs *models.SearchArgs) (*models.PersonHits, error) {
	req := es6SearchReq{
		From:  searchArgs.Start,
		Size:  searchArgs.Limit,
		Query: searchArgs.Query,
		Sort:  searchArgs.Sort,
	}
	res := es6SearchRes{}
	if err := ps.Client.Search(&req, &res); err != nil {
		return nil, fmt.Errorf("es6 search error: %w", err)
	}

	hits := &models.PersonHits{}
	hits.Start = searchArgs.Start
	hits.Limit = searchArgs.Limit
	hits.Total = res.Hits.Total

	for _, h := range res.Hits.Hits {
		var hit indexedPerson = indexedPerson{}
		if err := json.Unmarshal(h.Source, &hit); err != nil {
			return nil, fmt.Errorf("unable to decode es6 hit: %w", err)
		}
		person := hit.Person
		dc, _ := ParseTimeUTC(hit.DateCreated)
		du, _ := ParseTimeUTC(hit.DateUpdated)
		person.DateCreated = timestamppb.New(*dc)
		person.DateUpdated = timestamppb.New(*du)
		hits.Hits = append(hits.Hits, person)
	}

	return hits, nil
}

func (ps *PersonSearchService) Index(person *models.Person) error {
	payload, err := json.Marshal(newIndexedPerson(person))
	if err != nil {
		return err
	}
	ctx := context.Background()
	res, err := esapi.IndexRequest{
		Index:      ps.Client.Index,
		DocumentID: person.Id,
		Body:       bytes.NewReader(payload),
	}.Do(ctx, ps.Client.es)
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

func (ps *PersonSearchService) Delete(id string) error {
	ctx := context.Background()
	res, err := esapi.DeleteRequest{
		Index:      ps.Client.Index,
		DocumentID: id,
	}.Do(ctx, ps.Client.es)
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

func (ps *PersonSearchService) DeleteAll() error {
	ctx := context.Background()
	req := esapi.DeleteByQueryRequest{
		Index: []string{ps.Client.Index},
		Body: strings.NewReader(`{
            "query" : { 
                "match_all" : {}
            }
        }`),
	}
	res, err := req.Do(ctx, ps.Client.es)
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

func (ps *PersonSearchService) NewBulkIndexer(config models.BulkIndexerConfig) (models.BulkIndexer[*models.Person], error) {
	docFn := func(p *models.Person) (string, []byte, error) {
		doc, err := json.Marshal(newIndexedPerson(p))
		return p.Id, doc, err
	}
	return newBulkIndexer(ps.Client.es, ps.Client.Index, docFn, config)
}

func (ps *PersonSearchService) NewIndexSwitcher(config models.BulkIndexerConfig) (models.IndexSwitcher[*models.Person], error) {
	docFn := func(p *models.Person) (string, []byte, error) {
		doc, err := json.Marshal(newIndexedPerson(p))
		return p.Id, doc, err
	}
	return newIndexSwitcher(ps.Client.es, ps.Client.Index,
		ps.Client.Settings, ps.Client.IndexRetention, docFn, config)
}

func (ps *PersonSearchService) Suggest(query string) ([]*models.Person, error) {

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

	hits, err := ps.Search(searchArgs)
	if err != nil {
		return nil, err
	}

	return hits.Hits, nil
}
