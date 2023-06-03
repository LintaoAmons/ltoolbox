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

// toCamelCaseCmd represents the toCamelCase command
var toCamelCaseCmd = &cobra.Command{
	Use:   "toCamelCase",
	Short: "convert to aaaBbbCcc format",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		result := functions.ToCamelCase(args[0])
		clipboard.WriteAll(result)
		fmt.Println(result)
	},
}

func init() {
	formatConvertCmd.AddCommand(toCamelCaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toCamelCaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toCamelCaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
