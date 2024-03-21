package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// The global logger for all Client instances
var Logger = getLogger(os.Getenv("LOGLEVEL"))

func SetLogLevel(value string) {
	Logger = getLogger(value)
}

func getLogger(logLevel string) zerolog.Logger {
	var level zerolog.Level
	if logLevel != "" {
		parsedLevel, err := zerolog.ParseLevel(logLevel)
		if err != nil {
			panic("Invalid log level: " + logLevel)
		}
		level = parsedLevel
	} else {
		level = zerolog.InfoLevel
	}

	return zerolog.New(os.Stderr).Level(level)
}
