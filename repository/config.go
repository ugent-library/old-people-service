package repository

import "github.com/ugent-library/person-service/ent"

type Config struct {
	Client *ent.Client
	AesKey string
}
