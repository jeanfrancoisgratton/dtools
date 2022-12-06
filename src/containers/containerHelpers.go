// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/containers/containerHelpers.go
// 2022-10-08 13:56:41

package containers

import (
	"fmt"
	"github.com/docker/docker/api/types"
)

var QuietOutput = false

func prettifyPortsList(ports []types.Port) string {
	var portsString, sourcePort string
	for _, val := range ports {
		if val.IP == "" {
			sourcePort = ""
		} else {
			sourcePort = fmt.Sprintf("%d->", val.PublicPort)
		}
		portsString += fmt.Sprintf("%s/%s%d  ", val.Type, sourcePort, val.PrivatePort)
	}
	return portsString
}
