package monitor

import (
	"os"

	"unknwon.dev/clog/v2"
)

func InitClog() {
	logLevel := clog.LevelTrace
	if os.Getenv("ENVIRONMENT") == "production" {
		logLevel = clog.LevelWarn
	}
	err := clog.NewConsole(100, clog.ConsoleConfig{
		Level: logLevel,
	})
	if err != nil {
		panic("unable to create new logger: " + err.Error())
	}
	clog.Trace("Clog initialized")
}
