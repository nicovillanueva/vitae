package cmd

import (
	"github.com/nicovillanueva/api-vitae/server"
	"github.com/nicovillanueva/api-vitae/server/types"

	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var cvdata types.CVData

var rootCmd = &cobra.Command{
	Use:   "api-vitae",
	Short: "Start the server",

	Run: func(cmd *cobra.Command, args []string) {
		server.Start(cvdata)
	},
}

func init() {
	var data string
	var pdfPath string
	rootCmd.PersistentFlags().StringVarP(&data, "cv", "c", "", "Path to the CV data")
	rootCmd.PersistentFlags().StringVarP(&pdfPath, "pdf", "p", "", "Path to the downloadable PDF")
	cobra.OnInitialize(func() {
		viper.SetConfigType("yaml")
		viper.SetConfigFile(data)
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("fatal reading config: %+v", err))
		}
		if err := viper.Unmarshal(&cvdata); err != nil {
			panic(err)
		}
		var err error
		if cvdata.CVinPDF, err = ioutil.ReadFile(pdfPath); err != nil {
			panic(err)
		}
	})

}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
