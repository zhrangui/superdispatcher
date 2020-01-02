package config

type Constants struct {
	My    My
	MSSQL MSSQL
}

type My struct {
	LoggerType int
}

type MSSQL struct {
	ConnectionString string
}
