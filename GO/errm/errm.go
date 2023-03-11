package errm

import (
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

var racinefillename string

func init() {
	_, file2, _, _ := runtime.Caller(0)
	rfille := (filepath.Dir(filepath.Dir(filepath.Dir(file2))))
	racinefillename = filepath.Base(rfille)
}

func LogErr(err error) {
	if err != nil {
		file, line := getPath()
		log.Printf("[%s:%d] ERROR: %s\n", file, line, err)
	}
}

func getPath() (string, int) {
	_, file, line, _ := runtime.Caller(2)
	start := false
	filep := []string{}
	for _, filename := range strings.Split(file, "/") {
		if filename == racinefillename {
			start = true
		}
		if start {
			filep = append(filep, filename)
		}
	}
	return strings.Join(filep, "/"), line
}
