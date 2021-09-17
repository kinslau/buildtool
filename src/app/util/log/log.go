package main

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(logrus.TraceLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	log.Trace("trace msg")
	log.Debug("debug msg")
	log.Info("info msg")
	log.Warn("warn msg")
	log.Error("error msg")
	log.Fatal("fatal msg")
	log.Panic("panic msg")
}
