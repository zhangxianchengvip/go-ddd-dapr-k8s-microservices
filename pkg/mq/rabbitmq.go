package mq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQ(connStr string) (*amqp.Connection, error) {

	conn, err := amqp.Dial(connStr)

	if err != nil {
		log.Panicf("%s", err)
	}

	return conn, nil
}
