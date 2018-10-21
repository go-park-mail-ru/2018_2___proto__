package router

import (
	"log"
)

type ILogger interface {
	Debugf(msg string, v ...interface{})

	Info(v ...interface{})

	Notice(v ...interface{})

	Warning(v ...interface{})

	Error(v ...interface{})

	Critical(v ...interface{})
}

type DafaultLogger struct {

}

func NewDefaultLogger() *DafaultLogger {
	return &DafaultLogger{}
}

func (l *DafaultLogger) Debugf(msg string, v ...interface{}) {
	log.Printf(msg, v)
}

func (l *DafaultLogger) Info(v ...interface{}) {
	log.Println(v)
}

func (l *DafaultLogger) Notice(v ...interface{}) {
	log.Println(v)
}

func (l *DafaultLogger) Warning(v ...interface{}) {
	log.Println(v)
}

func (l *DafaultLogger) Error(v ...interface{}) {
	log.Println(v)
}

func (l *DafaultLogger) Critical(v ...interface{}) {
	log.Printf("CRITICAL %v", v)
}