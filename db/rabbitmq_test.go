package db

import (
	"superdispatcher/config"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	config, err := config.New("config", "../resources")
	assert.NoError(t, err)
	rabbitMQ, err := NewRabbitMQ(config)
	assert.NoError(t, err)

	name := "Test"
	message := "Raymond Test"
	err = rabbitMQ.Publish(name, message)
	assert.NoError(t, err)
	msgs, err := rabbitMQ.Consume()
	assert.NoError(t, err)
	for m := range msgs {
		assert.Equal(t, message, m.Body)
	}
}
