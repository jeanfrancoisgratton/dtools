// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/images/imageLs.go

package images

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"
)

func ImageList(allImg bool) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("Unable to create docker client: %s", err)
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		errmsg := fmt.Sprintf("%v", err)
		if errmsg == "Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?" {
			fmt.Println(errmsg)
			os.Exit(-1)
		} else {
			panic(err)
		}
	}

	for _, image := range images {
		getImageTag(image.RepoTags)
	}
}
