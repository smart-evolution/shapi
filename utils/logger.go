package utils

import (
	"log"
	"runtime"
)

// Log - print logs int standard output
func Log(args ...interface{}) {
	pc, _, line, ok := runtime.Caller(1)

	if ok {
		name := runtime.FuncForPC(pc).Name()
		params := append([]interface{}{name, line}, args[0])

		for _, v := range args[1:] {
			params = append(params, " / ", v)
		}

		log.Println(
			params...,
		)
		return
	}

	log.Println(
		args...,
	)
}
