package cmd

import (
	"fmt"

	"github.com/nicovillanueva/vitae/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version and exit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Server version: %s\nCLI version: %s", server.Version, version)
	},
}
