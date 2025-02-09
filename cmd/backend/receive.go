package main

import (
	"log"
	bkd "monitoring/internal/handlers/backend"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ReceiveMessages(hnd *bkd.BackendHandler, ch *amqp.Channel, queueName string, wg *sync.WaitGroup) {
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan struct{})

	go func() {
		defer wg.Done()
		for d := range msgs {
			hnd.UpdateContainers(d.Body)
		}
	}()

	<-forever
}
