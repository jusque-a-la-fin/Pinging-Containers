package queue

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func CreateQueue() (*amqp.Connection, *amqp.Channel, string) {
	username := viper.GetString("rabbitmq.username")
	password := viper.GetString("rabbitmq.password")
	port := viper.GetString("rabbitmq.port")
	url := fmt.Sprintf("amqp://%s:%s@rabbitmq:%s/", username, password, port)
	conn, err := amqp.Dial(url)

	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	queuename := viper.GetString("rabbitmq.queuename")
	queue, err := ch.QueueDeclare(
		queuename,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	return conn, ch, queue.Name
}
