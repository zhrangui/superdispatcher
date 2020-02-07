package db

import (
	"superdispatcher/config"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenDial(t *testing.T) {
	config, err := config.New("config", "../resources")
	assert.Nil(t, err)
	rabbitMQ, err := NewRabbitMQ(config)
	assert.Nil(t, err)
	_, err = rabbitMQ.Dial()

	assert.Nil(t, err)
}
