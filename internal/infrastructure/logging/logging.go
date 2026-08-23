// Package implementation for privacy transformation and sensitive-value protection.
package logging

import (
	"encoding/json"
	"log"
	"time"
)

type Logger struct{ base *log.Logger }

func New() *Logger { return &Logger{base: log.Default()} }
func (l *Logger) Event(level, msg string, fields map[string]any) {
	entry := make(map[string]any, len(fields)+3)
	for k, v := range fields {
		entry[k] = v
	}
	entry["level"] = level
	entry["message"] = msg
	entry["ts"] = time.Now().UTC().Format(time.RFC3339Nano)
	b, _ := json.Marshal(entry)
	l.base.Print(string(b))
}
func (l *Logger) Info(msg string, fields map[string]any)  { l.Event("info", msg, fields) }
func (l *Logger) Error(msg string, fields map[string]any) { l.Event("error", msg, fields) }
