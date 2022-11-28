// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/containers/Remove.go
// 2022-11-28 13:48:44

package containers

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

func RemoveContainer(containername string) error {
	ctx := context.Background()
	removeOptions := types.ContainerRemoveOptions{RemoveVolumes: true, Force: true}
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Printf("Unable to create docker client: %s", err)
	}

	if err := client.ContainerRemove(ctx, containername, removeOptions); err != nil {
		log.Printf("Unable to remove container: %s", err)
		return err
	}

	fmt.Printf("Container %s removed.\n", containername)
	return nil
}
