package models

type PersonHits struct {
	Start int       `json:"start"`
	Limit int       `json:"limit"`
	Total int       `json:"total"`
	Hits  []*Person `json:"hits"`
}

type M map[string]any

type SearchArgs struct {
	Start int
	Limit int
	Query M
	Sort  []M
}

type PersonSearchService interface {
	Search(*SearchArgs) (*PersonHits, error)
	Index(*Person) error
	Delete(string) error
	DeleteAll() error
	NewBulkIndexer(BulkIndexerConfig) (BulkIndexer[*Person], error)
	NewIndexSwitcher(BulkIndexerConfig) (IndexSwitcher[*Person], error)
	Suggest(string) ([]*Person, error)
}
