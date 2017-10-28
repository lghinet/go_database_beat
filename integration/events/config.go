package events

type Config struct {
	Fields     []eventField
	PrimaryKey string
	Topic      string
	Sql        string
	EventName  string
	Trigger    uint64
}
