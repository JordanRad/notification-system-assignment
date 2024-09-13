# Notification System Assignment

## Requirements
Notification-sending system requirements:

- The system needs to be able to send notifications via several different channels (email, sms, slack) and be easily extensible to support more channels in the future.

- The system needs to be horizontally scalable.

- The system must guarantee an "at least once" SLA for sending the message.

- The interface for accepting notifications to be sent is an HTTP API, however, that doesn’t mean message queues (e.g Kafka) can’t be used.