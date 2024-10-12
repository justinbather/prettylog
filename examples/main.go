package main

import "github.com/justinbather/prettylog"

func main() {
	logger := prettylog.New()

	logger.Info("This is an info log", "someKey", "someVal")
	logger.Warn("This is a warn log", "someKey", "someVal")
	logger.Error("This is an error log", "someKey", "someVal")
}
