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

	log.InfoMsg("Info message with message")
	log.InfoMsgf("Formatted info message: %s", "golang")

	// Логируем с полями
	fields := map[string]interface{}{
		"module": "main",
		"status": "initialized",
	}
	log.WarnFields(errors.New("warn err"), fields)

	// Инициализация логгера с уровнем Debug
	log.SetupLogger("debug")

	// Логируем с уровнем Debug
	log.Err(errors.New("err error"))
	log.ErrMsgf(errors.New("err error"), "Formatted err message: %s", "msg")

	// Логируем с полями 
	log.DebugFieldsf(fields, "Formatted debug message with fields: %s", "debugging")
}
```

```bash
{"level":"info","time":"19.11.2024 02:24:20","caller":"/logger/example/main.go:14","message":"Info message with message"}
{"level":"info","time":"19.11.2024 02:24:20","caller":"/logger/example/main.go:15","message":"Formatted info message: golang"}

02:24:20 ERR example/main.go:30 > Formatted err message: msg error="err error"
02:24:20 DBG example/main.go:34 > Formatted debug message with fields: debugging module=main status=initialized
```
