package events

import (
	"github.com/Garvit-Jethwani/notification-service/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type EventConsumer interface {
	Start() error
	Close() error
	Consume() (<-chan models.Event, error)
}

type KafkaConsumer struct {
	consumer  *kafka.Consumer
	eventChan chan models.Event
}

func NewKafkaConsumer(brokers []string, groupID, topics string) (*KafkaConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	consumer.SubscribeTopics([]string{topics}, nil)

	return &KafkaConsumer{
		consumer:  consumer,
		eventChan: make(chan models.Event),
	}, nil
}

func (k *KafkaConsumer) Start() error {
	go k.consumeEvents()
	return nil
}

func (k *KafkaConsumer) Close() error {
	return k.consumer.Close()
}

func (k *KafkaConsumer) Consume() (<-chan models.Event, error) {
	return k.eventChan, nil
}

func (k *KafkaConsumer) consumeEvents() {
	for {
		ev, err := k.consumer.ReadMessage(-1)
		if err != nil {
			// Handle error
			continue
		}

		event := models.Event{
			ID:        string(ev.Key),
			Type:      *ev.TopicPartition.Topic,
			Payload:   ev.Value,
			Timestamp: ev.Timestamp,
		}

		k.eventChan <- event
	}
}
