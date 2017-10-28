package events

type FieldResolver interface {
	Resolve(interface{}) interface{}
}
