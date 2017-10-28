package resolvers

import "charisma-beat/integration/events"

var FieldResolvers map[string]events.FieldResolver

func init() {
	FieldResolvers = map[string]events.FieldResolver{
		"date":    dateResolver{},
		"string":  stringResolver{},
		"decimal": decimalResolver{},
		"int":     intResolver{},
	}
}
