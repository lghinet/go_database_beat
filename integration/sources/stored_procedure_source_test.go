package sources

import (
	"charisma-beat/integration/events"
	"log"
	"testing"
)

func TestIterateFakeEventSource(t *testing.T) {

	source := &FakeSource{}

	source.Open()
	defer source.Close()

	var i int

	for _, ok, err := source.Next(); ok; _, ok, err = source.Next() {
		if err != nil {
			t.Fatal("Event failed:", err.Error())
			continue
		}
		//log.Printf("%+v", event.GetData())
		i++
	}

	log.Printf("nr of events %d", i)
}

func BenchmarkIterateSqlEventSource(b *testing.B) {

	source := NewStoredProcedureEventSource(
		"server=APPCOMDEV;user id=sa;password=admin;database=CLARETERP",
		"exec [dbo].[uspL3PaymentScheduleUpdated] ?",
		"PaymentScheduleUpdatedTest",
		&events.PaymentScheduleUpdatedEventMapper{},
		FakeGenerationStore{},
	)

	source.Open()
	defer source.Close()

	var i int

	for _, ok, err := source.Next(); ok; _, ok, err = source.Next() {
		if err != nil {
			b.Fatal("Event failed:", err.Error())
			continue
		}
		//log.Printf("%+v", event.GetData())
		i++
	}

	log.Printf("nr of events %d", i)
}
