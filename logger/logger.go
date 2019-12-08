package logger

import (
	"fmt"
)

var logLevel string

func Init(level string) {
	logLevel = level
}

func Info(a ...interface{}) {
	fmt.Println(a...)
}

func Debug(a ...interface{}) {
	if logLevel != "debug" {
		return
	}
	fmt.Println(a...)
}
