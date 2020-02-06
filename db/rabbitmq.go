package db

import (
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

func (rabbitMQ *RabbitMQ) Dial() error {
	cfg := rabbitMQ.config.Constants.RabbitMQ
	conn, err := amqp.Dial("amqp://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + string(cfg.Port))
	defer conn.Close()
	if err != nil {
		rabbitMQ.config.Logger.Fatal(err.Error())
	}

	return err
}
