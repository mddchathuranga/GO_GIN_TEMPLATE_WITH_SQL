package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *logrus.Logger

func InitLogger(fPath string, maxSize int, maxBackups int, maxAge int) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		logrus.Fatal("Failed to get the current working directory:", err)
	}
	// Define the relative path to the logs folder
	relativeLogFolderPath := "logs"
	// Join the current directory and the relative path to get the absolute logs folder path
	logsFolderPath := filepath.Join(currentDir, relativeLogFolderPath)
	// Ensure the logs folder exists or create it
	if _, err := os.Stat(logsFolderPath); os.IsNotExist(err) {
		if err := os.Mkdir(logsFolderPath, 0755); err != nil {
			logrus.Fatalf("Fail to create logs folder:%s", err)
		}

	}
	var filePathVal string

	if fPath == "default" {
		filePathVal = getLogFilePath(logsFolderPath)
	} else {
		filePathVal = fPath
	}

	// Create a new lumberjack logger with rotation settings

	logger := &lumberjack.Logger{
		Filename:   filePathVal,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}

	// Create a new logrus logger
	Log = logrus.New()

	// Create a multi-writer that writes to both the file and the console
	multiWriter := io.MultiWriter(logger, os.Stdout)

	// Set the log output to the multi-writer for both file and console logging
	Log.SetOutput(multiWriter)

	// Set log level (e.g., "info", "warn", "error", "debug", etc.)
	Log.SetLevel(logrus.InfoLevel)

	Log.Info("Logger initialized.")

}

// GetLogger returns the initialized logger instance
func GetLogger() *logrus.Logger {
	return Log
}

func getLogFilePath(logsFolderPath string) string {
	currentTime := time.Now().Format("2006-01-02_15-04-05") // Format: YYYY-MM-DD_HH-MM-SS
	return filepath.Join(logsFolderPath, fmt.Sprintf("app_%s.log", currentTime))
}
