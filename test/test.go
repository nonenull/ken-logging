package main

import (
	logging "ken-logging"
)

var logger *logging.Logger

func init() {
	conf := `
{
  "handlers": {
    "console": {
      "level": "DEBUG"
    },
    "omsFile": {
      "level": "INFO",
      "filename": "oms.log"
    },
    "celery": {
      "level": "INFO",
      "filename": "celery.log"
    },
    "tasks": {
      "level": "INFO",
      "filename": "tasks.log"
    },
    "db_handle": {
      "level": "INFO",
      "filename": "db.log"
    }
  },
  "loggers": {
    "root": {
      "handlers": [
        "console",
        "omsFile"
      ],
      "level": "DEBUG"
    }
  }
}
`
	logging.SetGlobalConf(conf)
	logger = logging.GetLogger("test")
}

func main() {
	logger.Debug("just do IT. ", " fuckyuou man ", "fuck")
	logger.Info("just do IT 11. ", " fuckyuou man ", "fuck")
	logger.Error("just do IT 22. ", " fuckyuou man ", "fuck")
	for {
	}
}
