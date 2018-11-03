package configuration

import (
	"fmt"
	"go-exporter-prometheus-z-way/extract_data/logs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func ReadConfiguration() (MainConfig){
	pathToFile := os.Getenv("EXPORTER_CONFIGURATION_FILE")
	if pathToFile == "" {
		pathToFile = "/home/pi/go/src/go-exporter-prometheus-z-way/go-exporter-conf.yaml"
	}
	yamlFile, err := ioutil.ReadFile(pathToFile)

	if err != nil {
		panic(err)
	}

	var config MainConfig

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	} else {
		logs.Info(config.Logger, config.Host, fmt.Sprint("Configuration Loaded : %s", config))
	}
	return config
}