/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/blackflame007/codeview/functions"
	"github.com/spf13/cobra"
)

// highlightCmd represents the highlight command
var highlightCmd = &cobra.Command{
	Use:   "highlight",
	Short: "A brief description of your command",
	Long:  `A longer description...`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		functions.Highlight()
	},
}

func init() {
	rootCmd.AddCommand(highlightCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// highlightCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// highlightCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
