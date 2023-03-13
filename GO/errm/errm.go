package errm

import (
	"log"
	"runtime"
)

func getFileAndLine() (string, int) {
	_, file, line, _ := runtime.Caller(2)
	return file, line
}

func LogErr(err error) {
	if err != nil {
		file, line := getFileAndLine()
		log.Printf("[%s:%d] ERROR: %s\n", file, line, err)
	}
}

func FatalErr(err error) {
	if err != nil {
		file, line := getFileAndLine()
		log.Fatalf("[%s:%d] FATAL ERROR: %s\n", file, line, err)
	}
}

func FuncExecIfErr(err error, f func()) {
	if err != nil {
		f()
	}
}
