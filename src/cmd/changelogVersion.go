// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/changelogVersion.go
// 2022-12-06 08:11:32

package cmd

import (
	"dtools/misc"
	"github.com/spf13/cobra"
)

var clVerCmd = &cobra.Command{
	Use:     "cl",
	Aliases: []string{"changelog"},
	Short:   "Shows the app changelog",
	Long:    `Show the application changelog.`,
	Run: func(cmd *cobra.Command, args []string) {
		misc.Changelog()
	},
}

func init() {
	rootCmd.AddCommand(clVerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// changelogVersionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// changelogVersionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
