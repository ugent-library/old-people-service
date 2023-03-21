package cmd

type ConfigDb struct {
	Url string `json:"url,omitempty"`
}

type ConfigNats struct {
	Url string `json:"url,omitempty"`
}

type Config struct {
	Production bool        `json:"production"`
	Db         *ConfigDb   `json:"db,omitempty"`
	Host       string      `json:"host,omitempty"`
	Port       int         `json:"port,omitempty"`
	Nats       *ConfigNats `json:"nats,omitempty"`
}
