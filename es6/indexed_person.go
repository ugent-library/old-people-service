package es6

import (
	"time"

	"github.com/ugent-library/people/models"
)

// time.RFC3339 does not include milliseconds
const TimeFormatUTC = "2006-01-02T15:04:05.999Z"

// force UTC
// force use of milliseconds
func formatTimeUTC(t *time.Time) string {
	return t.UTC().Format(TimeFormatUTC)
}
func ParseTimeUTC(ds string) (*time.Time, error) {
	t, e := time.Parse(TimeFormatUTC, ds)
	if e != nil {
		return nil, e
	}
	return &t, nil
}

type indexedPerson struct {
	*models.Person
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

func newIndexedPerson(person *models.Person) *indexedPerson {
	dateCreated := person.DateCreated.AsTime()
	dateUpdated := person.DateUpdated.AsTime()

	return &indexedPerson{
		Person:      person,
		DateCreated: formatTimeUTC(&dateCreated),
		DateUpdated: formatTimeUTC(&dateUpdated),
	}
}
