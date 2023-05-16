package cmd

import (
	"sync"

	"github.com/ugent-library/people/db"
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

func newServices() *models.Services {

	dbClient, err := db.OpenClient(config.Db.Url)
	if err != nil {
		panic(err)
	}

	dbConfig := &db.Config{
		Client: dbClient,
		AesKey: config.Db.AesKey,
	}

	personService, err := db.NewPersonService(dbConfig)

	if err != nil {
		panic(err)
	}

	organizationService, err := db.NewOrganizationService(dbConfig)

	if err != nil {
		panic(err)
	}

	return &models.Services{
		PersonService:              personService,
		PersonSuggestService:       personService,
		OrganizationService:        organizationService,
		OrganizationSuggestService: organizationService,
	}
}
