package log

import "time"

type Event struct {
	CreatedAt time.Time              `json:"time"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
	File      string                 `json:"file"`
}