package util

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var appLogger *logrus.Logger

func GetLogger() *logrus.Logger {

	logfile, err := os.Create("./gin_http.log")
	if err != nil {
		fmt.Println("Could not create log file")
	}

	logger := logrus.New()
	logger.Out = logfile
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger

}
