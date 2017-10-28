package integration

type Event interface {
	GetKey() interface{}
	GetData() interface{}
}
