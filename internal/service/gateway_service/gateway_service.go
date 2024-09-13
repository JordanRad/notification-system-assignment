package gateway_service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Notification struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

type MessageQueueProducer interface {
	PublishMessageToTopic(msg any, topic string) error
}

type Service struct {
	messageBusProducer MessageQueueProducer
}

func NewService(mbp MessageQueueProducer) *Service {
	return &Service{
		messageBusProducer: mbp,
	}
}

func (s *Service) HandleNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed. Only POST is supported.", http.StatusMethodNotAllowed)
		return
	}

	messageType := r.URL.Query().Get("message_type")
	if messageType == "" {
		http.Error(w, "Bad Request. Missing message_type parameter.", http.StatusBadRequest)
		return
	}

	// Process the payload based on the message_type
	switch messageType {
	case "email":
		var emailPayload EmailPayload
		err := json.NewDecoder(r.Body).Decode(&emailPayload)
		if err != nil {
			http.Error(w, "Bad Request. Invalid Email JSON.", http.StatusBadRequest)
			return
		}
		log.Printf("Received email notification: %+v", emailPayload)
		err = s.messageBusProducer.PublishMessageToTopic(emailPayload, "email")
		if err != nil {
			log.Println(fmt.Errorf("error sending a message to messsage bus: %w", err))
			http.Error(w, "Conflict. Message cannot be published to the message bus", http.StatusConflict)
			return
		}

	case "sms":
		var smsPayload SmsPayload
		err := json.NewDecoder(r.Body).Decode(&smsPayload)
		if err != nil {
			http.Error(w, "Bad Request. Invalid SMS JSON.", http.StatusBadRequest)
			return
		}
		log.Printf("Received SMS notification: %+v", smsPayload)
		err = s.messageBusProducer.PublishMessageToTopic(smsPayload, "sms")
		if err != nil {
			log.Println(fmt.Errorf("error sending a message to messsage bus: %w", err))
			http.Error(w, "Conflict. Message cannot be published to the message bus", http.StatusConflict)
			return
		}

	case "slack":
		var slackPayload SlackPayload
		err := json.NewDecoder(r.Body).Decode(&slackPayload)
		if err != nil {
			http.Error(w, "Bad Request. Invalid Slack JSON.", http.StatusBadRequest)
			return
		}
		log.Printf("Received Slack notification: %+v", slackPayload)
		err = s.messageBusProducer.PublishMessageToTopic(slackPayload, "slack")
		if err != nil {
			log.Println(fmt.Errorf("error sending a message to messsage bus: %w", err))
			http.Error(w, "Conflict. Message cannot be published to the message bus", http.StatusConflict)
			return
		}
	default:
		http.Error(w, "Bad Request. Unsupported message_type.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("%s notification handled successfully", messageType),
	}
	json.NewEncoder(w).Encode(response)
}
