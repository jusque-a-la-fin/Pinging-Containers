package queue

import (
	bkd "monitoring/internal/handlers/backend"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

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
