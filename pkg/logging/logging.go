package logging

// Logger intarface for application logging
type Logger interface {
	Log(Level, ...interface{})
	Trace(...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warning(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	Panic(...interface{})
}
