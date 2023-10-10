package internal

import (
	"time"
)

type Batch struct {
	Size int
	Wait time.Duration
}

func NewBatch() *Batch {
	return &Batch{
		Size: 1000,
		Wait: 3 * time.Second,
	}
}
