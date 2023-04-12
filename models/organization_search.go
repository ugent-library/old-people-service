package models

type OrganizationHits struct {
	Start int             `json:"start"`
	Limit int             `json:"limit"`
	Total int             `json:"total"`
	Hits  []*Organization `json:"hits"`
}

type OrganizationSearchService interface {
	Search(*SearchArgs) (*OrganizationHits, error)
	Index(*Organization) error
	Delete(string) error
	DeleteAll() error
	NewBulkIndexer(BulkIndexerConfig) (BulkIndexer[*Organization], error)
	NewIndexSwitcher(BulkIndexerConfig) (IndexSwitcher[*Organization], error)
	Suggest(string) ([]*Organization, error)
}
