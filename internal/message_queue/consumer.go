package message_queue

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	kafkaConsumer *kafka.Consumer
}

func NewConsumer(url, topic string) (*Consumer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers":  url,
		"group.id":           fmt.Sprintf("consumer-group-%s", topic),
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	}

	c, err := kafka.NewConsumer(config)
	if err != nil || c == nil {
		return nil, fmt.Errorf("failed to create Kafka consumer: %w", err)
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to topic: %w", err)
	}

	return &Consumer{kafkaConsumer: c}, nil
}

// ConsumeMessages listens for messages and returns the message and topic.
// The last returned argument is error
func (c *Consumer) ConsumeMessage() (any, error) {
	msg, err := c.kafkaConsumer.ReadMessage(-1) // -1 waits indefinitely for a message
	if err != nil {
		log.Printf("Error reading message: %v\n", err)
		return nil, err
	}

	log.Printf("Message received from topic %s: %s\n", *msg.TopicPartition.Topic, string(msg.Value))
	return msg, nil
}

// Ack acknowlegdes that the message has been successfully
// process and returns an error (if any occurs)
func (c *Consumer) Ack(msg any) error {
	// Type assertion to check if the message is of type *kafka.Message
	kafkaMsg, ok := msg.(*kafka.Message)
	if !ok {
		return fmt.Errorf("expected *kafka.Message, got %T", msg)
	}

	_, err := c.kafkaConsumer.CommitMessage(kafkaMsg)
	return err
}

func (c *Consumer) Close() {
	c.kafkaConsumer.Close()
}
