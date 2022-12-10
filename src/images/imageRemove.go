// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/images/imageRemove.go
// 2022-12-08 08:33:49


package images

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ImageRemove(args []string) {
	removalOptions := types.ImageRemoveOptions{ForceRemoval, ForceRemoval}
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("Unable to create docker client: %s", err)
	}

	cli.ImageRemove(ctx, "", removalOptions)
}
