package ken_logging

import (
	"encoding/json"
	"log"
)

var GlobalConf *LogConf

type FormatterConf struct {
	Format  string `yaml:"format"`
	Datefmt string `yaml:"datefmt"`
}

type LoggerConf struct {
	Handlers []string `yaml:"handlers"`
	Level    string   `yaml:"level"`
}

type HandlerConf struct {
	Level     string `yaml:"level"`
	Formatter string `yaml:"formatter"`
	Filename  string `yaml:"filename"`
}

type LogConf struct {
	Loggers    map[string]*LoggerConf    `yaml:"logger"`
	Formatters map[string]*FormatterConf `yaml:"formatters"`
	Handlers   map[string]*HandlerConf   `yaml:"handlers"`
}

func SetGlobalConf(configJSON string) {
	if err := json.Unmarshal([]byte(configJSON), &GlobalConf); err != nil {
		log.Fatal("log config load error: ", err)
	}
}
