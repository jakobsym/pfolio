/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addPositionCmd represents the addPosition command
var addPositionCmd = &cobra.Command{
	Use:   "addPosition",
	Short: "Add a new position to your portfolio.",
	Long: `Adds a new position to your portfolio where
		(ticker, size, entry price, and order date) are recorded.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addPosition called")
	},
}

func init() {
	rootCmd.AddCommand(addPositionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addPositionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addPositionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
