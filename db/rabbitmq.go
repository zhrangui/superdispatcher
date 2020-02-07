package db

import (
	"fmt"
	"net/url"
	"superdispatcher/config"

	"github.com/streadway/amqp"
)

// RabbitMQ interface
type RabbitMQ struct {
	config *config.Config
}

// NewRabbitMQ creates RabbitMQ client instance
func NewRabbitMQ(cfg *config.Config) (*RabbitMQ, error) {
	var rabbit = &RabbitMQ{}
	rabbit.config = cfg
	return rabbit, nil
}

func (rabbitMQ *RabbitMQ) Dial() (*amqp.Connection, error) {
	cfg := rabbitMQ.config.Constants.RabbitMQ
	connString := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.User, url.PathEscape(cfg.Password), cfg.Host, cfg.Port)
	conn, err := amqp.Dial(connString)
	if err != nil {
		rabbitMQ.config.Logger.Error(err, fmt.Sprintf("failed to establish RabbitMQ connection: %+v", connString))
		return nil, err
	}
	defer conn.Close()

	return conn, err
}
