package queue

import (
	"github.com/rabbitmq/amqp091-go"
	"rabbitmq-poc/conf"
)

func NewQueue(ch *amqp091.Channel) (amqp091.Queue, error) {
	return ch.QueueDeclare(
		conf.Config.QueueName, // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
}
