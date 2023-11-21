package cli

import (
	"github.com/ugent-library/old-people-service/jetstreamclient"
	"github.com/ugent-library/old-people-service/models"
	"github.com/ugent-library/old-people-service/repository"
	"github.com/ugent-library/old-people-service/ugentldap"
)

func newRepository() (models.Repository, error) {
	return repository.NewRepository(&repository.Config{
		DbUrl:  config.Db.Url,
		AesKey: config.Db.AesKey,
	})
}

func newUgentLdapClient() *ugentldap.Client {
	return ugentldap.NewClient(ugentldap.Config{
		Url:      config.Ldap.Url,
		Username: config.Ldap.Username,
		Password: config.Ldap.Password,
	})
}

func newJetstreamClient() (*jetstreamclient.Client, error) {
	return jetstreamclient.New(&jetstreamclient.Config{
		NatsUrl:    config.Nats.Url,
		StreamName: config.Nats.StreamName,
		Nkey:       config.Nats.Nkey,
		NkeySeed:   config.Nats.NkeySeed,
		Logger:     logger,
	})
}
