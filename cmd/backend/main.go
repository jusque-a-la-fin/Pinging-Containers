package main

import (
	"log"
	"monitoring/internal/backend"
	"monitoring/internal/datastore"
	bkd "monitoring/internal/handlers/backend"
	"monitoring/internal/queue"
	"monitoring/internal/shared/config"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

// Backend-сервис
func main() {
	var wg sync.WaitGroup
	configName1 := "backend"
	configName2 := "database"
	configName3 := "rabbitmq"
	if err := config.SetupBackendConfig(configName1, configName2, configName3); err != nil {
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
	rtr.HandleFunc("/list", backendHanlder.GetList).Methods("GET")
	rtr.HandleFunc("/container/{id}", backendHanlder.GetContainer).Methods("GET")

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
	conn, ch, queueName := queue.CreateQueue()
	defer conn.Close()
	defer ch.Close()

	queue.ReceiveMessages(backendHanlder, ch, queueName, &wg)
	wg.Wait()
}
