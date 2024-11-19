package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
			Out:        os.Stderr,
			TimeFormat: "15:04:05"}).
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

// Debug logs a debug message
func Debug(msg string) {
	log.Debug().Msg(msg)
}

// DebugMsgf logs a formatted debug message
func Debugf(msg string, args ...any) {
	log.Debug().Msgf(msg, args...)
}

// DebugFields logs a debug message with fields
func DebugFields(fields any, msg string) {
	log.Debug().Fields(fields).Msg(msg)
}

// DebugFieldsMsgf logs a debug formatted message with fields
func DebugFieldsf(fields any, msg string, args ...any) {
	log.Debug().Fields(fields).Msgf(msg, args...)
}

// Info logs an message
func Info(msg string) {
	log.Info().Msg(msg)
}

// Infof logs a formatted message
func Infof(msg string, args ...any) {
	log.Info().Msgf(msg, args...)
}

// InfoFields logs an message with fields
func InfoFields(fields any, msg string) {
	log.Info().Fields(fields).Msg(msg)
}

// InfoFieldsf logs a formatted message with fields
func InfoFieldsf(fields any, msg string, args ...any) {
	log.Info().Fields(fields).Msgf(msg, args...)
}

// Warn logs a warning error
func Warn(err error) {
	log.Warn().Err(err)
}

// WarnFields logs a warning error with fields
func WarnFields(err error, fields any) {
	log.Warn().Err(err).Fields(fields)
}

// WarnMsg logs a warning error with a message
func WarnMsg(err error, msg string) {
	log.Warn().Err(err).Msg(msg)
}

// WarnMsgf logs a warning error with a formatted message
func WarnMsgf(err error, msg string, args ...any) {
	log.Warn().Err(err).Msgf(msg, args...)
}

// WarnFieldsMsg logs a warning error with fields and a message
func WarnFieldsMsg(err error, fields any, msg string) {
	log.Warn().Err(err).Fields(fields).Msg(msg)
}

// WarnFieldsMsgf logs a warning error with fields and a formatted message
func WarnFieldsMsgf(err error, fields any, msg string, args ...any) {
	log.Warn().Err(err).Fields(fields).Msgf(msg, args...)
}

// Err logs an error
func Err(err error) {
	log.Error().Err(err)
}

// ErrMsg logs an error with a message
func ErrMsg(err error, msg string) {
	log.Error().Err(err).Msg(msg)
}

// ErrMsgf logs an error with a formatted message
func ErrMsgf(err error, msg string, args ...any) {
	log.Error().Err(err).Msgf(msg, args...)
}

// ErrFields logs an error with fields
func ErrFields(err error, fields any) {
	log.Error().Err(err).Fields(fields)
}

// ErrFieldsMsg logs an error with fields and a message
func ErrFieldsMsg(err error, fields any, msg string) {
	log.Error().Err(err).Fields(fields).Msg(msg)
}

// ErrFieldsMsgf logs an error with fields and a formatted message
func ErrFieldsMsgf(err error, fields any, msg string, args ...any) {
	log.Error().Err(err).Fields(fields).Msgf(msg, args...)
}

// Fatal logs a fatal error 
func Fatal(err error) {
	log.Fatal().Err(err)
}

// FatalMsg logs a fatal error  with a message
func FatalMsg(err error, msg string) {
	log.Fatal().Err(err).Msg(msg)
}

// FatalMsgf logs a fatal error with a formatted message
func FatalMsgf(err error, msg string, args ...any) {
	log.Fatal().Err(err).Msgf(msg, args...)
}

// FatalFields logs a fatal error with fields
func FatalFields(err error, fields any) {
	log.Fatal().Err(err).Fields(fields)
}

// FatalFieldsMsg logs a fatal error with fields and a message
func FatalFieldsMsg(err error, fields any, msg string) {
	log.Fatal().Err(err).Fields(fields).Msg(msg)
}

// FatalFieldsMsgf logs a fatal error with fields and a formatted message
func FatalFieldsMsgf(err error, fields any, msg string, args ...any) {
	log.Fatal().Err(err).Fields(fields).Msgf(msg, args...)
}
