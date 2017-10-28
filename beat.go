package main

import (
	"charisma-beat/integration/events"
	"charisma-beat/integration/events/resolvers"
	"charisma-beat/integration/processors"
	"charisma-beat/integration/sources"
	"charisma-beat/scheduler"
	"charisma-beat/scheduler/triggers"
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
)

var (
	dsn                  = flag.String("dsn", "", "Data source name (server=##;user id=##;password=##;database=##) ")
	brokerList           = flag.String("broker-list", "", "Kafka broker list (10.1.3.##:9094) ")
	schedulerInstance    scheduler.Scheduler
	kafkaDefaultProducer sarama.AsyncProducer
	eventsConfiguration  []events.Config
)

func main() {
	flag.Parse()
	validateFlags()

	kafkaDefaultProducer = buildKafkaProducer()

	cleanupDone := make(chan bool)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	cleanup(signalChan, cleanupDone)

	schedulerInstance,_ = CreateScheduler(kafkaDefaultProducer, *dsn)
	schedulerInstance.Start()

	//blocks
	<-cleanupDone
}

func validateFlags() {
	if *dsn == ""{
		log.Fatal("Data source name was not defined")
	}

	if *brokerList == ""{
		log.Fatal("Kafka broker list was not defined")
	}
}

func CreateScheduler(producer sarama.AsyncProducer, dsn string) (scheduler.Scheduler, error) {
	scheduler1 := scheduler.Scheduler{}
	fileStore := sources.NewFileGenerationStore()

	//dynamic events
	for _, config := range eventsConfiguration {
		scheduler1.RegisterJob(scheduler.EventSourceJob{
			sources.NewStoredProcedureEventSource(dsn, config.Sql, config.EventName,
				events.NewDynamicEventMapper(config.Fields, config.PrimaryKey, resolvers.FieldResolvers),
				fileStore),
			processors.KafkaProcessor{
				Producer: producer,
				Topic:    config.Topic,
			},
			&triggers.TimeTrigger{Duration: time.Second * time.Duration(config.Trigger)},
			config.EventName,
		})
	}

	//static events
	/*
		scheduler1.RegisterJob(scheduler.EventSourceJob{
			sources.NewStoredProcedureEventSource(dsn, "exec [dbo].[uspL3ContractCreatedOrUpdated] ?",
				"ContractCreatedOrUpdated",
				&events.ContractCreatedOrUpdatedEventMapper{},
				fileStore),
			processors.KafkaProcessor{
				Producer: producer,
				Topic:    "ch.events.Charisma.Portal.Contracts.CharismaEventProcessor.IntegrationEvents.Contracts.ContractCreatedOrUpdated",
			},
			&triggers.TimeTrigger{Duration: time.Second * 3},
			"ContractCreatedOrUpdated",
		})

		scheduler1.RegisterJob(scheduler.EventSourceJob{
			sources.NewStoredProcedureEventSource(dsn, "exec [dbo].[uspL3PaymentScheduleUpdated] ?",
				"PaymentScheduleUpdated",
				&events.PaymentScheduleUpdatedEventMapper{},
				fileStore),
			processors.KafkaProcessor{
				Producer: producer,
				Topic:    "ch.events.Charisma.Portal.Contracts.CharismaEventProcessor.IntegrationEvents.Contracts.PaymentScheduleUpdated",
			},
			triggers.Every30Seconds(),
			"PaymentScheduleUpdated",
		})

		scheduler1.RegisterJob(scheduler.EventSourceJob{
			sources.NewStoredProcedureEventSource(dsn, "exec [dbo].[uspL3FinancedAssetAddedOrUpdated] ?",
				"FinancedAssetAddedOrUpdated",
				&events.FinancedAssetAddedOrUpdatedEventMapper{},
				fileStore),
			processors.KafkaProcessor{
				Producer: producer,
				Topic:    "ch.events.Charisma.Portal.Contracts.CharismaEventProcessor.IntegrationEvents.Contracts.FinancedAssetAddedOrUpdated",
			},
			triggers.Every30Seconds(),
			"FinancedAssetAddedOrUpdated",
		})

		scheduler1.RegisterJob(scheduler.EventSourceJob{
			sources.NewStoredProcedureEventSource(dsn, "exec [dbo].[uspL3PartnerCreated] ?",
				"PartnerCreated",
				&events.PartnerCreatedEventMapper{},
				fileStore),
			processors.KafkaProcessor{
				Producer: producer,
				Topic:    "ch.events.Charisma.Portal.Partners.CharismaEventProcessor.IntegrationEvents.Partners.PartnerCreated",
			},
			triggers.Every30Seconds(),
			"PartnerCreated",
		})

		scheduler1.RegisterJob(scheduler.EventSourceJob{
			sources.NewStoredProcedureEventSource(dsn, "exec [dbo].[uspL3InsuranceCreated] ?",
				"InsuranceCreated",
				&events.InsuranceCreatedEventMapper{},
				fileStore),
			processors.KafkaProcessor{
				Producer: producer,
				Topic:    "ch.events.Charisma.Portal.Contracts.CharismaEventProcessor.IntegrationEvents.Contracts.InsuranceCreated",
			},
			triggers.Every30Seconds(),
			"InsuranceCreated",
		})
	*/
	return scheduler1, nil
}
func buildKafkaProducer() sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Return.Successes = false
	config.Metadata.Retry.Max = 10
	config.Metadata.Retry.Backoff = time.Second
	config.Producer.Partitioner = sarama.NewHashPartitioner
	//metrics.UseNilMetrics = true

	//https://github.com/Shopify/sarama/issues/959
	sarama.MaxRequestSize = 1000000

	producer, err := sarama.NewAsyncProducer(strings.Split(*brokerList, ","), config)
	if err != nil {
		log.Fatalln("Failed to open Kafka producer:", err.Error())
	}

	return producer
}
func closeProducer(producer sarama.AsyncProducer) {
	if err := producer.Close(); err != nil {
		log.Fatalln("Failed to close Kafka producer cleanly:", err)
	}
}
func cleanup(signalChan chan os.Signal, cleanupDone chan bool) {
	go func() {
		for range signalChan {
			fmt.Println("\nReceived an interrupt, stopping services...\n")
			schedulerInstance.Stop()
			closeProducer(kafkaDefaultProducer)
			cleanupDone <- true
		}
	}()
}
