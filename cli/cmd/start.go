package cmd

import (
	"io/ioutil"

	"github.com/nicovillanueva/vitae/server"
	"github.com/nicovillanueva/vitae/server/types"
	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

var cvdata types.CVData

var cvPath string
var pdfPath string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if cvdata.CVinPDF, err = ioutil.ReadFile(pdfPath); err != nil {
			panic(err)
		}
		if c, err := ioutil.ReadFile(cvPath); err != nil {
			panic(err)
		} else if err := yaml.Unmarshal(c, &cvdata); err != nil {
			panic(err)
		}

		server.Start(cvdata)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVarP(&cvPath, "cv", "c", "", "Path to the CV data")
	startCmd.Flags().StringVarP(&pdfPath, "pdf", "p", "", "Path to the downloadable PDF")
	startCmd.MarkFlagRequired("cv")
	startCmd.MarkFlagRequired("pdf")
}
