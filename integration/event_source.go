package integration

type EventSource interface {
	EventIterator
	Open() error
	Close()
}
