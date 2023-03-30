package cmd

type ConfigDb struct {
	Url string `json:"url,omitempty"`
}

type ConfigNats struct {
	Url string `json:"url,omitempty"`
}

type ConfigApi struct {
	Host          string `json:"host,omitempty"`
	Port          int    `json:"port,omitempty"`
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	TlsEnabled    bool   `json:"tls_enabled"`
	TlsServerCert string `json:"tls_server_cert,omitempty"`
	TlsServerKey  string `json:"tls_server_key,omitempty"`
}

type ConfigApiProxy struct {
	Host string `json:"host,omitempty"`
	Port int    `json:"port,omitempty"`
}

type ConfigApiClient struct {
	Insecure bool   `json:"insecure"`
	CaCert   string `json:"cacert,omitempty"`
}

type Config struct {
	Production bool            `json:"production"`
	Db         ConfigDb        `json:"db,omitempty" mapstructure:"db"`
	Nats       ConfigNats      `json:"nats,omitempty" mapstructure:"nats"`
	Api        ConfigApi       `json:"api,omitempty" mapstructure:"api"`
	ApiProxy   ConfigApiProxy  `json:"api_proxy,omitempty" mapstructure:"api_proxy"`
	ApiClient  ConfigApiClient `json:"api_client,omitempty" mapstructure:"api_client"`
}
