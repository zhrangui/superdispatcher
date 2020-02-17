package config

type Constants struct {
	My       My
	MSSQL    MSSQL
	RabbitMQ RabbitMQ
	Log      string
}

type My struct {
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
	Qname    string
}
