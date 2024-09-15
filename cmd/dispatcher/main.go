package main

import (
	"fmt"
	"log"

	"github.com/JordanRad/notification-system-assignment/internal/message_queue"
	"github.com/JordanRad/notification-system-assignment/internal/service/dispatcher_service"
)

func main() {
	env, err := configFromEnv()
	if err != nil {
		panic(fmt.Errorf("error reading the config: %w", err))
	}

	// Declare 3 different consumers for 3 different channels
	emailConsumer, err := message_queue.NewConsumer(env.Kafka.URL, "email")
	if err != nil {
		panic(fmt.Errorf("cannot open a connection for email topic: %w", err))
	}
	defer emailConsumer.Close()

	smsConsumer, err := message_queue.NewConsumer(env.Kafka.URL, "sms")
	if err != nil {
		panic(fmt.Errorf("cannot open a connection for sms topic: %w", err))
	}
	defer smsConsumer.Close()

	slackConsumer, err := message_queue.NewConsumer(env.Kafka.URL, "slack")
	if err != nil {
		panic(fmt.Errorf("cannot open a connection for email topic: %w", err))
	}
	defer slackConsumer.Close()

	// Declare 3 different dispatchers
	emailDispatcher := dispatcher_service.NewEmailDisptacher()
	smsDispatcher := dispatcher_service.NewSmsDisptacher()
	slackDispatcher := dispatcher_service.NewSlackDisptacher()

	// Decalre 3 working services for each topic
	emailService := dispatcher_service.NewService(emailDispatcher, emailConsumer)
	go listenForTopic(emailService, "Email Notification Service", "email")

	smsService := dispatcher_service.NewService(smsDispatcher, smsConsumer)
	go listenForTopic(smsService, "SMS Notification Service", "sms")

	slackSerivce := dispatcher_service.NewService(slackDispatcher, slackConsumer)
	go listenForTopic(slackSerivce, "Slack Notification Service", "slack")

	// This select unblocks the main thread
	select {}
}

func listenForTopic(service *dispatcher_service.Service, name, topic string) {
	log.Printf("%s started to listen for topic: %s\n", name, topic)
	service.Process()
}
