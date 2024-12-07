package main

import (
	"errors"

	log "github.com/youroffer/logger"
)

func main() {
	// Инициализация логгера с уровнем Info
	log.SetupLogger("info")

	log.Info("This is an info message")
	log.Infof("Formatted info message: %s", "golang")

	// Логируем с полями
	fields := log.Fields{
		"module":     "main",
		"status":     "initialized",
		"request_id": 1,
	}

	log.WarnFields(errors.New("warn err"), fields)
	log.WarnFieldsMsgf(errors.New("warn err"), fields, "Warn message with fields")

	// Инициализация логгера с уровнем Debug
	log.SetupLogger("debug")

	// Логируем с уровнем Debug
	log.ErrFields(errors.New("err error"), fields)
	log.ErrMsgf(errors.New("err error"), "Formatted err message: %s", "msg")

	// Логируем с полями
	log.DebugFields("Debug message with fields", fields)
	log.DebugFieldsf(fields, "Formatted debug message with fields: %s", "debugging")
}
