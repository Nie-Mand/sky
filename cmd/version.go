package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Sky",
	Long:  `All software has versions. This is Sky-Scraper's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sky [SkyScraper] v0.1")
	},
}
