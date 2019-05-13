package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "0.2"

var rootCmd = &cobra.Command{
	Use: "vitae",
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
