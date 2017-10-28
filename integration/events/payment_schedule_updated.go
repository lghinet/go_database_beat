package events

import (
	//"time"
	"charisma-beat/integration"
)

type paymentScheduleLine struct {
	PaymentScheduleLineId interface{}
	SiteId                interface{}
	RateName              interface{}
	RateNumber            interface{}
	DueDate               interface{}
	Installment           string
	Principal             string
	Interest              string
}

type PaymentScheduleUpdated struct {
	ContractId      interface{}
	ContractSiteId  interface{}
	PaymentSchedule []paymentScheduleLine
}

func (event PaymentScheduleUpdated) GetKey() interface{} {
	return event.ContractId
}
func (event PaymentScheduleUpdated) GetData() interface{} {
	return event
}

type PaymentScheduleUpdatedEventMapper struct {
	event PaymentScheduleUpdated
}

func (evm *PaymentScheduleUpdatedEventMapper) ShouldAddToCollection(row map[string]interface{}) bool {
	return evm.event.ContractId == row["ContractId"]
}
func (evm *PaymentScheduleUpdatedEventMapper) AddToCollection(row map[string]interface{}) {
	scheduleLine := paymentScheduleLine{}
	scheduleLine.PaymentScheduleLineId = row["PaymentScheduleLineId"]
	scheduleLine.SiteId = row["PaymentScheduleLineSiteId"]
	scheduleLine.RateName = row["RateName"]
	scheduleLine.RateNumber = row["RateNumber"]
	scheduleLine.DueDate = row["DueDate"]
	if row["Installment"] != nil {
		scheduleLine.Installment = string(row["Installment"].([]byte))
	}
	if row["Principal"] != nil {
		scheduleLine.Principal = string(row["Principal"].([]byte))
	}
	if row["Interest"] != nil {
		scheduleLine.Interest = string(row["Interest"].([]byte))
	}

	evm.event.PaymentSchedule = append(evm.event.PaymentSchedule, scheduleLine)
}

func (evm *PaymentScheduleUpdatedEventMapper) GetEvent() integration.Event {
	return evm.event
}

func (evm *PaymentScheduleUpdatedEventMapper) Map(row map[string]interface{}) {
	evm.event = PaymentScheduleUpdated{}
	evm.event.ContractId = row["ContractId"]
	evm.event.ContractSiteId = row["SiteId"]

	scheduleLine := paymentScheduleLine{}
	scheduleLine.PaymentScheduleLineId = row["PaymentScheduleLineId"]
	scheduleLine.SiteId = row["PaymentScheduleLineSiteId"]
	scheduleLine.RateName = row["RateName"]
	scheduleLine.RateNumber = row["RateNumber"]
	scheduleLine.DueDate = row["DueDate"]
	if row["Installment"] != nil {
		scheduleLine.Installment = string(row["Installment"].([]byte))
	}
	if row["Principal"] != nil {
		scheduleLine.Principal = string(row["Principal"].([]byte))
	}
	if row["Interest"] != nil {
		scheduleLine.Interest = string(row["Interest"].([]byte))
	}

	evm.event.PaymentSchedule = make([]paymentScheduleLine, 0, 100)
	evm.event.PaymentSchedule = append(evm.event.PaymentSchedule, scheduleLine)
}
