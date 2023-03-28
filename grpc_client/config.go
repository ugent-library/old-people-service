package grpc_client

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Insecure bool
	CaCert   string
}
