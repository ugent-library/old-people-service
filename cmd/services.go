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

	personService, err := db.NewPersonService(&db.PersonConfig{
		DB: config.Db.Url,
	})

	if err != nil {
		panic(err)
	}

	organizationService, err := db.NewOrganizationService(&db.OrganizationConfig{
		DB: config.Db.Url,
	})

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
