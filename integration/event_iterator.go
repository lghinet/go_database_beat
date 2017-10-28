package integration


type EventIterator interface {
	Next() (event Event, ok bool, err error)
}
