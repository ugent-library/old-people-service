package cli

import "fmt"

type ConfigDb struct {
	Url    string `env:"URL" envDefault:"postgres://people:people@localhost:5432/authority?sslmode=disable"`
	AesKey string `env:"AES_KEY,notEmpty"`
}

type ConfigNats struct {
	Url        string `env:"URL" envDefault:"nats://localhost:4222"`
	Nkey       string `env:"NKEY"`
	NkeySeed   string `env:"NKEY_SEED"`
	StreamName string `env:"STREAM_NAME,notEmpty"`
}

type ConfigApi struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port int    `env:"PORT" envDefault:"3999"`
	Key  string `env:"KEY,notEmpty"`
}

type ConfigLdap struct {
	Url      string `env:"URL,notEmpty"`
	Username string `env:"USERNAME,notEmpty"`
	Password string `env:"PASSWORD,notEmpty"`
}

type Config struct {
	Version struct {
		Branch string `env:"SOURCE_BRANCH"`
		Commit string `env:"SOURCE_COMMIT"`
		Image  string `env:"IMAGE_NAME"`
	}
	Production bool       `env:"PEOPLE_PRODUCTION"`
	Db         ConfigDb   `envPrefix:"PEOPLE_DB_"`
	Nats       ConfigNats `envPrefix:"PEOPLE_NATS_"`
	Api        ConfigApi  `envPrefix:"PEOPLE_API_"`
	Ldap       ConfigLdap `envPrefix:"PEOPLE_LDAP_"`
	IPRanges   string     `env:"PEOPLE_IP_RANGES"`
}

func (ca ConfigApi) Addr() string {
	return fmt.Sprintf("%s:%d", ca.Host, ca.Port)
}
