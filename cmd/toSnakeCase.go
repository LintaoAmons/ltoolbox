/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/LintaoAmons/toolbox/functions"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// toSnakeCaseCmd represents the toSnakeCase command
var toSnakeCaseCmd = &cobra.Command{
	Use:   "toSnakeCase",
	Short: "convert to aaa_bbb_ccc format",
	Run: func(cmd *cobra.Command, args []string) {
		result := functions.ToSnakeCase(args[0])
		clipboard.WriteAll(result)
		fmt.Println(result)
	},
}

func init() {
	formatConvertCmd.AddCommand(toSnakeCaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toSnakeCaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toSnakeCaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
