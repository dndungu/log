package log

type Interface interface {
	Close()

	Log(level Level, message string)
	Logf(level Level, message string, args ...interface{})

	Debug(message string)
	Debugf(message string, args ...interface{})

	Error(message string)
	Errorf(message string, args ...interface{})

	Info(message string)
	Infof(message string, args ...interface{})

	Fatal(message string)
	Fatalf(message string, args ...interface{})

	Warn(message string)
	Warnf(message string, args ...interface{})
}
