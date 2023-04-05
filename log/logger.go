package log

import (
	"log"
	"os"
)

var logger *log.Logger

func Initialize(file *os.File) {
	logger = log.New(file, "", log.LstdFlags)
}

func Default() *log.Logger {
	return logger
}
