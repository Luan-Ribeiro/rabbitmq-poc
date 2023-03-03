package main

import (
	"context"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"rabbitmq-poc/conf"
	"rabbitmq-poc/queue"
	"time"
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
		log.Panicf("Failed to inicialize queue: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "first message"
	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Panicf("Failed to publish on queue: %s", err)
	}
	log.Printf(" Sent %s\n", body)
}
