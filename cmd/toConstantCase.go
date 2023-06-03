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

// toConstantCaseCmd represents the toConstantCase command
var toConstantCaseCmd = &cobra.Command{
	Use:   "toConstantCase",
	Short: "convert to AAA_BBB_CCC format",
	Run: func(cmd *cobra.Command, args []string) {
		result := functions.ToConstantCase(args[0])
		clipboard.WriteAll(result)
		fmt.Println(result)
	},
}

func init() {
	formatConvertCmd.AddCommand(toConstantCaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toConstantCaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toConstantCaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
