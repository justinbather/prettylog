package main

import (
	"github.com/justinbather/prettylog"
)

func main() {
	logger := prettylog.New()

	logger.Info("This is an info log", "add string")
	logger.Infof("This is a formatted info %s", "log")
	logger.Debug("This is a debug log")
	logger.Warn("This is a warn log")
	logger.Error("This is an error log")
}
