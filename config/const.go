package config

type Constants struct {
	My    My
	MSSQL MSSQL
}

type My struct {
	LoggerType int
	Protocal   string
	IpAddress  string
}

type MSSQL struct {
	ConnectionString string
}
