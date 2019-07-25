package ken_logging

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
)

var GlobalConf = &LogConf{}

type LoggerConf struct {
	Handlers []string `yaml:"handlers"`
	Level    string   `yaml:"level"`
}

type HandlerConf struct {
	Level    string `yaml:"level"`
	Filename string `yaml:"filename"`
}

type LogConf struct {
	Loggers  map[string]*LoggerConf  `yaml:"loggers"`
	Handlers map[string]*HandlerConf `yaml:"handlers"`
}

func SetGlobalConf(confType string, configStr []byte) {
	switch confType {
	case "yaml":
		if err := yaml.Unmarshal(configStr, GlobalConf); err != nil {
			log.Fatal("parse yaml log config  error: ", err)
		}
		break
	case "json":
		if err := json.Unmarshal(configStr, GlobalConf); err != nil {
			log.Fatal("parse json log config  error: ", err)
		}
		break
	default:
		log.Fatal("Only json and yaml are supported, current conf type : ", confType)
	}
}

func SetGlobalConfFormFile(filePath string) {
	conf, err := ioutil.ReadFile(path.Join(GetCurrentDirectory(), filePath))
	if err != nil {
		log.Fatal("open log config file error: ", err)
	}
	fileSuffix := path.Ext(filePath)
	if fileSuffix == ".yml" {
		fileSuffix = ".yaml"
	}
	SetGlobalConf(fileSuffix[1:], conf)
}
