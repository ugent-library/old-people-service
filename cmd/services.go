package cmd

import (
	"sync"

	"github.com/ugent-library/person-service/models"
	"github.com/ugent-library/person-service/repository"
	"github.com/ugent-library/person-service/ugentldap"
)

var (
	_services       *models.Services
	_servicesOnce   sync.Once
	_ldapClient     *ugentldap.UgentLdap
	_ldapClientOnce sync.Once
)

func Services() *models.Services {
	_servicesOnce.Do(func() {
		_services = newServices()
	})
	return _services
}

func LDAPClient() *ugentldap.UgentLdap {
	_ldapClientOnce.Do(func() {
		_ldapClient = ugentldap.NewClient(ugentldap.Config{
			Url:      config.Ldap.Url,
			Username: config.Ldap.Username,
			Password: config.Ldap.Password,
		})
	})
	return _ldapClient
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
