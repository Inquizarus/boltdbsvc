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

// Print passes information to logrus Print func
func (ll *LogrusLogger) Print(args ...interface{}) {
	ll.log.Print(args...)
}

// Printf passes information to logrus Printf func
func (ll *LogrusLogger) Printf(format string, args ...interface{}) {
	ll.log.Printf(format, args...)
}

// Println passes information to logrus Println func
func (ll *LogrusLogger) Println(args ...interface{}) {
	ll.log.Println(args...)
}

// Log passes information to logrus Log func
func (ll *LogrusLogger) Log(level Level, args ...interface{}) {
	ll.log.Log(ll.mapToLogrusLevel(level), args...)
}

// Logf passes information to logrus Logf func
func (ll *LogrusLogger) Logf(level Level, format string, args ...interface{}) {
	ll.log.Logf(ll.mapToLogrusLevel(level), format, args...)
}

// Logln passes information to logrus Logln func
func (ll *LogrusLogger) Logln(level Level, args ...interface{}) {
	ll.log.Logln(ll.mapToLogrusLevel(level), args...)
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

// Fatalf ...
func (ll *LogrusLogger) Fatalf(format string, args ...interface{}) {
	ll.Logf(FatalLevel, format, args...)
}

// Fatalln ...
func (ll *LogrusLogger) Fatalln(args ...interface{}) {
	ll.Logln(FatalLevel, args...)
}

// Panic ...
func (ll *LogrusLogger) Panic(args ...interface{}) {
	ll.Log(PanicLevel, args...)
}

// Panicf ...
func (ll *LogrusLogger) Panicf(format string, args ...interface{}) {
	ll.Logf(PanicLevel, format, args...)
}

// Panicln ...
func (ll *LogrusLogger) Panicln(args ...interface{}) {
	ll.Logln(PanicLevel, args...)
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
