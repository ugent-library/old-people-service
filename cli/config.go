package cli

import "fmt"

type ConfigDb struct {
	Url    string `json:"url,omitempty" env:"URL" envDefault:"postgres://people:people@localhost:5432/authority?sslmode=disable"`
	AesKey string `json:"-" env:"AES_KEY,notEmpty"`
}

type ConfigNats struct {
	Url      string `json:"url,omitempty" env:"URL" envDefault:"nats://localhost:4222"`
	Nkey     string `json:"-" env:"NKEY"`
	NkeySeed string `json:"-" env:"NKEY_SEED"`
}

type ConfigApi struct {
	Host string `json:"host,omitempty" env:"HOST" envDefault:"localhost"`
	Port int    `json:"port,omitempty" env:"PORT" envDefault:"3999"`
	Key  string `json:"key,omitempty" env:"KEY,notEmpty"`
}

type ConfigLdap struct {
	Url      string `json:"url,omitempty" env:"URL,notEmpty"`
	Username string `json:"username,omitempty" env:"USERNAME,notEmpty"`
	Password string `json:"password,omitempty" env:"PASSWORD,notEmpty"`
}

type Config struct {
	Production bool       `json:"production" env:"PRODUCTION"`
	Db         ConfigDb   `json:"db,omitempty" envPrefix:"DB_"`
	Nats       ConfigNats `json:"nats,omitempty" envPrefix:"NATS_"`
	Api        ConfigApi  `json:"api,omitempty" envPrefix:"API_"`
	Ldap       ConfigLdap `json:"ldap,omitempty" envPrefix:"LDAP_"`
}

func (ca ConfigApi) Addr() string {
	return fmt.Sprintf("%s:%d", ca.Host, ca.Port)
}
