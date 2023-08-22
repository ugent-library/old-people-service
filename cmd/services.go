package cmd

import (
	"sync"

	"github.com/ugent-library/people-service/models"
	"github.com/ugent-library/people-service/repository"
	"github.com/ugent-library/people-service/ugentldap"
)

var (
	_repository     models.Repository
	_repositoryOnce sync.Once
	_ldapClient     *ugentldap.UgentLdap
	_ldapClientOnce sync.Once
)

func Repository() models.Repository {
	_repositoryOnce.Do(func() {
		_repository = newRepository()
	})
	return _repository
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

func newRepository() models.Repository {
	repo, err := repository.NewRepository(&repository.Config{
		DbUrl:  config.Db.Url,
		AesKey: config.Db.AesKey,
	})

	if err != nil {
		panic(err)
	}

	return repo
}
