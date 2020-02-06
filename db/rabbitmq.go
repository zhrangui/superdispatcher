package db

import (
	"fmt"
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
	connString := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
	conn, err := amqp.Dial(connString)
	defer conn.Close()
	rabbitMQ.config.Logger.FailOnError(err, fmt.Sprintf("failed to establish RabbitMQ connection: %s", connString))
	return conn, err
}
