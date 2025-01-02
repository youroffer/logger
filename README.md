Настройка формата zerolog для всего проекта

```bash
{"level":"info","time":"07.12.2024 18:50:00","caller":"/logger/example/main.go:14","message":"Formatted info message: golang"}
{"level":"warn","error":"warn err","module":"main","request_id":1,"status":"initialized","time":"07.12.2024 18:50:00","caller":"/logger/example/main.go:24","message":"Warn message with fields"}

18:50:00 ERR example/main.go:30 > request_id=1 error="err error" module=main status=initialized
18:50:00 ERR example/main.go:31 > Formatted err message: msg error="err error"
18:50:00 DBG example/main.go:34 > Debug message with fields request_id=1 module=main status=initialized
18:50:00 DBG example/main.go:35 > Formatted debug message with fields: debugging request_id=1 module=main status=initialized
```
