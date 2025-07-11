package logger

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/bongochat/utils/resterrors"
)

var (
	fileLogger  *log.Logger
	once        sync.Once
	logFilePath = filepath.Join("logs", "debug.log") // logs/debug.log
)

// initialize file logger once
func initFileLogger() {
	once.Do(func() {
		// Create logs directory if it doesn't exist
		if err := os.MkdirAll(filepath.Dir(logFilePath), os.ModePerm); err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}

		// Open log file in append mode, create if not exists
		file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}

		// Create a new logger instance
		fileLogger = log.New(file, "", log.LstdFlags|log.Lshortfile)
	})
}

// InfoLog writes an info message
func InfoLog(message string) {
	initFileLogger()
	fileLogger.Println("[INFO] " + message)
}

// ErrorMsgLog logs error messages
func ErrorMsgLog(message string) {
	initFileLogger()
	fileLogger.Println("[ERROR] " + message)
}

// ErrorLog logs error objects
func ErrorLog(err error) {
	initFileLogger()
	fileLogger.Println("[ERROR] ", err)
}

// RestErrorLog logs custom RestError
func RestErrorLog(err resterrors.RestError) {
	initFileLogger()
	fileLogger.Println("[ERROR] ", err)
}
