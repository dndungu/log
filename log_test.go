package log

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"
)

type Buffer struct {
	data  [][]byte
	mutex sync.Mutex
}

func (b *Buffer) Write(v []byte) (int, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.data = append(b.data, v)
	return len(v), nil
}

func (b *Buffer) Get(k int) []byte {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.data[k]
}

func TestLog(t *testing.T) {
	tests := []struct {
		level   string
		message string
	}{
		{level: "info", message: "1"},
		{level: "warning", message: "2"},
		{level: "error", message: "3"},
		{level: "fatal", message: "4"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			buf := Buffer{}
			l := New(
				WithWriter(&buf),
				WithExitFunc(func(_ int) {}),
				WithField("level", test.level),
			)
			l.Log(test.level, test.message)
			// wait for the l to write to the buffer.
			time.Sleep(10 * time.Millisecond)
			actualEvent := Event{}
			err := json.Unmarshal(buf.Get(0), &actualEvent)
			if err != nil {
				t.Error(err.Error())
			}
			if actualEvent.Level != test.level {
				t.Errorf(
					"Error, expected to find level `%s` found `%s`",
					test.level,
					actualEvent.Level,
				)
			}
			if actualEvent.Message != test.message {
				t.Errorf(
					"Error, expected to find message: `%s` found `%s`",
					test.message,
					actualEvent.Message,
				)
			}
		})
	}
}

func TestInfo(t *testing.T) {
	tests := []struct {
		message string
	}{
		{message: "1"},
		{message: "2"},
		{message: "3"},
		{message: "4"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			buf := Buffer{}
			l := New(WithWriter(&buf), WithExitFunc(func(_ int) {}))
			l.Info(test.message)
			// wait for the l to write to the buffer.
			time.Sleep(10 * time.Millisecond)
			actualEvent := Event{}
			err := json.Unmarshal(buf.Get(0), &actualEvent)
			if err != nil {
				t.Error(err.Error())
			}
			if actualEvent.Message != test.message {
				t.Errorf(
					"Error, expected to find message: `%s` found `%s`",
					test.message,
					actualEvent.Message,
				)
			}
		})
	}
}
func TestInfof(t *testing.T) {
	foo := "foo"
	tests := []struct {
		message string
	}{
		{message: "1 %s"},
		{message: "2 %s"},
		{message: "3 %s"},
		{message: "4 %s"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			buf := Buffer{}
			l := New(WithWriter(&buf), WithExitFunc(func(_ int) {}))
			l.Infof(test.message, foo)
			// wait for the log to write to the buffer.
			time.Sleep(10 * time.Millisecond)
			actualEvent := Event{}
			err := json.Unmarshal(buf.Get(0), &actualEvent)
			if err != nil {
				t.Error(err.Error())
			}
			expectedMessage := fmt.Sprintf(test.message, foo)
			if actualEvent.Message != expectedMessage {
				t.Errorf(
					"Error, expected to find message: `%s` found `%s`",
					expectedMessage,
					actualEvent.Message,
				)
			}
		})
	}
}

func TestWarning(t *testing.T) {
	tests := []struct {
		message string
	}{
		{message: "1"},
		{message: "2"},
		{message: "3"},
		{message: "4"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			buf := Buffer{}
			l := New(WithWriter(&buf), WithExitFunc(func(_ int) {}))
			l.Warning(test.message)
			// wait for the l to write to the buffer.
			time.Sleep(10 * time.Millisecond)
			actualEvent := Event{}
			err := json.Unmarshal(buf.Get(0), &actualEvent)
			if err != nil {
				t.Error(err.Error())
			}
			if actualEvent.Message != test.message {
				t.Errorf(
					"Error, expected to find message: `%s` found `%s`",
					test.message,
					actualEvent.Message,
				)
			}
		})
	}
}

func TestWarningf(t *testing.T) {
	foo := "foo"
	tests := []struct {
		message string
	}{
		{message: "1 %s"},
		{message: "2 %s"},
		{message: "3 %s"},
		{message: "4 %s"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			buf := Buffer{}
			l := New(WithWriter(&buf), WithExitFunc(func(_ int) {}))
			l.Warningf(test.message, foo)
			// wait for the l to write to the buffer.
			time.Sleep(10 * time.Millisecond)
			actualEvent := Event{}
			err := json.Unmarshal(buf.Get(0), &actualEvent)
			if err != nil {
				t.Error(err.Error())
			}
			expectedMessage := fmt.Sprintf(test.message, foo)
			if actualEvent.Message != expectedMessage {
				t.Errorf(
					"Error, expected to find message: `%s` found `%s`",
					expectedMessage,
					actualEvent.Message,
				)
			}
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		message string
	}{
		{message: "1"},
		{message: "2"},
		{message: "3"},
		{message: "4"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			buf := Buffer{}
			l := New(WithWriter(&buf), WithExitFunc(func(_ int) {}))
			l.Error(test.message)
			// wait for the l to write to the buffer.
			time.Sleep(10 * time.Millisecond)
			actualEvent := Event{}
			err := json.Unmarshal(buf.Get(0), &actualEvent)
			if err != nil {
				t.Error(err.Error())
			}
			if actualEvent.Message != test.message {
				t.Errorf(
					"Error, expected to find message: `%s` found `%s`",
					test.message,
					actualEvent.Message,
				)
			}
		})
	}
}

func TestErrorf(t *testing.T) {
	foo := "foo"
	tests := []struct {
		message string
	}{
		{message: "1 %s"},
		{message: "2 %s"},
		{message: "3 %s"},
		{message: "4 %s"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			buf := Buffer{}
			l := New(WithWriter(&buf), WithExitFunc(func(_ int) {}))
			l.Errorf(test.message, foo)
			// wait for the l to write to the buffer.
			time.Sleep(10 * time.Millisecond)
			actualEvent := Event{}
			err := json.Unmarshal(buf.Get(0), &actualEvent)
			if err != nil {
				t.Error(err.Error())
			}
			expectedMessage := fmt.Sprintf(test.message, foo)
			if actualEvent.Message != expectedMessage {
				t.Errorf(
					"Error, expected to find message: `%s` found `%s`",
					expectedMessage,
					actualEvent.Message,
				)
			}
		})
	}
}

func TestFatal(t *testing.T) {
	tests := []struct {
		message string
	}{
		{message: "1"},
		{message: "2"},
		{message: "3"},
		{message: "4"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			buf := Buffer{}
			l := New(WithWriter(&buf), WithExitFunc(func(_ int) {}))
			l.Fatal(test.message)
			// wait for the l to write to the buffer.
			time.Sleep(10 * time.Millisecond)
			actualEvent := Event{}
			err := json.Unmarshal(buf.Get(0), &actualEvent)
			if err != nil {
				t.Error(err.Error())
			}
			if actualEvent.Message != test.message {
				t.Errorf(
					"Error, expected to find message: `%s` found `%s`",
					test.message,
					actualEvent.Message,
				)
			}
		})
	}
}

func TestFatalf(t *testing.T) {
	foo := "foo"
	tests := []struct {
		message string
	}{
		{message: "1 %s"},
		{message: "2 %s"},
		{message: "3 %s"},
		{message: "4 %s"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			buf := Buffer{}
			l := New(WithWriter(&buf), WithExitFunc(func(_ int) {}))
			l.Fatalf(test.message, foo)
			// wait for the l to write to the buffer.
			time.Sleep(10 * time.Millisecond)
			actualEvent := Event{}
			err := json.Unmarshal(buf.Get(0), &actualEvent)
			if err != nil {
				t.Error(err.Error())
			}
			expectedMessage := fmt.Sprintf(test.message, foo)
			if actualEvent.Message != expectedMessage {
				t.Errorf(
					"Error, expected to find message: `%s` found `%s`",
					expectedMessage,
					actualEvent.Message,
				)
			}
		})
	}
}
