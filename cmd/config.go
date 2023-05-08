package cmd

type ConfigDb struct {
	Url string `json:"url,omitempty" env:"URL" envDefault:"postgres://people:people@localhost:5432/authority?sslmode=disable"`
}

type ConfigNats struct {
	Url string `json:"url,omitempty" env:"URL" envDefault:"nats://localhost:4222"`
}

type ConfigApi struct {
	Host          string `json:"host,omitempty" env:"HOST" envDefault:"localhost"`
	Port          int    `json:"port,omitempty" env:"PORT" envDefault:"3999"`
	Username      string `json:"username,omitempty" env:"USERNAME,notEmpty"`
	Password      string `json:"password,omitempty" env:"PASSWORD,notEmpty"`
	TlsEnabled    bool   `json:"tls_enabled" env:"TLS_ENABLED"`
	TlsServerCert string `json:"tls_server_cert,omitempty" env:"TLS_SERVER_CERT"`
	TlsServerKey  string `json:"tls_server_key,omitempty" env:"TLS_SERVER_KEY"`
}

type ConfigApiProxy struct {
	Host string `json:"host,omitempty" env:"HOST" envDefault:"localhost"`
	Port int    `json:"port,omitempty" env:"PORT" envDefault:"4001"`
}

type ConfigApiClient struct {
	Host     string `json:"host,omitempty" env:"HOST" envDefault:"localhost"`
	Port     int    `json:"port,omitempty" env:"PORT" envDefault:"3999"`
	Username string `json:"username,omitempty" env:"USERNAME,notEmpty"`
	Password string `json:"password,omitempty" env:"PASSWORD,notEmpty"`
	Insecure bool   `json:"insecure" env:"INSECURE"`
	CaCert   string `json:"cacert,omitempty" env:"CACERT"`
}

type Config struct {
	Production bool            `json:"production" env:"PRODUCTION"`
	Db         ConfigDb        `json:"db,omitempty" envPrefix:"DB_"`
	Nats       ConfigNats      `json:"nats,omitempty" envPrefix:"NATS_"`
	Api        ConfigApi       `json:"api,omitempty" envPrefix:"API_"`
	ApiProxy   ConfigApiProxy  `json:"api_proxy,omitempty" envPrefix:"API_PROXY_"`
	ApiClient  ConfigApiClient `json:"api_client,omitempty" envPrefix:"API_CLIENT_"`
}
