package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfiguration(t *testing.T) {
	config, err := New("config.staging")
	if err != nil {
		t.Errorf("Unable to initiate config: %+v", err)
	}
	assert.NotEmpty(t, config.MSSQL.ConnectionString)
}
