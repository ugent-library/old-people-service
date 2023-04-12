package models

type M map[string]any

type SearchArgs struct {
	Start int
	Limit int
	Query M
	Sort  []M
}
