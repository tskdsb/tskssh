package util

import (
	"log"
	"runtime"
)

func LogErr(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d %s\n", file, line, err)
	}
}
