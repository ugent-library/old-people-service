package cmd

import (
	"fmt"
	"os"
	"sync"

	"github.com/elastic/go-elasticsearch/v6"
	"github.com/ugent-library/people/db"
	"github.com/ugent-library/people/es6"
	"github.com/ugent-library/people/models"
)

var (
	_services     *models.Services
	_servicesOnce sync.Once
)

func Services() *models.Services {
	_servicesOnce.Do(func() {
		_services = newServices()
	})
	return _services
}

func newPersonService() (models.PersonService, error) {
	return db.NewPersonService(&db.PersonConfig{
		DB: config.Db.Url,
	})
}

func newOrganizationService() (models.OrganizationService, error) {
	return db.NewOrganizationService(&db.OrganizationConfig{
		DB: config.Db.Url,
	})
}

func newPersonSearchService() (models.PersonSearchService, error) {
	filePath := "etc/es6/authority_person.json"
	settings, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable read from file %s: %w", filePath, err)
	}
	es6Client, err := es6.NewClient(es6.Config{
		Index:    "authority_person",
		Settings: string(settings),
		ClientConfig: elasticsearch.Config{
			Addresses: []string{
				"http://localhost:9200",
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("unable to initialize es6 client: %w", err)
	}
	return &es6.PersonSearchService{
		Client: *es6Client,
	}, nil
}

func newOrganizationSearchService() (models.OrganizationSearchService, error) {
	filePath := "etc/es6/authority_organization.json"
	settings, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable read from file %s: %w", filePath, err)
	}
	es6Client, err := es6.NewClient(es6.Config{
		Index:    "authority_organization",
		Settings: string(settings),
		ClientConfig: elasticsearch.Config{
			Addresses: []string{
				"http://localhost:9200",
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("unable to initialize es6 client: %w", err)
	}
	return &es6.OrganizationSearchService{
		Client: *es6Client,
	}, nil
}

func newServices() *models.Services {

	personService, err := newPersonService()
	if err != nil {
		panic(err)
	}

	personSearchService, err := newPersonSearchService()
	if err != nil {
		panic(err)
	}

	organizationService, err := newOrganizationService()
	if err != nil {
		panic(err)
	}

	organizationSearchService, err := newOrganizationSearchService()
	if err != nil {
		panic(err)
	}

	return &models.Services{
		PersonService:             personService,
		PersonSearchService:       personSearchService,
		OrganizationService:       organizationService,
		OrganizationSearchService: organizationSearchService,
	}
}
