package logging

// Logger intarface for application logging
type Logger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Log(Level, ...interface{})
	Logf(Level, string, ...interface{})
	Logln(Level, ...interface{})

	Trace(...interface{})

	Debug(...interface{})

	Info(...interface{})

	Warning(...interface{})

	Error(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}
