package configuration

import (
	"github.com/apsdehal/go-logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var log, err = logger.New("test", 1, os.Stdout)


func ReadConfiguration() (MainConfig){
	pathToFile := os.Getenv("EXPORTER_CONFIGURATION_FILE")
	if pathToFile == "" {
		pathToFile = "go-exporter-conf.yaml"
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
		log.NoticeF("Configuration Loaded : %+v", config)
	}
	return config

}