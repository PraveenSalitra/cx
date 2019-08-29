package aws

import (
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func UpdateRegion(region string) {
	//config
	home, _ := homedir.Dir()
	awsConfigPath := filepath.Join(home, ".aws", "config")

	var awsConfig []byte
	var awsConfigLines []string

	awsConfigLines = append(awsConfigLines, "[default]")
	awsConfigLines = append(awsConfigLines, "output = json")
	awsConfigLines = append(awsConfigLines, "region = "+region)

	awsConfig = []byte (strings.Join(awsConfigLines, "\n"))
	err := ioutil.WriteFile(awsConfigPath, awsConfig, 0600)
	if err != nil {
		panic(err)
	}
}
