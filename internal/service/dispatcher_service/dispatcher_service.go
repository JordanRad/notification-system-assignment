package dispatcher_service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/JordanRad/notification-system-assignment/internal/notification"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type NotificationDispatcher interface {
	Dispatch(notification notification.Notification) error
}

type MessageBusConsumer interface {
	ConsumeMessage() (any, error)
	Ack(msg any) error
}

type Service struct {
	dispatcher NotificationDispatcher
	consumer   MessageBusConsumer
}

func NewService(d NotificationDispatcher, c MessageBusConsumer) *Service {
	return &Service{dispatcher: d, consumer: c}
}

func (s *Service) Process() {
	for {
		// Consume the message
		msg, err := s.consumer.ConsumeMessage()
		if err != nil {
			log.Println(fmt.Errorf("error consuming a message: %w", err))
			continue
		}
		// Parse it into Kafka message
		kafkaMsg, ok := msg.(*kafka.Message)
		if !ok || kafkaMsg == nil {
			log.Println(fmt.Errorf("expected *kafka.Message, got %T or nil", msg))
			continue
		}

		// Convert it into unified Notification struct
		topic := *kafkaMsg.TopicPartition.Topic
		notification, err := s.convertKafkaMessageToMessage(topic, kafkaMsg.Value)
		if !ok {
			log.Println(fmt.Errorf("error converting to Notification struct: %w", err))
			continue
		}

		// Dispatch the notification
		err = s.dispatcher.Dispatch(notification)
		if err != nil {
			log.Println(fmt.Errorf("error dispatching %v notification", *kafkaMsg.TopicPartition.Topic))
			continue
		}

		// Acknowledge the dispatched notification
		err = s.consumer.Ack(kafkaMsg)
		if err != nil {
			log.Println(fmt.Errorf("error acknowledging a message: %w", err))
			continue
		}
	}
}

func (s *Service) convertKafkaMessageToMessage(topic string, kafkaMessage []byte) (notification.Notification, error) {
	var n notification.Notification

	switch topic {
	case "email":
		var emailPayload notification.EmailPayload
		if err := json.Unmarshal(kafkaMessage, &emailPayload); err != nil {
			return notification.Notification{}, fmt.Errorf("failed to unmarshal email payload: %w", err)
		}
		n = notification.Notification{
			Channel:   "email",
			Recipient: emailPayload.Email,
			Subject:   emailPayload.Subject,
			Body:      emailPayload.Text,
		}

	case "sms":
		var smsPayload notification.SmsPayload
		if err := json.Unmarshal(kafkaMessage, &smsPayload); err != nil {
			return notification.Notification{}, fmt.Errorf("failed to unmarshal sms payload: %w", err)
		}
		n = notification.Notification{
			Channel:   "sms",
			Recipient: smsPayload.PhoneNumber,
			Body:      smsPayload.Text,
		}

	case "slack":
		var slackPayload notification.SlackPayload
		if err := json.Unmarshal(kafkaMessage, &slackPayload); err != nil {
			return notification.Notification{}, fmt.Errorf("failed to unmarshal slack payload: %w", err)
		}
		n = notification.Notification{
			Channel:   "slack",
			Recipient: slackPayload.ReceiverID,
			Body:      slackPayload.Text,
		}

	default:
		return notification.Notification{}, fmt.Errorf("unknown topic: %s", topic)
	}

	return n, nil
}
