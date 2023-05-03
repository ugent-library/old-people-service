package models

type Services struct {
	PersonService              PersonService
	PersonSuggestService       PersonSuggestService
	OrganizationService        OrganizationService
	OrganizationSuggestService OrganizationSuggestService
}
