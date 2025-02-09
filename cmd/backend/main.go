package main

import (
	"log"
	"monitoring/internal/backend"
	"monitoring/internal/datastore"
	bkd "monitoring/internal/handlers/backend"
	"monitoring/internal/shared/config"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func main() {
	var wg sync.WaitGroup
	configName1 := "backend"
	configName2 := "database"
	if err := config.SetupBackendConfig(configName1, configName2); err != nil {
		log.Fatalf("failed to load the config file: %s", err.Error())
	}

	dtb, err := datastore.CreateNewDB()
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}

	bnd := backend.NewDBRepo(dtb)
	backendHanlder := &bkd.BackendHandler{
		BackendRepo: bnd,
	}

	rtr := mux.NewRouter()
	rtr.HandleFunc("/logs", backendHanlder.GetLogs).Methods("GET")

	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET"},
	})

	handler := crs.Handler(rtr)
	port := viper.GetString("port")

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = http.ListenAndServe(":"+port, handler)
		if err != nil {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	wg.Add(1)
	InitReceiver(backendHanlder, &wg)
	wg.Wait()
}

func InitReceiver(hnd *bkd.BackendHandler, wg *sync.WaitGroup) {
	conn, err := amqp.Dial("amqp://user:password@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"monitoring",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	ReceiveMessages(hnd, ch, queue.Name, wg)
}
