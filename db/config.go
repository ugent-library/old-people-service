package db

import "github.com/ugent-library/people/ent"

type Config struct {
	Client *ent.Client
	AesKey string
}
