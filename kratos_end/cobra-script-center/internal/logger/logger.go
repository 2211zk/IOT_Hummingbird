package logger

import (
	"os"

	"cobra-script-center/internal/config"

	"github.com/sirupsen/logrus"
)

// Logger is the global logger instance
var Logger *logrus.Logger

// Init initializes the logger with the given configuration
func Init(cfg *config.LoggingConfig) error {
	Logger = logrus.New()

	// Set log level
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		return err
	}
	Logger.SetLevel(level)

	// Set log format
	if cfg.Format == "json" {
		Logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	// Set output
	if cfg.File != "" {
		file, err := os.OpenFile(cfg.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		Logger.SetOutput(file)
	} else {
		Logger.SetOutput(os.Stdout)
	}

	return nil
}

// GetLogger returns the global logger instance
func GetLogger() *logrus.Logger {
	if Logger == nil {
		Logger = logrus.New()
	}
	return Logger
}
