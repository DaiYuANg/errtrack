package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "CLI tool to manage your application",
	Run: func(cmd *cobra.Command, args []string) {
		container().Run()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
