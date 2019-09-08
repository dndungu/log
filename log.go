package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

type Log struct {
	done     chan struct{}
	exit     func(code int)
	messages chan []byte
	fields   map[string]interface{}
	writer   io.Writer
}

func (l Log) Debug(message string) {
	l.Log(DEBUG, message)
}

func (l Log) Debugf(message string, args ...interface{}) {
	l.Logf(DEBUG, message, args...)
}
func (l Log) Info(message string) {
	l.Log(INFO, message)
}

func (l Log) Infof(message string, args ...interface{}) {
	l.Logf(INFO, message, args...)
}

func (l Log) Warn(message string) {
	l.Log(WARNING, message)
}

func (l Log) Warnf(message string, args ...interface{}) {
	l.Logf(WARNING, message, args...)
}

func (l Log) Error(message string) {
	l.Log(ERROR, message)
}

func (l Log) Errorf(message string, args ...interface{}) {
	l.Logf(ERROR, message, args...)
}

func (l Log) Fatal(message string) {
	l.Log(FATAL, message)

	l.Close()

	l.exit(1)
}

func (l Log) Fatalf(message string, args ...interface{}) {
	l.Logf(FATAL, message, args...)

	l.Close()

	l.exit(1)
}

func (l Log) Close() {
	close(l.messages)
	<-l.done
}

func (l Log) Log(level Level, message string) {
	e := Event{
		Fields:    l.fields,
		Message:   message,
		Level:     level.String(),
		CreatedAt: time.Now().UTC(),
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

func (l Log) Logf(level Level, message string, args ...interface{}) {
	l.Log(level, fmt.Sprintf(message, args...))
}

func New(options ...Option) Interface {
	l := Log{
		done:     make(chan struct{}),
		exit:     os.Exit,
		fields:   make(map[string]interface{}),
		messages: make(chan []byte),
		writer:   os.Stdout,
	}

	for _, option := range options {
		option(&l)
	}

	go func() {
		for m := range l.messages {
			l.writer.Write(append(m, 10))
		}

		l.done <- struct{}{}
	}()

	return &l
}
