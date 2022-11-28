// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/containers/imageLs.go
// 2022-10-07 12:48:18

package containers

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"time"
)

func ContainerList(all bool) {
	//clo := types.ContainerListOptions{Quiet: false, Size: true, All: true, Latest: true}

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	//containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{Latest: true})
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: all})
	if err != nil {
		errmsg := fmt.Sprintf("%v", err)
		if errmsg == "Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?" {
			fmt.Println(errmsg)
			os.Exit(-1)
		} else {
			panic(err)
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Container image", "Container name", "Container Created", "Exposed ports", "State", "Status"})
	for _, container := range containers {
		// This is a design decision: I'll take only the first name in the container slice
		cn := container.Names[0]
		ports := prettifyPortsList(container.Ports)
		t.AppendRow([]interface{}{container.ID[:10], container.Image, cn[1:], time.Unix(container.Created, 0).Format("2006.01.02 15:04:05"), ports, container.State, container.Status})
	}
	t.SortBy([]table.SortBy{
		{Name: "Container name", Mode: table.Asc},
	})
	t.SetStyle(table.StyleBold)
	//t.Style().Options.DrawBorder = false
	//t.Style().Options.SeparateColumns = false
	t.Style().Format.Header = text.FormatDefault
	t.SetRowPainter(func(row table.Row) text.Colors {
		switch row[5] {
		case "running":
			//return text.Colors{text.BgBlack, text.FgHiGreen}
			return text.Colors{text.FgHiGreen}
			//case "Crashed":
			//	return text.Colors{text.BgBlack, text.FgHiRed}
			//case "Blocked":
			//case "Suspended":
			//case "Paused":
			//	return text.Colors{text.BgHiBlack, text.FgHiYellow}
		}
		return nil
	})
	t.Render()
}
