package models

import "context"

type BulkIndexerConfig struct {
	OnError      func(error)
	OnIndexError func(string, error)
}

type IndexSwitcher[T any] interface {
	Index(context.Context, T) error
	Switch(context.Context) error
}

type BulkIndexer[T any] interface {
	Index(context.Context, T) error
	Close(context.Context) error
}
