// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/containerRm.go
// 2022-11-28 13:51:08

package cmd

import (
	"dtools/containers"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove container",
	Long:  `This will remove a stopped container.`,
	Run: func(cmd *cobra.Command, args []string) {
		containers.RemoveContainer(args[0])
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
