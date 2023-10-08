package loki

import (
	"encoding/json"
)

type Sink struct {
	Pusher *Pusher
}

func NewSink(lp *Pusher) Sink {
	return Sink{
		Pusher: lp,
	}
}

func (s Sink) Sync() error {
	return nil
}

func (s Sink) Close() error {
	return nil
}

func (s Sink) Write(bytes []byte) (count int, err error) {
	entry := new(Entry)
	if ue := json.Unmarshal(bytes, entry); nil != ue {
		err = ue
	} else {
		entry.raw = string(bytes)
		s.Pusher.entries <- *entry
		count = len(bytes)
	}

	return
}
