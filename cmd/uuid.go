/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate UUID",
	Run: func(cmd *cobra.Command, args []string) {
		// get the count flag
		count, _ := cmd.Flags().GetString("count")
		// convert count to int
		countInt, err := strconv.Atoi(count)
		if err != nil {
			fmt.Println("count must be a number")
		}

		uuids := []string{}
		for i := 0; i < countInt; i++ {
			uuids = append(uuids, uuid.New().String())
		}
		uuidsString := strings.Join(uuids, "\n")
		clipboard.WriteAll(uuidsString)
		fmt.Println(uuidsString)
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	uuidCmd.PersistentFlags().String("count", "1", "Number of UUID to generate")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uuidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
