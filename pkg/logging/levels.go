package logging

// Level is the base type for logging levels
type Level uint32

const (
	// TraceLevel for trace level logging
	TraceLevel = iota
	// DebugLevel for debug level logging
	DebugLevel
	// InfoLevel for info level logging
	InfoLevel
	// WarningLevel for warning level logging
	WarningLevel
	// ErrorLevel for error level logging
	ErrorLevel
	// FatalLevel for fatal level logging
	FatalLevel
	// PanicLevel for panic level logging
	PanicLevel
)
