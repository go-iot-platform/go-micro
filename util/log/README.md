# Log

DEPRECATED: use github.com/go-iot-platform/go-micro/logger interface

This is the global logger for all micro based libraries.

## Set Logger

Set the logger for micro libraries

```go
// import go-micro/util/log
import "github.com/go-iot-platform/go-micro/util/log"

// SetLogger expects github.com/go-iot-platform/go-micro/debug/log.Log interface
log.SetLogger(mylogger)
```
