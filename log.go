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
	exit     func(code int)
	messages chan []byte
	fields   map[string]interface{}
	writer   io.Writer
}

func (l Log) Info(message string) {
	l.Log(INFO, message)
}

func (l Log) Infof(message string, args ...interface{}) {
	l.Log(INFO, fmt.Sprintf(message, args...))
}

func (l Log) Warn(message string) {
	l.Log(WARNING, message)
}

func (l Log) Warnf(message string, args ...interface{}) {
	l.Log(WARNING, fmt.Sprintf(message, args...))
}

func (l Log) Error(message string) {
	l.Log(ERROR, message)
}

func (l Log) Errorf(message string, args ...interface{}) {
	l.Log(ERROR, fmt.Sprintf(message, args...))
}

func (l Log) Fatal(message string) {
	l.Log(FATAL, message)

	close(l.messages)
}

func (l Log) Fatalf(message string, args ...interface{}) {
	l.Log(FATAL, fmt.Sprintf(message, args...))

	close(l.messages)
}

func (l Log) Log(level, message string) {
	e := Event{
		Fields:    l.fields,
		Message:   message,
		Level:     level,
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

func (l Log) watch() {
	var err error
	for m := range l.messages {
		_, err = l.writer.Write(m)

		if err != nil {
			return
		}

		_, err = l.writer.Write([]byte("\n"))

		if err != nil {
			return
		}
	}

	l.exit(1)
}

func New(options ...Option) *Log {
	l := Log{
		exit:     os.Exit,
		fields:   make(map[string]interface{}),
		messages: make(chan []byte),
		writer:   os.Stdout,
	}

	for _, option := range options {
		option(&l)
	}

	go l.watch()

	return &l
}
