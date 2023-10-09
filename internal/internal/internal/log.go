package internal

import (
	"time"
)

type Log struct {
	Timestamp time.Time `json:"timestamp"`
	raw       string
}

func (l *Log) Raw() string {
	return l.raw
}
