package notification

type EmailPayload struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
}

type SmsPayload struct {
	PhoneNumber string `json:"phone_number"`
	Text        string `json:"text"`
}

type SlackPayload struct {
	Channel    string `json:"channel"`
	ReceiverID string `json:"receiver_id"`
	Text       string `json:"text"`
}

type Notification struct {
	Channel   string // "email", "sms", "slack"
	Recipient string // Email address, phone number, or Slack user/channel ID
	Subject   string // Optional (used for emails or slack)
	Body      string // Message content for SMS, email body, or Slack message
}
