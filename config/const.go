package config

type Constants struct {
	My       My
	MSSQL    MSSQL
	RabbitMQ RabbitMQ
}

type My struct {
	LoggerType  int
	Protocal    string
	ServerIP    string
	HostAddress string
}

type MSSQL struct {
	ConnectionString string
}

type RabbitMQ struct {
	Scheme   string
	Host     string
	Port     int
	User     string
	Password string
	Vhost    string
}
