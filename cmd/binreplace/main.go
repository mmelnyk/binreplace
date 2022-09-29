package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   filepath.Base(os.Args[0]),
	Short: filepath.Base(os.Args[0]) + " is a tool to patch binaries in different ways",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var dryrun bool //	Dry Run operation

func main() {
	rootCmd.PersistentFlags().BoolVarP(&dryrun, "dryrun", "d", false, "dry run operation")
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
