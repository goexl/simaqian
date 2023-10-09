package internal

import (
	"encoding/json"
)

type Sink struct {
	pusher Pusher
}

func NewSink(pusher Pusher) *Sink {
	return &Sink{
		pusher: pusher,
	}
}

func (s *Sink) Sync() error {
	return nil
}

func (s *Sink) Close() error {
	return nil
}

func (s *Sink) Write(bytes []byte) (count int, err error) {
	entry := new(Log)
	if ue := json.Unmarshal(bytes, entry); nil != ue {
		err = ue
	} else {
		entry.raw = string(bytes)
		s.pusher.Push(entry)
		count = len(bytes)
	}

	return
}
