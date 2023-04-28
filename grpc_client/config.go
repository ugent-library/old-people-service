package grpc_client

type ServerConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Insecure bool
	CaCert   string
}

type ClientConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Insecure bool
	CaCert   string
}
