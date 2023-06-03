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

// postgresCreateTableToStructCmd represents the postgresCreateTableToStruct command
var postgresCreateTableToStructCmd = &cobra.Command{
	Use: "postgresCreateTableToStruct",
	Run: func(cmd *cobra.Command, args []string) {
		result, _ := functions.CreateTableToStruct(args[0])
		fmt.Println(result)
		clipboard.WriteAll(result)
	},
}

func init() {
	structConvertCmd.AddCommand(postgresCreateTableToStructCmd)
}
