package backend

import (
	"context"
	"fmt"
	"log"
	"monitoring/internal/utils"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func (repo *BackendDBRepository) GetList() ([]string, error) {
	excludedContainers := []string{"pinger", "rabbitmq", "database", "backend", "frontend", "nginx"}
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed containers: %w", err)
	}

	var ips []string
	for _, container := range containers {
		containerName := container.Names[0][1:]

		if utils.Сontains(excludedContainers, containerName) {
			continue
		}

		ip, err := utils.GetContainerIP(cli, container.ID)
		if err != nil {
			log.Fatalf("Error getting container IP: %v", err)
		}
		ips = append(ips, ip)
	}
	return ips, nil
}
