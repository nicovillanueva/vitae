package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nicovillanueva/vitae/server/types"
	"gopkg.in/yaml.v2"
)

/*
	This is a mess.
*/

func readConfig() types.CVData {
	var cvconfig types.CVData

	envPath, _ := os.LookupEnv(configEnvVariable)

	if flagPath != "" { // If a path is provided via flag, it's first priority
		readConfigFile(flagPath, &cvconfig)
	} else if envPath != "" { // otherwise try with the environment variable
		readConfigFile(envPath, &cvconfig)
	} else { // default to ./cv-data.yaml
		readConfigFile(defaultConfigPath, &cvconfig)
	}

	if p := os.Getenv("CV_PDF_PATH"); p != "" {
		// Override PDF path declared in YAML
		cvconfig.PathToPDF = p
	}
	fmt.Printf("Reading pdf from: %s\n", cvconfig.PathToPDF)
	if b, err := ioutil.ReadFile(cvconfig.PathToPDF); err != nil {
		panic(err)
	} else {
		cvconfig.PDFpayload = b
	}

	return cvconfig
}

func readConfigFile(pathToConfig string, cfg *types.CVData) error {
	fmt.Printf("Reading config from: %s\n", pathToConfig)
	var b []byte
	var err error
	if b, err = ioutil.ReadFile(pathToConfig); err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(b, cfg); err != nil {
		panic(err)
	}
	return nil
}
