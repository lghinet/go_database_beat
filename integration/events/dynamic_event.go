package events

import (
	"charisma-beat/integration"
	"log"
)

type eventDefinition struct {
	fields     []eventField
	primaryKey string
}

type DynamicEvent struct {
	data       map[string]interface{}
	primaryKey string
}

func (event DynamicEvent) GetKey() interface{} {
	return event.data[event.primaryKey]
}
func (event DynamicEvent) GetData() interface{} {
	return event.data
}

type dynamicEventMapper struct {
	eventDefinition eventDefinition
	event           DynamicEvent
	fieldResolvers  map[string]FieldResolver
}

func NewDynamicEventMapper(fields []eventField, primaryKey string, fieldResolvers map[string]FieldResolver) *dynamicEventMapper {
	return &dynamicEventMapper{
		eventDefinition: eventDefinition{fields: fields, primaryKey: primaryKey},
		fieldResolvers:  fieldResolvers,
	}
}

func (mapper *dynamicEventMapper) ShouldAddToCollection(row map[string]interface{}) bool {
	if mapper.event.data != nil {
		return mapper.event.data[mapper.eventDefinition.primaryKey] == row[mapper.eventDefinition.primaryKey]
	}
	return false
}

func (mapper *dynamicEventMapper) AddToCollection(row map[string]interface{}) {
	for _, elem := range mapper.eventDefinition.fields {
		if elem.FieldType == "list" {
			newElem := make(map[string]interface{})
			mapper.innerMap(elem.FieldComplexType, row, newElem)
			mapper.event.data[elem.FieldName] = append(mapper.event.data[elem.FieldName].([]map[string]interface{}), newElem)
		}
	}
}

func (mapper *dynamicEventMapper) GetEvent() integration.Event {
	return mapper.event
}

func (mapper *dynamicEventMapper) Map(row map[string]interface{}) {
	mapper.event = DynamicEvent{}
	mapper.event.data = make(map[string]interface{})
	mapper.innerMap(mapper.eventDefinition.fields, row, mapper.event.data)
}

func (mapper *dynamicEventMapper) innerMap(fields []eventField, source map[string]interface{}, dest map[string]interface{}) {
	for _, field := range fields {
		if field.FieldType == "list" {
			dest[field.FieldName] = make([]map[string]interface{}, 0, 100)
			elem := make(map[string]interface{})
			mapper.innerMap(field.FieldComplexType, source, elem)
			dest[field.FieldName] = append(dest[field.FieldName].([]map[string]interface{}), elem)
		} else {
			resolver, ok := mapper.fieldResolvers[field.FieldType]
			if !ok {
				log.Fatalf("error fieldResolver for %s not found", field.FieldType)
			}
			dest[field.FieldName] = resolver.Resolve(source[field.FieldName])
		}
	}
}
