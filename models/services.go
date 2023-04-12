package models

type Services struct {
	PersonService             PersonService
	PersonSearchService       PersonSearchService
	OrganizationService       OrganizationService
	OrganizationSearchService OrganizationSearchService
}
