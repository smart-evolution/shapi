package utils

import (
    "log"
    "runtime"
)

// Log - print logs int standard output
func Log(msg interface{}, args ...interface{}) {
    pc, file, line, ok := runtime.Caller(1)
    if ok {
        log.Println(
            file,
            runtime.FuncForPC(pc).Name(),
            line,
            msg,
        )
    }
}
