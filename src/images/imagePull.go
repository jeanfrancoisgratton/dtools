// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/images/imagePull.go
// 2022-12-05 11:55:16

package images

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"os"
	"runtime"
)

func Pull(args []string) {
	ctx := context.Background()
	bAllImages := false
	f, err := os.Create("/tmp/dockerpull.txt")
	defer f.Close()

	//if args[0] == "-a" || args[0] == "--all" {
	//	bAllImages = true
	//}

	pullOptions := types.ImagePullOptions{bAllImages, "", nil, runtime.GOARCH}
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("Unable to create docker client: %s", err)
	}

	fmt.Printf("Pulling image %s... ", args[0])
	ioreader, err := cli.ImagePull(ctx, args[0], pullOptions)
	defer ioreader.Close()
	if err != nil {
		panic(err)
	}
	io.Copy(io.Discard, ioreader)
	fmt.Println("done !")
}
