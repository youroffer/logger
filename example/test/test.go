package test

import log "github.com/youroffer/logger"

func ShowLogsAgain() {
	log.Trace("Hello, Trace!")
	log.Info("Hello, World!")
	log.Debug("Debugging message")
	log.Warn("Warn")
	log.Error("Error")
	log.Fatal("Fatlse")

}
