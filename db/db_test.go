package db

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"superdispatcher/config"

	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

func TestInvalidConnectionString(t *testing.T) {

	config, _ := config.New("config.staging")
	connString := config.MSSQL.ConnectionString

	conn, _ := sql.Open("mssql", connString)
	defer conn.Close()

	rows, _ := conn.Query("SELECT COUNT(*) FROM AtlasUsers")
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}

	assert.True(t, count > 1)
}
