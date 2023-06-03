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

// ToKebabCaseCmd represents the ToKebabCase command
var ToKebabCaseCmd = &cobra.Command{
	Use:   "ToKebabCase",
	Short: "convert to aaa-bbb-ccc format",
	Run: func(cmd *cobra.Command, args []string) {
		result := functions.ToKebabCase(args[0])
		clipboard.WriteAll(result)
		fmt.Println(result)
	},
}

func init() {
	formatConvertCmd.AddCommand(ToKebabCaseCmd)
}
