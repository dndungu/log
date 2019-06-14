package log

import "io"

type Option func(*Log)

func WithField(k string, v interface{}) Option {
	return func(l *Log) {
		l.fields[k] = v
	}
}

func WithWriter(w io.Writer) Option {
	return func(l *Log) {
		l.writer = w
	}
}

func WithExitFunc(f func(code int)) Option {
	return func(l *Log) {
		l.exit = f
	}
}
