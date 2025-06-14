package mocks

import (
	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/stretchr/testify/mock"
)

// MockLogger is a mock implementation of logging.Logger
type MockLogger struct {
	mock.Mock
}

// Debug logs debug level message
func (m *MockLogger) Debug(args ...interface{}) {
	m.Called(args...)
}

// Info logs info level message
func (m *MockLogger) Info(args ...interface{}) {
	m.Called(args...)
}

// Warn logs warning level message
func (m *MockLogger) Warn(args ...interface{}) {
	m.Called(args...)
}

// Error logs error level message
func (m *MockLogger) Error(args ...interface{}) {
	m.Called(args...)
}

// Fatal logs fatal level message and exits
func (m *MockLogger) Fatal(args ...interface{}) {
	m.Called(args...)
}

// Debugf logs debug level formatted message
func (m *MockLogger) Debugf(format string, args ...interface{}) {
	m.Called(format, args)
}

// Infof logs info level formatted message
func (m *MockLogger) Infof(format string, args ...interface{}) {
	m.Called(format, args)
}

// Warnf logs warning level formatted message
func (m *MockLogger) Warnf(format string, args ...interface{}) {
	m.Called(format, args)
}

// Errorf logs error level formatted message
func (m *MockLogger) Errorf(format string, args ...interface{}) {
	m.Called(format, args)
}

// Fatalf logs fatal level formatted message and exits
func (m *MockLogger) Fatalf(format string, args ...interface{}) {
	m.Called(format, args)
}

// WithField adds a single field to the logger
func (m *MockLogger) WithField(key string, value interface{}) logging.Logger {
	args := m.Called(key, value)
	return args.Get(0).(logging.Logger)
}

// WithFields adds multiple fields to the logger
func (m *MockLogger) WithFields(fields map[string]interface{}) logging.Logger {
	args := m.Called(fields)
	return args.Get(0).(logging.Logger)
}

// WithError adds an error field to the logger
func (m *MockLogger) WithError(err error) logging.Logger {
	args := m.Called(err)
	return args.Get(0).(logging.Logger)
}
