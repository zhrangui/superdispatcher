package db

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"

	"testing"

	"github.com/stretchr/testify/assert"

	"superdispatcher/config"
)

func TestInvalidConnectionString(t *testing.T) {

	config, _ := config.New("config", "resources")
	connString := config.Constants.MSSQL.ConnectionString

	conn, _ := sql.Open("mssql", connString)
	defer conn.Close()

	rows, _ := conn.Query("SELECT COUNT(*) FROM AtlasUsers")
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}

	assert.True(t, count > 1)
}
