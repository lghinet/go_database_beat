package events

import (
	//"time"
	"charisma-beat/integration"
)

type InsuranceCreated struct {
	InsuranceId           interface{}
	InsuranceSiteId       interface{}
	ContractId            interface{}
	ContractSiteId        interface{}
	FinancedAssetId       interface{}
	FinancedAssetSiteId   interface{}
	PaymentScheduleId     interface{}
	PaymentScheduleSiteId interface{}
	InsuranceType         interface{}
	InsuranceNumber       interface{}
	InsuranceCompany      interface{}
	StartDate             interface{}
	EndDate               interface{}
	PremiumsCount         interface{}
	Currency              interface{}
	TotalPremium          string
	InsuredValue          string
	Status                interface{}
}

func (event InsuranceCreated) GetKey() interface{} {
	return event.InsuranceId
}
func (event InsuranceCreated) GetData() interface{} {
	return event
}

type InsuranceCreatedEventMapper struct {
	event InsuranceCreated
}

func (evm *InsuranceCreatedEventMapper) ShouldAddToCollection(row map[string]interface{}) bool {
	return false
}

func (evm *InsuranceCreatedEventMapper) AddToCollection(row map[string]interface{}) {
}

func (evm *InsuranceCreatedEventMapper) GetEvent() integration.Event {
	return evm.event
}
func (evm *InsuranceCreatedEventMapper) Map(row map[string]interface{}) {
	evm.event = InsuranceCreated{}
	evm.event.InsuranceId = row["InsuranceId"]
	evm.event.InsuranceSiteId = row["InsuranceSiteId"]
	evm.event.ContractId = row["ContractId"]
	evm.event.ContractSiteId = row["ContractSiteId"]
	evm.event.FinancedAssetId = row["FinancedAssetId"]
	evm.event.FinancedAssetSiteId = row["FinancedAssetSiteId"]
	evm.event.PaymentScheduleId = row["PaymentScheduleId"]
	evm.event.PaymentScheduleSiteId = row["PaymentScheduleSiteId"]
	evm.event.InsuranceType = row["InsuranceType"]
	evm.event.InsuranceNumber = row["InsuranceNumber"]
	evm.event.InsuranceCompany = row["InsuranceCompany"]
	evm.event.StartDate = row["StartDate"]
	evm.event.EndDate = row["EndDate"]
	evm.event.PremiumsCount = row["PremiumsCount"]
	evm.event.Currency = row["Currency"]
	evm.event.Status = row["Status"]

	if row["TotalPremium"] != nil {
		evm.event.TotalPremium = string(row["TotalPremium"].([]byte))
	}
	if row["InsuredValue"] != nil {
		evm.event.InsuredValue = string(row["InsuredValue"].([]byte))
	}
}
