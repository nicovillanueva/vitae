package cmd

import (
	"github.com/nicovillanueva/api-vitae/server"
	"github.com/nicovillanueva/api-vitae/server/types"

	"fmt"
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
	rootCmd.PersistentFlags().StringVarP(&data, "cv", "c", "", "Path to the CV data")
	rootCmd.MarkFlagRequired("cv")
	cobra.OnInitialize(func() {
		viper.SetConfigType("yaml")
		viper.SetConfigFile(data)
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("fatal reading config: %+v", err))
		}
		if err := viper.Unmarshal(&cvdata); err != nil {
			panic(err)
		}
		// fmt.Printf("unmarshalled: %+v", cvdata)
	})

}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
