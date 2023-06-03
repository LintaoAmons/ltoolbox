/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/LintaoAmons/nric-go/nric"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// nricCmd represents the nric command
var nricCmd = &cobra.Command{
	Use:   "nric",
	Short: "Generate NRIC",
	Run: func(cmd *cobra.Command, args []string) {
		// get the count flag
		count, _ := cmd.Flags().GetString("count")
		// convert count to int
		countInt, err := strconv.Atoi(count)
		if err != nil {
			fmt.Println("count must be a number")
		}

		result := nric.GenerateNricMultiple(countInt)
		clipboard.WriteAll(result)
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(nricCmd)

	nricCmd.PersistentFlags().String("count", "1", "Number of NRIC to generate")
}
