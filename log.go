package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

const (
	INFO    = "info"
	WARNING = "warning"
	ERROR   = "error"
	FATAL   = "fatal"
)

// Event -
type Event struct {
	Time    time.Time              `json:"time"`
	Level   string                 `json:"level"`
	Message string                 `json:"message"`
	Fields  map[string]interface{} `json:"fields"`
	File    string                 `json:"file"`
}

// Log -
type Log struct {
	exitFunc func(code int)
	messages chan []byte
	fields   map[string]interface{}
	writer   io.Writer
}

// Option -
type Option func(*Log)

// WithWriter -
func WithWriter(w io.Writer) Option {
	return func(l *Log) {
		l.writer = w
	}
}

// WithExitFunc
func WithExitFunc(f func(code int)) Option {
	return func(l *Log) {
		l.exitFunc = f
	}
}

// New -
func New(options ...Option) *Log {
	l := Log{
		exitFunc: os.Exit,
		fields:   make(map[string]interface{}),
		messages: make(chan []byte, 1024),
		writer:   os.Stdout,
	}
	for _, option := range options {
		option(&l)
	}
	go l.watch()
	return &l
}

//WithFields -
func (l *Log) WithField(k string, v interface{}) *Log {
	l.fields[k] = v
	return l
}

// Info
func (l Log) Info(message string) {
	l.Log(INFO, message)
}

// Infof
func (l Log) Infof(message string, args ...interface{}) {
	l.Log(INFO, fmt.Sprintf(message, args...))
}

// Warning
func (l Log) Warning(message string) {
	l.Log(WARNING, message)
}

// Warningf
func (l Log) Warningf(message string, args ...interface{}) {
	l.Log(WARNING, fmt.Sprintf(message, args...))
}

// Error
func (l Log) Error(message string) {
	l.Log(ERROR, message)
}

// Errorf
func (l Log) Errorf(message string, args ...interface{}) {
	l.Log(ERROR, fmt.Sprintf(message, args...))
}

// Fatal
func (l Log) Fatal(message string) {
	l.Log(FATAL, message)
}

// Fatalf
func (l Log) Fatalf(message string, args ...interface{}) {
	l.Log(FATAL, fmt.Sprintf(message, args...))
}

// Log -
func (l Log) Log(level, message string) {
	e := Event{
		Fields:  l.fields,
		Message: message,
		Level:   level,
		Time:    time.Now().UTC(),
	}
	_, file, line, ok := runtime.Caller(2)
	if ok {
		e.File = fmt.Sprintf("%s:%d", file, line)
	}
	b, err := json.Marshal(e)
	if err == nil {
		l.messages <- b
	}
	if level == FATAL {
		close(l.messages)
	}
}

// watch
func (l Log) watch() {
	for m := range l.messages {
		l.writer.Write(m)
		l.writer.Write([]byte("\n"))
	}
	l.exitFunc(1)
}
