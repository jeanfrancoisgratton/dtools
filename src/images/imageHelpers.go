// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/image/imageLs.go
// 2022-10-07 12:46:11

package images

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// REPOSITORY              TAG              IMAGE ID       CREATED      SIZE
// rpmbuilder              latest           d7f4c25238e4   3 days ago   414MB
// nexus:9820/rpmbuilder   10.00.00-arm64   d7f4c25238e4   3 days ago   414MB
// rocky                   test             662704dd4eee   3 days ago   301MB

type imageSpec struct {
	id, repo, name, tag, created, formattedSize string
	size                                        int64
}

var ForceRemoval bool

func splitImageTag(imagetag string) (string, string, string) {
	dockerRegistry := "--"
	imgtag := "latest"
	imgname := ""

	slashIndex := strings.Index(imagetag, "/")
	columnIndex := strings.LastIndex(imagetag, ":")

	// This means: no remote registry
	if slashIndex == -1 {
		//if columnIndex != -1 {
		imgname = imagetag[:columnIndex]
		imgtag = imagetag[columnIndex+1:]
		//}
	} else {
		dockerRegistry = imagetag[:slashIndex]
		imgname = imagetag[slashIndex+1 : columnIndex]
		imgtag = imagetag[columnIndex+1:]
	}
	return dockerRegistry, imgname, imgtag
}
func getImageTag(id string, imageTagSlice []string, created int64, size int64) []imageSpec {
	var imgspecSlice []imageSpec
	var imgspec imageSpec

	for _, imagetag := range imageTagSlice {
		// First, we split image.RepoTags in two parts: reponame w/ port, and image name w/ tag
		imgspec.repo, imgspec.name, imgspec.tag = splitImageTag(imagetag)
		imgspec.id = id[7:]
		// Then we add creation time & size
		imgspec.created = time.Unix(created, 0).Format("2006.01.02 15:04:05")
		imgspec.formattedSize = formatImageSize(size)
		imgspecSlice = append(imgspecSlice, imgspec)
	}
	return imgspecSlice
}

// formatImageSize() : just so the image size shows MB or GB when needed
func formatImageSize(sz int64) string {
	numSize := (float32)(sz) / 1000.0 / 1000.0 // this will give us the size in MB

	if (int)(math.Log10((float64(numSize)))) > 2 {
		return fmt.Sprintf("%.3f GB", numSize/1000.0)
	} else {
		return fmt.Sprintf("%.3f MB", numSize)
	}
}
