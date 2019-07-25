package main

import (
	logging "ken-logging"
)

var logger *logging.Logger

func init() {
	logging.SetGlobalConfFormFile("test/log.yml")
	//logging.SetGlobalConfFormFile("test/log.json")
	logger = logging.GetLogger("test")
}

func main() {
	logger.Debug("just do IT. ", " fuckyuou man ", "fuck")
	logger.Info("just do IT 11. ", " fuckyuou man ", "fuck")
	logger.Error("just do IT 22. ", " fuckyuou man ", "fuck")
	for {
	}
}
