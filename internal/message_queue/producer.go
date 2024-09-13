package message_queue

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	kafkaProducer *kafka.Producer
}

func NewProducer(url string) (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": url,
	})
	if err != nil {
		return nil, fmt.Errorf("error crerating Kafka producer: %w", err)
	}

	return &Producer{kafkaProducer: p}, nil
}

func (p *Producer) PublishMessageToTopic(msg any, topic string) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal payload message: %v", err)
	}

	// Produce the message to the specific topic
	err = p.kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	// Wait for message delivery
	p.kafkaProducer.Flush(250)
	log.Printf("Message sent to topic %s: %s\n", topic, string(data))

	return nil
}

func (p *Producer) Close() {
	p.kafkaProducer.Close()
}
