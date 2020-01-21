package config

type Constants struct {
	My    My
	MSSQL MSSQL
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
