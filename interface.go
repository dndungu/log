package log

type Interface interface {
	Error(message string)
	Errorf(message string, args ...interface{})
	Info(message string)
	Infof(message string, args ...interface{})
	Log(level, message string)
	Fatal(message string)
	Fatalf(message string, args ...interface{})
	Warn(message string)
	Warnf(message string, args ...interface{})
}
