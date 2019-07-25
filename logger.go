package ken_logging

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var LevelMap = map[string]int{
	"DEBUG":   0,
	"WARNING": 1,
	"INFO":    2,
	"ERROR":   3,
}

type Logger struct {
	Name     string
	Level    string
	Handlers []*Handler
}

func (self *Logger) OutPut(curLevel string, v ...interface{}) {
	if LevelMap[curLevel] < LevelMap[self.Level] {
		return
	}
	prefix := "[" + curLevel + "] "

	formatStr := fmt.Sprintf("%v", v[0])
	var content string
	if strings.Contains(formatStr, "%") {
		content = fmt.Sprintf(fmt.Sprintf("%v", v[0]), v[1:]...)
	} else {
		content = fmt.Sprint(append([]interface{}{prefix}, v...)...)
	}
	for _, handler := range self.Handlers {
		err := handler.Output(curLevel, content)
		if err != nil {
			log.Println("写日志发生错误: ", err)
		}
	}
}

func (self *Logger) Debug(v ...interface{}) {
	self.OutPut("DEBUG", v...)
}

func (self *Logger) Warning(v ...interface{}) {
	self.OutPut("WARNING", v...)
}

func (self *Logger) Info(v ...interface{}) {
	self.OutPut("INFO", v...)
}

func (self *Logger) Error(v ...interface{}) {
	self.OutPut("ERROR", v...)
}

func (self *Logger) Exception(v ...interface{}) {
	self.Error(v...)
	os.Exit(1)
}

func NewLogger(name string, loggerConf *LoggerConf) *Logger {
	if name == "" {
		name = "root"
	}
	var handles []*Handler
	for _, v := range loggerConf.Handlers {
		h := GlobalConf.Handlers[v]
		handler := NewHandler(h)
		handles = append(handles, handler)
	}

	myLogger := Logger{
		Name:     name,
		Level:    loggerConf.Level,
		Handlers: handles,
	}
	return &myLogger
}
