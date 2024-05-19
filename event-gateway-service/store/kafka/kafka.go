package kafka

import (
	"github.com/Garvit-Jethwani/event-gateway-service/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaStore struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaStore(brokers []string, topic string) (*KafkaStore, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
	})
	if err != nil {
		return nil, err
	}

	return &KafkaStore{
		producer: producer,
		topic:    topic,
	}, nil
}

func (k *KafkaStore) Publish(event models.Event) (string, error) {
	value := event.Payload

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &k.topic, Partition: kafka.PartitionAny},
		Value:          value,
	}

	err := k.producer.Produce(msg, nil)
	if err != nil {
		return "", err
	}

	return msg.TopicPartition.String(), nil
}

func (k *KafkaStore) GetEventStatus(eventID string) (models.EventStatus, error) {
	// Implement the logic to retrieve the event status from Kafka
	return "", nil
}

func (k *KafkaStore) ListEvents(filters ...models.EventFilter) ([]models.EventSummary, error) {
	// Implement the logic to list events from Kafka based on filters
	return nil, nil
}

func (k *KafkaStore) Close() {
	k.producer.Close()
}
