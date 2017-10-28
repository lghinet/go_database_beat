package main

import (
	"charisma-beat/integration/processors"
	"charisma-beat/integration/sources"
	"charisma-beat/scheduler"
	"charisma-beat/scheduler/triggers"
	"os"
	"testing"
	"charisma-beat/integration/events"
	"charisma-beat/integration/events/resolvers"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	*dsn = "server=APPCOMDEV;user id=sa;password=admin;database=CLARETERP"
	*brokerList = "10.1.3.63:9094"
}

func shutdown() {

}

func TestRunOneFakeJobAndSendToKafka(t *testing.T) {
	producer := buildKafkaProducer()

	if producer == nil {
		t.Fatal("no producer")
	}

	schedulerInstance := scheduler.Scheduler{}
	schedulerInstance.RegisterJob(scheduler.EventSourceJob{
		&sources.FakeSource{},
		processors.KafkaProcessor{
			Producer: producer,
			Topic:    "test.go.TestCreateFakeEvent",
		},
		&triggers.OnceTrigger{},
		"TestCreateFakeEvent",
	})

	if len(schedulerInstance.Jobs) == 0 {
		t.Fatalf("No jobs to schedule")
	}
	schedulerInstance.Start()
	schedulerInstance.Stop()
	closeProducer(producer)
}

func BenchmarkRunAllJobs(b *testing.B) {

	producer := buildKafkaProducer()

	if producer == nil {
		b.Fatal("no producer")
	}

	schedulerInstance := scheduler.Scheduler{}
	fileStore := sources.FakeGenerationStore{}

	//dynamic events
	for _, config := range eventsConfiguration {
		schedulerInstance.RegisterJob(scheduler.EventSourceJob{
			sources.NewStoredProcedureEventSource(*dsn, config.Sql, config.EventName,
				events.NewDynamicEventMapper(config.Fields, config.PrimaryKey, resolvers.FieldResolvers),
				fileStore),
			processors.KafkaProcessor{
				Producer: producer,
				Topic:    config.Topic,
			},
			&triggers.OnceTrigger{},
			config.EventName,
		})
	}

	schedulerInstance.Start()
	schedulerInstance.Stop()
	closeProducer(producer)
}