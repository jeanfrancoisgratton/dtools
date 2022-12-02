/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dtools/images"
	"github.com/spf13/cobra"
)

// imglsCmd represents the imgls command
var imglsCmd = &cobra.Command{
	Use:     "imgls",
	Aliases: []string{"lsi", "imagels", "ils"},
	Short:   "Image list",
	Long:    `Similar to docker images, this will give you an inventory of all images on the host.`,
	Run: func(cmd *cobra.Command, args []string) {
		allImages := false
		if len(args) > 0 && args[0] == "all" {
			allImages = true
		}
		images.ImageList(allImages)
	},
}

func init() {
	rootCmd.AddCommand(imglsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imglsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imglsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
