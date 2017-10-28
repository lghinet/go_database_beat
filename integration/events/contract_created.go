package events

import (
	//"time"
	"charisma-beat/integration"
)

type ContractCreatedOrUpdated struct {
	ContractId        interface{}
	SiteId            interface{}
	PartnerId         interface{}
	PartnerName       interface{}
	ContractType      interface{}
	ContractNumber    interface{}
	SigningDate       interface{}
	ActivationDate    interface{}
	EffectiveDate     interface{}
	TerminationDate   interface{}
	Address           interface{}
	ResponsiblePerson interface{}
	DurationMonths    interface{}
	ContractCurrency  interface{}
	ContractValue     string
	DownPayment       string
	FinancedAmount    string
	ResidualValue     string
	Status            interface{}
}

func (event ContractCreatedOrUpdated) GetKey() interface{} {
	return event.ContractId
}
func (event ContractCreatedOrUpdated) GetData() interface{} {
	return event
}

type ContractCreatedOrUpdatedEventMapper struct {
	event ContractCreatedOrUpdated
}

func (evm *ContractCreatedOrUpdatedEventMapper) ShouldAddToCollection(row map[string]interface{}) bool {
	return false
}

func (evm *ContractCreatedOrUpdatedEventMapper) AddToCollection(row map[string]interface{}) {
}

func (evm *ContractCreatedOrUpdatedEventMapper) GetEvent() integration.Event {
	return evm.event
}

func (evm *ContractCreatedOrUpdatedEventMapper) Map(row map[string]interface{}) {
	evm.event = ContractCreatedOrUpdated{}
	evm.event.ContractId = row["ContractId"]
	evm.event.SiteId = row["SiteId"]
	evm.event.PartnerId = row["PartnerId"]
	evm.event.PartnerName = row["PartnerName"]
	evm.event.ContractType = row["ContractType"]
	evm.event.ContractNumber = row["ContractNumber"]
	evm.event.SigningDate = row["SigningDate"]
	evm.event.ActivationDate = row["ActivationDate"]
	evm.event.EffectiveDate = row["EffectiveDate"]
	evm.event.TerminationDate = row["TerminationDate"]
	evm.event.Address = row["Address"]
	evm.event.ResponsiblePerson = row["ResponsiblePerson"]
	evm.event.DurationMonths = row["DurationMonths"]
	evm.event.ContractCurrency = row["ContractCurrency"]
	if row["ContractValue"] != nil {
		evm.event.ContractValue = string(row["ContractValue"].([]byte))
	}
	if row["DownPayment"] != nil {
		evm.event.DownPayment = string(row["DownPayment"].([]byte))
	}
	if row["FinancedAmount"] != nil {
		evm.event.FinancedAmount = string(row["FinancedAmount"].([]byte))
	}
	if row["ResidualValue"] != nil {
		evm.event.ResidualValue = string(row["ResidualValue"].([]byte))
	}
	evm.event.Status = row["Status"]
}
