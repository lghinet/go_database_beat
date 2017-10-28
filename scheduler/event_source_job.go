package scheduler

import (
	"charisma-beat/integration"
	"log"
)

type EventSourceJob struct {
	Source    integration.EventSource
	Processor integration.EventProcessor
	Trigger   Trigger
	Name      string
}

func (job EventSourceJob) Run() {
	job.Source.Open()
	defer job.Source.Close()

	for event, ok, err := job.Source.Next(); ok; event, ok, err = job.Source.Next() {
		if err != nil {
			log.Println("Event failed:", err.Error())
			continue
		}

		err = job.Processor.Process(event)
		if err != nil {
			log.Println("Process failed:", err.Error())
			continue
		}
		//log.Printf("%+v", event.GetData())
	}
}

func (job EventSourceJob) GetTrigger() Trigger {
	return job.Trigger
}

func (job EventSourceJob) GetName() string {
	return job.Name
}
