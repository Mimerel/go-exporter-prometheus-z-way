package extract_data

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

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