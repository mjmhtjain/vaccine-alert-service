package logger

import (
	"log"
	"os"
)

var (
	WARN  *log.Logger
	INFO  *log.Logger
	ERROR *log.Logger
	DEBUG *log.Logger
)

func init() {
	INFO = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WARN = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ERROR = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	DEBUG = log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
