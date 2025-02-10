package queue

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SendMessages(ch *amqp.Channel, data []byte, queueName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ch.PublishWithContext(ctx,
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(data),
		})
	failOnError(err, "Failed to publish a message")
}
