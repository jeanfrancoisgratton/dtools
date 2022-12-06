// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/containersCommands.go
// 2022-11-28 13:55:19

package cmd

import (
	"dtools/containers"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Aliases: []string{"down", "containerdown"},
	Short:   "Cleanly stops a running container",
	Long:    `This will attempt to gracefully shut a container down.`,
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

var stopallCmd = &cobra.Command{
	Use:   "stopall",
	Short: "Cleanly stops all running containers",
	Long:  `This will attempt to gracefully shut containers down.`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.Stopall()
	},
}

var killallCmd = &cobra.Command{
	Use:   "killall",
	Short: "Kills all running containers",
	Long:  `Will SIGTERM all running containers.`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.Killall()
	},
}

var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"up", "containerup"},
	Short:   "Starts a stopped container",
	Run: func(cmd *cobra.Command, args []string) {
		containers.StartContainer(args[0])
	},
}

var startCallmd = &cobra.Command{
	Use:   "startall",
	Short: "Starts all containers",
	Run: func(cmd *cobra.Command, args []string) {
		containers.Startall()
	},
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts a container",
	Run: func(cmd *cobra.Command, args []string) {
		containers.RestartContainer(args[0])
	},
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove container",
	Long:  `This will remove a stopped container.`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.RemoveContainer(args[0])
	},
}

var lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"lsc", "containerls"},
	Short:   "List all containers",
	Long:    `Equivalent to docker ps [-a].`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.ContainerList(containers.QuietOutput)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(killCmd)
	rootCmd.AddCommand(stopallCmd)
	rootCmd.AddCommand(killallCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(startCallmd)
	rootCmd.AddCommand(restartCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
