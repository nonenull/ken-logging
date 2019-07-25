package ken_logging

var LoggerDict = make(map[string]*Logger)

func GetLogger(name string) *Logger {
	var logger *Logger
	logger, ok := LoggerDict[name]
	if !ok {
		loggerConf, ok := GlobalConf.Loggers[name]
		if !ok {
			loggerConf, _ = GlobalConf.Loggers["root"]
		}
		logger = NewLogger(name, loggerConf)
	}
	LoggerDict[name] = logger
	return logger
}
