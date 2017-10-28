package events

type FakeEventLine struct {
	PaymentScheduleLineId interface{}
	SiteId                interface{}
	RateName              interface{}
	RateNumber            interface{}
	DueDate               interface{}
	Installment           string
	Principal             string
	Interest              string
}

type FakeEvent struct {
	ContractId      interface{}
	ContractSiteId  interface{}
	PaymentSchedule []FakeEventLine
}

func (event FakeEvent) GetKey() interface{} {
	return event.ContractId
}
func (event FakeEvent) GetData() interface{} {
	return event
}
