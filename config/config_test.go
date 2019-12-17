package config

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestReadConfiguration(t *testing.T) {
	connect := viper.Get("connectionString")
	assert.NotEmpty(t, connect)
}
