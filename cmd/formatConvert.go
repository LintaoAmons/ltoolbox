/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// formatConvertCmd represents the formatConvert command
var formatConvertCmd = &cobra.Command{
	Use:   "formatConvert",
	Short: "convert between different format",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please call formatConvert with sub-commands")
	},
}

func init() {
	rootCmd.AddCommand(formatConvertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// formatConvertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// formatConvertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
