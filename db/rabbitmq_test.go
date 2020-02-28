package db

import (
	"superdispatcher/config"
	"superdispatcher/logger"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	config, err := config.New("config", "resources")
	assert.NoError(t, err)
	logger, err := logger.NewLog(config)
	rabbitMQ, err := NewRabbitMQ(config)
	assert.NoError(t, err)
	rabbitMQ.Logger = logger
	name := "Test"
	message := "Raymond Test"
	err = rabbitMQ.Publish(name, message)
	assert.NoError(t, err)
	msgs, err := rabbitMQ.Consume()
	assert.NoError(t, err)
	m := <-msgs
	s := string(m.Body)
	assert.Equal(t, message, s)
}
