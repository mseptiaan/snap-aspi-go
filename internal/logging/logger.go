package logging

import (
	"os"
	"strings"

	"github.com/mseptiaan/snap-aspi-go/internal/config"
	"github.com/sirupsen/logrus"
)

// Logger interface for structured logging
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})

	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
	WithError(err error) Logger
}

// LogrusLogger wraps logrus.Logger to implement our Logger interface
type LogrusLogger struct {
	*logrus.Logger
	entry *logrus.Entry
}

// NewLogger creates a new structured logger based on configuration
func NewLogger(cfg *config.LoggerConfig) Logger {
	logger := logrus.New()

	// Set log level
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	// Set log format
	switch strings.ToLower(cfg.Format) {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		})
	case "text":
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		})
	default:
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		})
	}

	// Set output destination
	switch strings.ToLower(cfg.Output) {
	case "stdout":
		logger.SetOutput(os.Stdout)
	case "stderr":
		logger.SetOutput(os.Stderr)
	default:
		logger.SetOutput(os.Stdout)
	}

	return &LogrusLogger{
		Logger: logger,
		entry:  logrus.NewEntry(logger),
	}
}

// WithField adds a single field to the logger
func (l *LogrusLogger) WithField(key string, value interface{}) Logger {
	return &LogrusLogger{
		Logger: l.Logger,
		entry:  l.entry.WithField(key, value),
	}
}

// WithFields adds multiple fields to the logger
func (l *LogrusLogger) WithFields(fields map[string]interface{}) Logger {
	return &LogrusLogger{
		Logger: l.Logger,
		entry:  l.entry.WithFields(logrus.Fields(fields)),
	}
}

// WithError adds an error field to the logger
func (l *LogrusLogger) WithError(err error) Logger {
	return &LogrusLogger{
		Logger: l.Logger,
		entry:  l.entry.WithError(err),
	}
}

// Debug logs debug level message
func (l *LogrusLogger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

// Info logs info level message
func (l *LogrusLogger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

// Warn logs warning level message
func (l *LogrusLogger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

// Error logs error level message
func (l *LogrusLogger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

// Fatal logs fatal level message and exits
func (l *LogrusLogger) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

// Debugf logs debug level formatted message
func (l *LogrusLogger) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

// Infof logs info level formatted message
func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

// Warnf logs warning level formatted message
func (l *LogrusLogger) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

// Errorf logs error level formatted message
func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

// Fatalf logs fatal level formatted message and exits
func (l *LogrusLogger) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

// Default logger instance
var defaultLogger Logger

// InitDefaultLogger initializes the default logger
func InitDefaultLogger(cfg *config.LoggerConfig) {
	defaultLogger = NewLogger(cfg)
}

// GetDefaultLogger returns the default logger instance
func GetDefaultLogger() Logger {
	if defaultLogger == nil {
		// Fallback to basic logger if not initialized
		defaultLogger = NewLogger(&config.LoggerConfig{
			Level:  "info",
			Format: "json",
			Output: "stdout",
		})
	}
	return defaultLogger
}

// Package-level convenience functions using default logger

// Debug logs debug level message using default logger
func Debug(args ...interface{}) {
	GetDefaultLogger().Debug(args...)
}

// Info logs info level message using default logger
func Info(args ...interface{}) {
	GetDefaultLogger().Info(args...)
}

// Warn logs warning level message using default logger
func Warn(args ...interface{}) {
	GetDefaultLogger().Warn(args...)
}

// Error logs error level message using default logger
func Error(args ...interface{}) {
	GetDefaultLogger().Error(args...)
}

// Fatal logs fatal level message using default logger and exits
func Fatal(args ...interface{}) {
	GetDefaultLogger().Fatal(args...)
}

// Debugf logs debug level formatted message using default logger
func Debugf(format string, args ...interface{}) {
	GetDefaultLogger().Debugf(format, args...)
}

// Infof logs info level formatted message using default logger
func Infof(format string, args ...interface{}) {
	GetDefaultLogger().Infof(format, args...)
}

// Warnf logs warning level formatted message using default logger
func Warnf(format string, args ...interface{}) {
	GetDefaultLogger().Warnf(format, args...)
}

// Errorf logs error level formatted message using default logger
func Errorf(format string, args ...interface{}) {
	GetDefaultLogger().Errorf(format, args...)
}

// Fatalf logs fatal level formatted message using default logger and exits
func Fatalf(format string, args ...interface{}) {
	GetDefaultLogger().Fatalf(format, args...)
}
