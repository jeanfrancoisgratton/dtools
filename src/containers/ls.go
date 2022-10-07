// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/containers/ls.go
// 2022-10-07 12:48:18

package containers

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"time"
)

func Ls(all bool) {
	//clo := types.ContainerListOptions{Quiet: false, Size: true, All: true, Latest: true}

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	//containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{Latest: true})
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: all})
	if err != nil {
		panic(err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Container image", "Container name", "Created", "State", "Status"})
	for _, container := range containers {
		t.AppendRow([]interface{}{container.ID[:10], container.Image, container.Names, time.Unix(container.Created, 0).String(), container.State, container.Status})
	}
	t.SortBy([]table.SortBy{
		{Name: "Container name", Mode: table.Asc},
	})
	t.SetStyle(table.StyleBold)
	//t.Style().Options.DrawBorder = false
	//t.Style().Options.SeparateColumns = false
	t.Style().Format.Header = text.FormatDefault
	t.SetRowPainter(func(row table.Row) text.Colors {
		switch row[4] {
		case "running":
			return text.Colors{text.BgBlack, text.FgHiGreen}
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
