package cmd

import (
	"fmt"

	"github.com/nicovillanueva/vitae/server"

	"github.com/spf13/cobra"
)

var flagPath string

const (
	defaultConfigPath = "./cv-data.yaml"
	configEnvVariable = "CV_DATA_PATH"
	pdfEnvVariable    = "CV_PDF_PATH"
)

var startCmd = &cobra.Command{
	Use: "start",
	Long: "Start the API server. Configuration file is read from: \n1. The path provided via`--config`\n" +
		"2. The path provided via the env variable " + configEnvVariable + "\n" +
		"3. The default path, " + defaultConfigPath + "\n" +
		"The path to the served PDF is overridable via the env variable CV_PDF_PATH, " +
		"otherwise it's configured in the YAML",

	Run: func(cmd *cobra.Command, args []string) {
		c := readConfig()
		server.Start(c)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVarP(&flagPath, "config", "c", "", "Path to the CV configuration. "+
		fmt.Sprintf("Overrides default path (%s)", defaultConfigPath))
}
