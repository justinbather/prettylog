package main

import "github.com/justinbather/prettylog"

func main() {
	logger := prettylog.New()

	logger.Info("This is a log", "someKey", "someVal")
}
