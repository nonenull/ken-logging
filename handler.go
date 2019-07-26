package ken_logging

import (
	"io"
	"log"
	"os"
)

type Handler struct {
	Logger   *log.Logger
	level    string
	console  bool
	filename string
}

func (self *Handler) createLog() {
	var output io.Writer
	if self.filename == "" {
		output = os.Stdout
	} else {
		checkPath(self.filename)
		var err error
		output, err = os.OpenFile(self.filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			log.Fatal("open log file error: ", err)
		}
	}
	if self.console {
		output = io.MultiWriter(os.Stdout, output)
	}
	self.Logger = log.New(output, "", log.LstdFlags|log.Lshortfile)
}

func (self *Handler) Output(level string, s string) error {
	if LevelMap[level] < LevelMap[self.level] {
		return nil
	}
	return self.Logger.Output(4, s)
}

func NewHandler(handlerConf *HandlerConf) *Handler {
	h := &Handler{
		level:    handlerConf.Level,
		console:  handlerConf.Console,
		filename: handlerConf.Filename,
	}
	h.createLog()
	return h
}
