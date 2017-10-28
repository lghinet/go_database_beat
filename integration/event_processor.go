package integration

type EventProcessor interface {
	Process(event Event) error
}