package db

import (
	"fmt"
	"net/url"
	"superdispatcher/config"

	"github.com/streadway/amqp"
)

// RabbitMQ interface
type RabbitMQ struct {
	config     *config.Config
	connection *amqp.Connection
}

func (rabbitMQ *RabbitMQ) declareQueue(channel *amqp.Channel) (amqp.Queue, error) {
	q, err := channel.QueueDeclare(
		rabbitMQ.config.Constants.RabbitMQ.Qname,
		false,
		true,
		false,
		false,
		nil,
	)
	rabbitMQ.logError(err, "Failed to declare a queue")
	return q, err
}

// NewRabbitMQ creates RabbitMQ client instance
func NewRabbitMQ(cfg *config.Config) (*RabbitMQ, error) {
	var rabbit = &RabbitMQ{}
	rabbit.config = cfg
	return rabbit, nil
}

func (rabbitMQ *RabbitMQ) logError(err error, message string) {
	rabbitMQ.config.Logger.Error(err, message)
}

func (rabbitMQ *RabbitMQ) connect() (*amqp.Connection, error) {
	if rabbitMQ.connection != nil {
		return rabbitMQ.connection, nil
	}
	cfg := rabbitMQ.config.Constants.RabbitMQ
	connString := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", cfg.User, url.PathEscape(cfg.Password), cfg.Host, cfg.Port, cfg.Vhost)
	conn, err := amqp.Dial(connString)
	rabbitMQ.logError(err, fmt.Sprintf("failed to establish RabbitMQ connection: %+v", connString))
	if err == nil {
		rabbitMQ.connection = conn
	}
	return conn, err
}

func (rabbitMQ *RabbitMQ) Publish(name string, message string) error {
	conn, err := rabbitMQ.connect()
	if err != nil {
		return err
	}
	ch, err := conn.Channel()
	rabbitMQ.logError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := rabbitMQ.declareQueue(ch)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	rabbitMQ.logError(err, fmt.Sprintf("Failed to publish: %+v", message))
	return err
}

func (rabbitMQ *RabbitMQ) Consume() (<-chan amqp.Delivery, error) {
	conn, err := rabbitMQ.connect()
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	rabbitMQ.logError(err, "Failed to open a channel")
	q, err := rabbitMQ.declareQueue(ch)
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	rabbitMQ.logError(err, "Failed to connect consume channel")
	return msgs, err
}
