package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goc",
	Short: "goc is a tool to manage your git repositories",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
