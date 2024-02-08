package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type LoggerHandler struct {
	directory string
}

func Logger(dir string) (l *LoggerHandler, err error) {
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return
	}
	l = &LoggerHandler{
		directory: dir,
	}
	return
}

func (l *LoggerHandler) Info(arg ...any) {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	text := fmt.Sprintf("[%s] INFO %s(%s:%d) Message:%s", time.Now().Format("2006-01-02T15:04:05.000"), f.Name(), filepath.Base(file), line, fmt.Sprint(arg...))

	io.WriteString(os.Stdout, text)
}

func (l *LoggerHandler) Warn(arg ...any) {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	text := fmt.Sprintf("[%s] Warn %s(%s:%d) Message:%s", time.Now().Format("2006-01-02T15:04:05.000"), f.Name(), filepath.Base(file), line, fmt.Sprint(arg...))

	io.WriteString(os.Stdout, text)
}

func (l *LoggerHandler) Error(arg ...any) {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	text := fmt.Sprintf("[%s] ERROR %s(%s:%d) Message:%s", time.Now().Format("2006-01-02T15:04:05.000"), f.Name(), filepath.Base(file), line, fmt.Sprint(arg...))

	io.WriteString(os.Stderr, text)
}
