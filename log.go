package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Fields map[string]interface{}

const (
	RequestID = "request_id"
)

// SetupLogger initializes the global logger with the specified log level
func SetupLogger(level string) {
	zlevel, err := zerolog.ParseLevel(level)
	if err != nil {
		zlevel = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(zlevel)
	zerolog.TimeFieldFormat = "02.01.2006 15:04:05"

	switch zlevel {
	case zerolog.DebugLevel:
		log.Logger = zerolog.New(zerolog.ConsoleWriter{
			Out:         os.Stderr,
			TimeFormat:  "15:04:05",
			FieldsOrder: []string{RequestID}}).
			With().
			Timestamp().
			CallerWithSkipFrameCount(3).
			Logger()
	default:
		log.Logger = zerolog.New(os.Stdout).
			With().
			Timestamp().
			CallerWithSkipFrameCount(3).
			Logger()
	}
}
