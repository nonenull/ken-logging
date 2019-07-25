package ken_logging

import (
	"log"
	"os"
)

type Handler struct {
	logger    *log.Logger
	level     string
	formatter string
	filename  string
}

func (self *Handler) createLog() {
	if self.filename == "" {
		self.logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	} else {
		checkPath(self.filename)
		logFile, err := os.OpenFile(self.filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			log.Fatal("open log file error: ", err)
		}
		self.logger = log.New(logFile, "", log.LstdFlags|log.Lshortfile)
	}
}

func (self *Handler) Output(level string, s string) error {
	if LevelMap[level] < LevelMap[self.level] {
		return nil
	}
	return self.logger.Output(4, s)
}

func NewHandler(handlerConf *HandlerConf) *Handler {
	h := &Handler{
		level:    handlerConf.Level,
		filename: handlerConf.Filename,
	}
	h.createLog()
	return h
}
