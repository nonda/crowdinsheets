package main

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// CrowdinSheetsConfig Project config, see crowdin.yml.exmaple for details
type CrowdinSheetsConfig struct {
	ProjectID    string `yaml:"projectId,omitempty"`
	APIToken     string `yaml:"apiToken,omitempty"`
	Languages    []string
	Files        []string
	OutputFolder string `yaml:"outputFolder"`
}

// ReadConfig reads specified yml config file, see crowdin.yml.exmaple for details
func ReadConfig(filename string) (CrowdinSheetsConfig, error) {
	filenameAbs, err := filepath.Abs(filename)
	if err != nil {
		return CrowdinSheetsConfig{}, err
	}

	ymlSource, err := ioutil.ReadFile(filenameAbs)
	if err != nil {
		return CrowdinSheetsConfig{}, err
	}

	var config CrowdinSheetsConfig
	err = yaml.Unmarshal(ymlSource, &config)
	return config, err
}
