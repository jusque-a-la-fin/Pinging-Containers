package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"monitoring/internal/shared/config"
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

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	excludedContainers := []string{"pinger"}
	cns, err := getAllContainersList(excludedContainers)
	if err != nil {
	}

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

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var wg sync.WaitGroup

		for _, ctr := range cns {
			wg.Add(1)
			go pingIP(ctr.IPv4, &wg, ch, queue.Name)
		}
		wg.Wait()
	}

	configName := "pinger"
	if err := config.SetupPingerConfig(configName); err != nil {
		log.Fatalf("failed to load the config file: %s", err.Error())
	}

	port := viper.GetString("port")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
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

		if contains(names, containerName) {
			continue
		}

		ip, err := getContainerIP(cli, container.ID)
		if err != nil {
			log.Fatalf("Error getting container IP: %v", err)
		}

		ctr := Container{}
		ctr.IPv4 = ip
		cns = append(cns, ctr)
	}
	return cns, nil
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func getContainerIP(cli *client.Client, containerID string) (string, error) {
	inspect, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return "", fmt.Errorf("fetching IP fails for container %s: %w", containerID, err)
	}

	for _, network := range inspect.NetworkSettings.Networks {
		return network.IPAddress, nil
	}
	return "", fmt.Errorf("no IP address found for container %s", containerID)
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
	} else {
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonData, err := json.Marshal(ctr)
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}

	err = ch.PublishWithContext(ctx,
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(jsonData),
		})
	failOnError(err, "Failed to publish a message")
}
