// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/container/ls.go
// 2022-10-07 12:48:18

package container

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Ls(all bool) {
	//clo := types.ContainerListOptions{Quiet: false, Size: true, All: true, Latest: true}

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	//containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{Latest: true})
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: all})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s %s %d %d %s %s\n",
			container.ID[:10], container.Image, container.Names, container.SizeRw/1024, container.SizeRootFs/1024, container.State, container.Status)
	}
}
