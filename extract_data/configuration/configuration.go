package configuration

import (
	"github.com/Mimerel/go-logger-client"
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
		config.Logger = logs.New(config.Elasticsearch.Url, config.Host)
		config.Logger.Info("Configuration Loaded : %s", config)
	}
	return config
}