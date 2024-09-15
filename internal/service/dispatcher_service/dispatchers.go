package dispatcher_service

import (
	"fmt"
	"log"

	"github.com/JordanRad/notification-system-assignment/internal/notification"
)

type EmailDispatcher struct{}

func NewEmailDisptacher() *EmailDispatcher {
	return &EmailDispatcher{}
}

func (e *EmailDispatcher) Dispatch(notification notification.Notification) error {
	if notification.Channel != "email" {
		return fmt.Errorf("invalid channel for EmailDispatcher")
	}
	// Dummy logic
	log.Printf("Sending email to %s with subject %s and body {%s}\n", notification.Recipient, notification.Subject, notification.Body)
	return nil
}

type SmsDispatcher struct{}

func NewSmsDisptacher() *SmsDispatcher {
	return &SmsDispatcher{}
}

func (s *SmsDispatcher) Dispatch(notification notification.Notification) error {
	if notification.Channel != "sms" {
		return fmt.Errorf("invalid channel for SmsDispatcher")
	}
	// Dummy logic
	log.Printf("Sending SMS to %s with body {%s}\n", notification.Recipient, notification.Body)
	return nil
}

type SlackDispatcher struct{}

func NewSlackDisptacher() *SlackDispatcher {
	return &SlackDispatcher{}
}
func (s *SlackDispatcher) Dispatch(notification notification.Notification) error {
	if notification.Channel != "slack" {
		return fmt.Errorf("invalid channel for SlackDispatcher")
	}
	// Dummy logic
	log.Printf("Sending Slack notification to %s with body {%s}\n", notification.Recipient, notification.Body)
	return nil
}
