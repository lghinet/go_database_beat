package sources

import (
	"charisma-beat/integration"
	"charisma-beat/integration/events"
	"time"
)

type FakeSource struct {
	count int
}

func (s *FakeSource) Next() (event integration.Event, ok bool, err error) {
	s.count--
	if s.count > 0 {
		ev := events.FakeEvent{ContractId: s.count, ContractSiteId: 0}
		ev.PaymentSchedule = make([]events.FakeEventLine, 0, 100)
		for j := 0; j < 100; j++ {
			ev.PaymentSchedule = append(ev.PaymentSchedule, events.FakeEventLine{
				DueDate:               time.Now(),
				Installment:           "235",
				Interest:              "23423",
				PaymentScheduleLineId: 234,
				RateName:              "sdfsdf",
				RateNumber:            j,
				Principal:             "234253",
				SiteId:                234,
			})
		}

		return ev, true, nil
	}
	return nil, false, nil
}

func (s *FakeSource) Open() error {
	s.count = 1000
	return nil
}

func (s *FakeSource) Close() {
	s.count = 0
}
