// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/constainerStopKill.go
// 2022-11-28 13:55:19

package cmd

import (
	"dtools/containers"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Cleanly stops a running container",
	Long:  `This will attempt to gracefully shut a container down.`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.StopContainer(args[0])
	},
}

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kills a running container",
	Long:  `Will SIGTERM a running container.`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.KillContainer(args[0])
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(killCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
