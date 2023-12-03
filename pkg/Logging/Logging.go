package Logging

import (
	"fmt"
	"reflect"
	"time"
)

var Colors = struct {
	Default, Black, Red, Green, Yellow, Blue, Magenta, Cyan, White, Gray string
}{
	Default: "\033[0m",
	Black:   "\033[30m",
	Red:     "\033[31m",
	Green:   "\033[32m",
	Yellow:  "\033[33m",
	Blue:    "\033[34m",
	Magenta: "\033[35m",
	Cyan:    "\033[36m",
	White:   "\033[37m",
	Gray:    "\033[90m",
}

const (
	Info = iota
	Trace
	Debug
	Warn
	Error
)

type Log struct {
	LogLevel uint8
}

func (l Log) logParent(logLevel uint8, message interface{}) {
	if l.LogLevel > logLevel {
		return
	}

	var prefix string

	switch logLevel {
	case Info:
		prefix = Colors.Gray + "[Info  " + time.Now().Format(time.RFC822) + "] " + Colors.Default
	case Trace:
		prefix = Colors.Cyan + "[Trace " + time.Now().Format(time.RFC822) + "] " + Colors.Default
	case Debug:
		prefix = Colors.Magenta + "[Debug " + time.Now().Format(time.RFC822) + "] " + Colors.Default
	case Warn:
		prefix = Colors.Yellow + "[Warn  " + time.Now().Format(time.RFC822) + "] " + Colors.Default
	case Error:
		prefix = Colors.Red + "[Error " + time.Now().Format(time.RFC822) + "] " + Colors.Default
	}

	fmt.Println(prefix, message, Colors.Gray+"("+reflect.TypeOf(message).String()+")")
}

func (l Log) Info(message interface{}) {
	l.logParent(Info, message)
}

func (l Log) Trace(message interface{}) {
	l.logParent(Trace, message)
}

func (l Log) Debug(message interface{}) {
	l.logParent(Debug, message)
}

func (l Log) Warn(message interface{}) {
	l.logParent(Warn, message)
}

func (l Log) Error(message string) {
	l.logParent(Error, message)
}
