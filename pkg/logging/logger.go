package logging

import (
	"github.com/sirupsen/logrus"
)

// LogrusLogger wraps logrus into an application logger
type LogrusLogger struct {
	log *logrus.Logger
}

func (ll *LogrusLogger) mapToLogrusLevel(level Level) logrus.Level {
	switch level {
	case TraceLevel:
		return logrus.TraceLevel
	case DebugLevel:
		return logrus.DebugLevel
	case InfoLevel:
		return logrus.InfoLevel
	case WarningLevel:
		return logrus.WarnLevel
	case ErrorLevel:
		return logrus.ErrorLevel
	case FatalLevel:
		return logrus.FatalLevel
	case PanicLevel:
		return logrus.PanicLevel
	}
	return logrus.Level(100) // Unknown level
}

// Log passes information to logrus Log func
func (ll *LogrusLogger) Log(level Level, args ...interface{}) {
	ll.log.Log(ll.mapToLogrusLevel(level), args...)
}

// Trace ...
func (ll *LogrusLogger) Trace(args ...interface{}) {
	ll.Log(TraceLevel, args...)
}

// Debug ...
func (ll *LogrusLogger) Debug(args ...interface{}) {
	ll.Log(DebugLevel, args...)
}

// Info ...
func (ll *LogrusLogger) Info(args ...interface{}) {
	ll.Log(InfoLevel, args...)
}

// Warning ...
func (ll *LogrusLogger) Warning(args ...interface{}) {
	ll.Log(WarningLevel, args...)
}

// Error ...
func (ll *LogrusLogger) Error(args ...interface{}) {
	ll.Log(ErrorLevel, args...)
}

// Fatal ...
func (ll *LogrusLogger) Fatal(args ...interface{}) {
	ll.Log(FatalLevel, args...)
}

// Panic ...
func (ll *LogrusLogger) Panic(args ...interface{}) {
	ll.Log(PanicLevel, args...)
}

// NewLogrusLogger creates a new LogrusLogger with passed or standard instance of logrus
func NewLogrusLogger(log *logrus.Logger) Logger {
	if nil == log {
		log = logrus.New()
		log.SetFormatter(&logrus.JSONFormatter{})
		log.SetLevel(logrus.InfoLevel)
	}
	return &LogrusLogger{
		log: log,
	}
}
