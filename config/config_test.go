package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadConfiguration(t *testing.T) {
	config, err := New("config", "../resources")
	if err != nil {
		t.Errorf("Unable to initiate config: %+v", err)
	}
	assert.NotEmpty(t, config.Constants.MSSQL.ConnectionString)
}
