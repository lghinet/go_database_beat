package integration


type EventMapper interface {
	Map(row map[string]interface{})
	GetEvent() Event
	ShouldAddToCollection(row map[string]interface{}) bool
	AddToCollection(row map[string]interface{})
}
