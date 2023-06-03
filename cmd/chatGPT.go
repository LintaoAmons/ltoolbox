/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/LintaoAmons/toolbox/functions/ai"
	"github.com/spf13/cobra"
)

// chatGPTCmd represents the chatGPT command
var chatGPTCmd = &cobra.Command{
	Use:   "chatGPT",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ai.AskChatGPT(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(chatGPTCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatGPTCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatGPTCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
