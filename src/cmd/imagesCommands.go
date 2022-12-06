/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dtools/images"
	"github.com/spf13/cobra"
)

var imglsCmd = &cobra.Command{
	Use:     "ils",
	Aliases: []string{"lsi", "imagels", "imgls"},
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

var imgpullCmd = &cobra.Command{
	Use:     "pull",
	Aliases: []string{"fetch", "get"},
	Short:   "Pulls an image from a registry",
	Long:    `Works exactly like docker pull.`,
	Run: func(cmd *cobra.Command, args []string) {
		images.Pull(args)
	},
}

func init() {
	rootCmd.AddCommand(imglsCmd)
	rootCmd.AddCommand(imgpullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imglsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imglsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
