package log

type Level uint8

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func (l Level) String() string {
	return []string{"debug", "info", "warning", "error", "fatal"}[l]
}
