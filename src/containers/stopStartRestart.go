// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/containers/stopStartRestart.go
// 2022-11-28 13:47:22

package containers

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

func StopContainer(containername string) error {
	ctx := context.Background()
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Printf("Unable to create docker client: %s", err)
	}

	if err := client.ContainerStop(ctx, containername, nil); err != nil {
		log.Printf("Unable to stop container %s: %s", containername, err)
	}

	fmt.Printf("Container %s is stopped.\n", containername)
	return nil
}

func KillContainer(containername string) error {
	ctx := context.Background()
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Printf("Unable to create docker client: %s", err)
	}

	if err := client.ContainerKill(ctx, containername, "TERM"); err != nil {
		log.Printf("Unable to kill container %s: %s", containername, err)
	}

	fmt.Printf("Container %s is killed.\n", containername)
	return nil
}

func StartContainer(containername string) error {
	ctx := context.Background()
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Printf("Unable to create docker client: %s", err)
	}

	if err := client.ContainerStart(ctx, containername, types.ContainerStartOptions{}); err != nil {
		log.Printf("Unable to start container %s: %s", containername, err)
	}

	fmt.Printf("Container %s is started.\n", containername)
	return nil
}

func RestartContainer(containername string) error {
	StopContainer(containername)
	StartContainer(containername)
	return nil
}

func Killall() error {
	containers := ContainerList(true)

	for _, container := range containers {
		KillContainer(container.Names[0])
	}

	return nil
}

func Stopall() error {
	containers := ContainerList(true)

	for _, container := range containers {
		StopContainer(container.Names[0])
	}

	return nil
}

func Startall() error {
	containers := ContainerList(true)

	for _, container := range containers {
		StartContainer(container.Names[0])
	}

	return nil
}
