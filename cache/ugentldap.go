package cache

import (
	"fmt"
	"os"
	"time"

	"github.com/bluele/gcache"
	"github.com/go-ldap/ldap/v3"
	"github.com/ugent-library/old-people-service/ugentldap"
)

type ugentLdapClientWithCache struct {
	searcher ugentldap.Searcher
	cache    gcache.Cache
}

func NewUgentLdapSearcher(client ugentldap.Searcher, size int, expiration time.Duration) ugentldap.Searcher {
	cache := gcache.New(size).
		Expiration(expiration).
		LRU().
		LoaderFunc(func(key any) (any, error) {
			var entries []*ldap.Entry
			err := client.SearchPeople(key.(string), func(e *ldap.Entry) error {
				entries = append(entries, e)
				return nil
			})
			return entries, err
		}).
		AddedFunc(func(key, val any) {
			fmt.Fprintf(os.Stderr, "added entry %s\n", key)
		}).
		Build()

	return &ugentLdapClientWithCache{
		searcher: client,
		cache:    cache,
	}
}

func (c *ugentLdapClientWithCache) SearchPeople(filter string, cb func(e *ldap.Entry) error) error {
	entries, err := c.cache.Get(filter)
	if err != nil {
		return err
	}
	for _, entry := range entries.([]*ldap.Entry) {
		if err := cb(entry); err != nil {
			return nil
		}
	}
	return nil
}
