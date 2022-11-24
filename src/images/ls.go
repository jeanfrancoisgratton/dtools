// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/images/ls.go

package images

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Ls(allImg bool) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		//lastTag := len(image.RepoTags)
		//getImageTag(image.RepoTags)
		fmt.Println(image.ID)
		fmt.Println(image)
	}
}
