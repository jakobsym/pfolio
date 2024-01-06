/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removePositionCmd represents the removePosition command
var removePositionCmd = &cobra.Command{
	Use:   "removePosition",
	Short: "Remove a position from your portfolio.",
	Long:  `Completely removes a position from your portfolio (Use when exiting a position).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("removePosition called")
	},
}

func init() {
	rootCmd.AddCommand(removePositionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removePositionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removePositionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
