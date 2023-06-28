package cmd

import (
	"sync"

	"github.com/ugent-library/person-service/models"
	"github.com/ugent-library/person-service/repository"
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

	repo, err := repository.NewRepository(dbConfig)

	if err != nil {
		panic(err)
	}

	return &models.Services{
		Repository: repo,
	}
}
