/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// editPositionCmd represents the editPosition command
var editPositionCmd = &cobra.Command{
	Use:   "editPosition",
	Short: "Edit a position in your portfolio.",
	Long:  `Edit the position size in your portfolio.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("editPosition called")
	},
}

func init() {
	rootCmd.AddCommand(editPositionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editPositionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editPositionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
