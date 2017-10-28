package processors

import (
	"charisma-beat/integration"
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"errors"
)

type KafkaProcessor struct {
	Producer sarama.AsyncProducer
	Topic    string
}

func (proc KafkaProcessor) Process(event integration.Event) error {
	jsonData, _ := json.Marshal(event.GetData())
	jsonKeyData, _ := json.Marshal(event.GetKey())

	//log.Println(string(jsonData))

	if jsonData == nil || jsonKeyData == nil {
		return errors.New("error empty event")
	}
	message := &sarama.ProducerMessage{
		Topic:     proc.Topic,
		Key:       sarama.ByteEncoder(jsonKeyData),
		Value:     sarama.ByteEncoder(jsonData),
		Partition: int32(-1)}

	proc.Producer.Input() <- message

	select {
	case msg := <-proc.Producer.Errors():
		log.Println("Failed to produce message:", msg.Err.Error())
		return msg.Err
	default:
	}

	return nil
}

