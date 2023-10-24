package cmd

import (
	"github.com/spf13/cobra"

	"rwietter/goc/repository"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list your local repositories",
	Run: func(cmd *cobra.Command, args []string) {
		repositories := repository.GetRepositories()

		for _, r := range repositories {
			println(r.Path)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
