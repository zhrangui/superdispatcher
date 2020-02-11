package db

import (
	"superdispatcher/config"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	config, err := config.New("config", "../resources")
	assert.Nil(t, err)
	rabbitMQ, err := NewRabbitMQ(config)
	assert.Nil(t, err)
	conn, err := rabbitMQ.Connect()
	defer conn.Close()
	rabbitMQ.Publish(conn, "Go Publish", "Raymond Test")
	assert.Nil(t, err)
}
