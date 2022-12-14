package main

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version info",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		showVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
