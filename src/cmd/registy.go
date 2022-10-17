/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// registyCmd represents the registy command
var registyCmd = &cobra.Command{
	Use:   "registy",
	Short: "Docker registry related subcommands",
	Long:  `This is where you will find all registry-related commands, such as list tags, list images, select registry, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("registy called")
	},
}

func init() {
	rootCmd.AddCommand(registyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
