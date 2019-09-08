package log

type Level uint8

const (
	FATAL = iota
	ERROR
	WARNING
	INFO
	DEBUG
)

func (l Level) String() string {
	return []string{"fatal", "error", "warning", "info", "debug"}[l]
}
