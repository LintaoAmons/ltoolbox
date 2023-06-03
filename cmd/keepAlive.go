/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

// keepAliveCmd represents the keepAlive command
var keepAliveCmd = &cobra.Command{
	Use:   "keepAlive",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Seed the random number generator
		rand.Seed(time.Now().UnixNano())

		for {
			// Get the screen size
			screenWidth, screenHeight := robotgo.GetScreenSize()

			// Generate a random position for the mouse
			x := rand.Intn(screenWidth)
			y := rand.Intn(screenHeight)

			// Move the mouse pointer to the random position
			robotgo.MoveMouseSmooth(x, y)

			// Wait for 20 seconds before moving the mouse again
			time.Sleep(20 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(keepAliveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keepAliveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keepAliveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
