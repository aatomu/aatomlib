package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type loggerHandler struct {
	directory string
}

func Logger(dir string) (l *loggerHandler, err error) {
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return
	}
	l = &loggerHandler{
		directory: dir,
	}
	return
}

func (l *loggerHandler) Info(arg ...any) {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	text := fmt.Sprintf("[%s] INFO %s(%s:%d) Message:%s", time.Now().Format("2006-01-02T15:04:05.000"), f.Name(), filepath.Base(file), line, fmt.Sprint(arg...))

	io.WriteString(os.Stdout, text)
}

func (l *loggerHandler) Error(arg ...any) {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	text := fmt.Sprintf("[%s] ERROR %s(%s:%d) Message:%s", time.Now().Format("2006-01-02T15:04:05.000"), f.Name(), filepath.Base(file), line, fmt.Sprint(arg...))

	io.WriteString(os.Stderr, text)
}
