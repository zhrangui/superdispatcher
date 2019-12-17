package db

import (
	"testing"

	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"

)

func TestInvalidConnectionString(t *testing.T) {

	connString := viper.GetString("connectionString")

	conn, _ := sql.Open("mssql", connString)
	defer conn.Close()

	rows, _ := conn.Query("SELECT COUNT(*) FROM AtlasUsers")
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}

	if count <= 0 {
		t.Fail()
	}
}
