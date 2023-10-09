package internal

type Pusher interface {
	Push(entry *Log)
}
