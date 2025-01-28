package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wordcloud",
	Short: "A CLI tool to generate word clouds",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Here you can add subcommands or flags if needed
}