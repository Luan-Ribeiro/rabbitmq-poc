package main

import (
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"rabbitmq-poc/conf"
	"rabbitmq-poc/queue"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	conf.Config.QueueName = os.Getenv("QUEUE_NAME")
	conf.Config.RabbitmqUrl = os.Getenv("RABBITMQ_URL")
}

func main() {
	conn, err := amqp.Dial(conf.Config.RabbitmqUrl)
	if err != nil {
		log.Panicf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("Failed to connect to Channel: %s", err)
	}
	defer ch.Close()
	q, err := queue.NewQueue(ch)
	if err != nil {
		log.Panicf("Failed to declare a queue: %s", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
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
