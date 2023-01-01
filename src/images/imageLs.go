// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/images/imageLs.go

package images

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
)

func ImageList(allImg bool) {
	var imgspecSlice []imageSpec

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
		//imgSpec := getImageTag(image.ID, image.RepoTags, image.Created, image.Size)
		imgspecSlice = append(imgspecSlice, getImageTag(image.ID, image.RepoTags, image.Created, image.Size)...)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Registry", "Image name", "Image tag", "Image ID", "Creation time", "Approx size"})
	for _, imgspec := range imgspecSlice {
		// This is a design decision: I'll take only the first name in the container slice
		t.AppendRow([]interface{}{imgspec.repo, imgspec.name, imgspec.tag, imgspec.id[:12], imgspec.created, imgspec.formattedSize})
	}
	t.SortBy([]table.SortBy{
		{Name: "Image name", Mode: table.Asc},
	})
	t.SetStyle(table.StyleBold)
	t.Style().Format.Header = text.FormatDefault
	t.Render()
}
