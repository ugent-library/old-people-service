package cmd

import (
	"sync"

	"github.com/ugent-library/people/models"
	"github.com/ugent-library/people/repository"
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

	dbClient, err := repository.OpenClient(config.Db.Url)
	if err != nil {
		panic(err)
	}

	dbConfig := &repository.Config{
		Client: dbClient,
		AesKey: config.Db.AesKey,
	}

	personService, err := repository.NewPersonService(dbConfig)

	if err != nil {
		panic(err)
	}

	organizationService, err := repository.NewOrganizationService(dbConfig)

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
