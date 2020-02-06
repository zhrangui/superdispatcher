package db

import (
	"superdispatcher/config"

	"github.com/pkg/errors"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenDial(t *testing.T) {
	config, err := config.New("config", "../resources")
	rabbitMQ, e := NewRabbitMQ(config)
	if e != nil {
		err = errors.Wrap(err, e.Error())
	}

	_, e = rabbitMQ.Dial()
	if e != nil {
		err = errors.Wrap(err, e.Error())
	}
	assert.Nil(t, err)
}
