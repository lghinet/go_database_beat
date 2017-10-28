package events

import (
	//"time"
	"charisma-beat/integration"
)

type FinancedAssetAddedOrUpdated struct {
	FinancedAssetId     interface{}
	FinancedAssetSiteId interface{}
	ContractId          interface{}
	ContractSiteId      interface{}
	AssetType           interface{}
	Description         interface{}
	Quantity            string
	Amount              string
}

func (event FinancedAssetAddedOrUpdated) GetKey() interface{} {
	return event.FinancedAssetId
}
func (event FinancedAssetAddedOrUpdated) GetData() interface{} {
	return event
}

type FinancedAssetAddedOrUpdatedEventMapper struct {
	event FinancedAssetAddedOrUpdated
}

func (evm *FinancedAssetAddedOrUpdatedEventMapper) ShouldAddToCollection(row map[string]interface{}) bool {
	return false
}

func (evm *FinancedAssetAddedOrUpdatedEventMapper) AddToCollection(row map[string]interface{}) {
}

func (evm *FinancedAssetAddedOrUpdatedEventMapper) GetEvent() integration.Event {
	return evm.event
}

func (evm *FinancedAssetAddedOrUpdatedEventMapper) Map(row map[string]interface{}) {
	evm.event = FinancedAssetAddedOrUpdated{}
	evm.event.FinancedAssetId = row["AssetId"]
	evm.event.FinancedAssetSiteId = row["SiteId"]
	evm.event.ContractId = row["ContractId"]
	evm.event.ContractSiteId = row["ContractSiteId"]
	evm.event.AssetType = row["AssetType"]
	evm.event.Description = row["Description"]

	if row["Quantity"] != nil {
		evm.event.Quantity = string(row["Quantity"].([]byte))
	}
	if row["Amount"] != nil {
		evm.event.Amount = string(row["Amount"].([]byte))
	}
}
