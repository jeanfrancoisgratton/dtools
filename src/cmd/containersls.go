/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	//"fmt"

	"dtools/containers"
	"github.com/spf13/cobra"
)

// containerlsCmd represents the containerls command
var containerlsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"cls", "containerls"},
	Short:   "List all containers",
	Long:    `Equivalent to docker ps [-a].`,
	Run: func(cmd *cobra.Command, args []string) {
		allContainers := false
		if len(args) > 0 && args[0] == "all" {
			allContainers = true
		}
		containers.Ls(allContainers)
	},
}

func init() {
	rootCmd.AddCommand(containerlsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// containerlsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// containerlsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
