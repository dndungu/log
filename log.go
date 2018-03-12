package log

import (
	"encoding/json"
	"fmt"
	"io"
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
	Message string    `json:"message"`
	File    string    `json:"file"`
	Level   string    `json:"level"`
	Time    time.Time `json:"time"`
}

// Log -
type Log struct {
	messages chan []byte
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

// New -
func New(options ...Option) *Log {
	l := Log{messages: make(chan []byte)}
	for _, option := range options {
		option(&l)
	}
	go l.watch()
	return &l
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
}

// watch
func (l Log) watch() {
	for m := range l.messages {
		l.writer.Write(m)
		l.writer.Write([]byte("\n"))
	}
}
