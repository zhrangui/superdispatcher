package config

type Constants struct {
	MSSQL MSSQL
}

type MSSQL struct {
	ConnectionString string
}
