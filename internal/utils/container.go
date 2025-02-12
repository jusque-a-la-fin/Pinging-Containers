package utils

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func Сontains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func GetContainerIP(cli *client.Client, containerID string) (string, error) {
	inspect, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return "", fmt.Errorf("fetching IP fails for container %s: %w", containerID, err)
	}

	for _, network := range inspect.NetworkSettings.Networks {
		return network.IPAddress, nil
	}
	return "", fmt.Errorf("no IP address found for container %s", containerID)
}
