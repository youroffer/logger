Обертка над zerolog для использования единого формата логгера в микросервисах проекта youroffer.

## Example

```go
package main

import (
	"errors"

	log "github.com/youroffer/logger"
)

func main() {
	// Инициализация логгера с уровнем Info
	log.SetupLogger("info")

	log.Infof("Formatted info message: %s", "golang")

	// Логируем с полями
	fields := log.Fields{
		"module": "main",
		"status": "initialized",
		log.RequestID: 1,
	}

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
```

```bash
{"level":"info","time":"07.12.2024 18:50:00","caller":"/logger/example/main.go:14","message":"Formatted info message: golang"}
{"level":"warn","error":"warn err","module":"main","request_id":1,"status":"initialized","time":"07.12.2024 18:50:00","caller":"/logger/example/main.go:24","message":"Warn message with fields"}

18:50:00 ERR example/main.go:30 > request_id=1 error="err error" module=main status=initialized
18:50:00 ERR example/main.go:31 > Formatted err message: msg error="err error"
18:50:00 DBG example/main.go:34 > Debug message with fields request_id=1 module=main status=initialized
18:50:00 DBG example/main.go:35 > Formatted debug message with fields: debugging request_id=1 module=main status=initialized
```
