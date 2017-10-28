package events

type eventField struct {
	FieldType        string
	FieldName        string
	FieldComplexType []eventField
}
