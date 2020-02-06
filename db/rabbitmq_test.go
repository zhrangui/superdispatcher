package db

import (
	"log"
	"superdispatcher/config"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenDial(t *testing.T) {
	config, err := config.New("config", "../resources")
	rabbitMQ, err := NewRabbitMQ(config)
	if err != nil {
		log.Fatal(err)
	}
	err = rabbitMQ.Dial()
	assert.Nil(t, err)
}
