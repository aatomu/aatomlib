package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type LogLevel int

const (
	Info LogLevel = iota
	Warn
	Error
)

type LoggerHandler struct {
	Level LogLevel
}

func (l *LoggerHandler) Info(arg ...any) {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	text := fmt.Sprintf("[%s] INFO %s(%s:%d)   %s\n", time.Now().Format("2006-01-02T15:04:05.000"), f.Name(), filepath.Base(file), line, fmt.Sprint(arg...))

	l.logPrint(Info, text)
}

func (l *LoggerHandler) Warn(arg ...any) {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	text := fmt.Sprintf("[%s] Warn %s(%s:%d)   %s\n", time.Now().Format("2006-01-02T15:04:05.000"), f.Name(), filepath.Base(file), line, fmt.Sprint(arg...))

	l.logPrint(Warn, text)
}

func (l *LoggerHandler) Error(arg ...any) {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	text := fmt.Sprintf("[%s] ERROR %s(%s:%d)   %s\n", time.Now().Format("2006-01-02T15:04:05.000"), f.Name(), filepath.Base(file), line, fmt.Sprint(arg...))

	l.logPrint(Error, text)
}

func (l *LoggerHandler) logPrint(level LogLevel, text string) {
	if l.Level >= level {
		io.WriteString(os.Stdout, text)
	} else {
		io.WriteString(os.Stderr, text)
	}
}
