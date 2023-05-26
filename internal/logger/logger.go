package logger

import "log"

type Logger struct{}

func New() Logger {
	return Logger{}
}

func (logger Logger) Log(msg string) {
	log.Println(msg)
}
