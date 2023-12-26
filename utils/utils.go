package utils

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"syscall"
	"time"
)

// Regexp Match
func RegMatch(text string, check string) (match bool) {
	return regexp.MustCompile(check).MatchString(text)
}

// Regexp Replace
func RegReplace(fromText string, toText string, check string) (replaced string) {
	return regexp.MustCompile(check).ReplaceAllString(fromText, toText)
}

// Rand Generate 0~max-1
func Rand(max int) (result int) {
	result = rand.New(rand.NewSource(time.Now().UnixNano())).Int() % max
	return
}

// String Cut
func StrCut(text, suffix string, max int) (result string) {
	textArray := strings.Split(text, "")
	if len(textArray) < max {
		return text
	}
	for i := 0; i < max; i++ {
		result += textArray[i]
	}
	result += suffix
	return
}

// Listen Kill,Term,Interupt to Channel
func BreakSignal() (sc chan os.Signal) {
	sc = make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	return
}

// Error表示
func PrintError(message string, err error) (errored bool) {
	if err != nil {
		trackBack := ""
		// 原因を特定
		for i := 1; true; i++ {
			pc, file, line, _ := runtime.Caller(i)
			trackBack += fmt.Sprintf("> %s:%d %s()\n", filepath.Base(file), line, RegReplace(runtime.FuncForPC(pc).Name(), "", "^.*/"))
			_, _, _, ok := runtime.Caller(i + 3)
			if !ok {
				break
			}
			// インデント
			for j := 0; j < i; j++ {
				trackBack += "  "
			}
		}
		//表示
		SetPrintWordColor(255, 0, 0)
		fmt.Printf("[Error] Message:\"%s\" Error:\"%s\"\n", message, err.Error())
		fmt.Printf("%s", trackBack)
		ResetPrintWordColor()
		return true
	}
	return false
}

// Byte を intに
func ConvBtoI(b []byte) int {
	n := 0
	length := len(b)
	for i := 0; i < length; i++ {
		m := 1
		for j := 0; j < length-i-1; j++ {
			m = m * 256
		}
		m = m * int(b[i])
		n = n + m
	}
	return n
}
