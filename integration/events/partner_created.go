package events

import (
	//"time"
	"charisma-beat/integration"
)

type PartnerCreated struct {
	PartnerId   interface{}
	PartnerName interface{}
}

func (event PartnerCreated) GetKey() interface{} {
	return event.PartnerId
}
func (event PartnerCreated) GetData() interface{} {
	return event
}

type PartnerCreatedEventMapper struct {
	event PartnerCreated
}

func (evm *PartnerCreatedEventMapper) ShouldAddToCollection(row map[string]interface{}) bool {
	return false
}

func (evm *PartnerCreatedEventMapper) AddToCollection(row map[string]interface{}) {
}

func (evm *PartnerCreatedEventMapper) GetEvent() integration.Event {
	return evm.event
}
func (evm *PartnerCreatedEventMapper) Map(row map[string]interface{}) {
	evm.event = PartnerCreated{}
	evm.event.PartnerId = row["PartnerId"]
	evm.event.PartnerName = row["PartnerName"]
}
