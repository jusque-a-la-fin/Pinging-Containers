package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"monitoring/internal/queue"
	"monitoring/internal/shared/config"
	"monitoring/internal/utils"
	"net/http"
	"os/exec"
	"sync"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

type Container struct {
	IPv4            string `json:"IPv4"`
	PingDuration    string `json:"PingDuration"`
	SuccessPingTime string `json:"SuccessPingTime"`
	IsSuccess       bool   `json:"IsSuccess"`
}

func main() {
	excludedContainers := []string{"pinger", "rabbitmq", "database", "backend", "frontend", "nginx"}
	cns, err := getAllContainersList(excludedContainers)
	if err != nil {
	}

	configName1 := "pinger"
	configName2 := "rabbitmq"
	if err := config.SetupPingerConfig(configName1, configName2); err != nil {
		log.Fatalf("failed to load the config file: %s", err.Error())
	}

	conn, ch, queueName := queue.CreateQueue()
	defer conn.Close()
	defer ch.Close()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var wg sync.WaitGroup

		for _, ctr := range cns {
			wg.Add(1)
			go pingIP(ctr.IPv4, &wg, ch, queueName)
		}
		wg.Wait()
	}

	port := viper.GetString("port")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func getAllContainersList(names []string) ([]Container, error) {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed containers: %w", err)
	}

	var cns []Container
	for _, container := range containers {
		containerName := container.Names[0][1:]

		if utils.Сontains(names, containerName) {
			continue
		}

		ip, err := utils.GetContainerIP(cli, container.ID)
		if err != nil {
			log.Fatalf("Error getting container IP: %v", err)
		}

		ctr := Container{}
		ctr.IPv4 = ip
		cns = append(cns, ctr)
	}
	return cns, nil
}

func pingIP(ip string, wg *sync.WaitGroup, ch *amqp.Channel, queueName string) {
	defer wg.Done()
	start := time.Now()
	pingTime := start.Format("02-01-2006 15:04")
	cmd := exec.Command("ping", "-c", "1", ip)
	err := cmd.Run()
	duration := time.Since(start)

	ctr := Container{}
	ctr.IPv4 = ip
	ctr.PingDuration = duration.String()
	if err == nil {
		ctr.IsSuccess = true
		ctr.SuccessPingTime = pingTime
	}

	jsonData, err := json.Marshal(ctr)
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}

	queue.SendMessages(ch, jsonData, queueName)
}
