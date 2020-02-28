package db

import (
	_ "github.com/denisenkom/go-mssqldb"

	"testing"

	"github.com/stretchr/testify/assert"

	"superdispatcher/config"
)

func TestOpenDB(t *testing.T) {
	config, err := config.New("config", "resources")
	assert.NoError(t, err)
	sql, err := NewSQL(config)
	assert.NoError(t, err)
	err = sql.Open()
	defer sql.Close()
}

func TestTableQuery(t *testing.T) {
	config, err := config.New("config", "resources")
	assert.NoError(t, err)
	sql, err := NewSQL(config)
	assert.NoError(t, err)
	err = sql.Open()
	assert.NoError(t, err)
	defer sql.Close()
	var users = AtlasUsers{}
	sql.Unscoped().Debug().First(&users)
	assert.NotEmpty(t, users.ComputerName)
}
