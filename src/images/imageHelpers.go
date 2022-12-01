// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/image/imageLs.go
// 2022-10-07 12:46:11

package images

import (
	"fmt"
	"strings"
)

func splitImageTag(imagetag string) (string, string) {
	dockerRegistry := "LOCAL"
	taggedImageName := ""
	tag := "latest"

	slashIndex := strings.Index(imagetag, "/")
	columnIndex := strings.LastIndex(imagetag, ":")

	// This means: no remote registry
	if slashIndex != -1 {
		dockerRegistry = imagetag
	}

}
func getImageTag(imageTagSlice []string) bool {

	for _, imagetag := range imageTagSlice {
		// First, we split image.RepoTags in two parts: reponame w/ port, and image name w/ tag
		dockerRegistry, imageName := splitImageTag(imagetag)
		imageRepository := "LOCAL"
		//tagLen := len(imageTag)

		if strings.Contains(imagetag, "/") {
			slashPos := strings.Index(imagetag, "/") + 1
			imageRepository = imagetag[:slashPos-1]
			imageName = imagetag[slashPos:len(imagetag)]
		} else {
			imageName = imagetag[:strings.Index(imagetag, ":")]
		}
		fmt.Println("Full name: ", imagetag)
		fmt.Println("Docker registry: ", imageRepository)
		fmt.Println("Image name: ", imageName)
		fmt.Println("Image tag: ")
		fmt.Println()

	}
	return true
}
