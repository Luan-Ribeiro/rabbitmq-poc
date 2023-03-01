package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panicf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	queueName := "rabbitmq-test"
	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("Failed to connect to Channel: %s", err)
	}
	defer ch.Close()
	queue, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Panicf("Failed to declare a queue: %s", err)
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	var forever chan struct{}

	go func() {
		for m := range msgs {
			log.Printf("Received a message: %s", m.Body)
		}
	}()

	log.Printf(" Waiting for messages. To exit press CTRL+C")
	<-forever

}
