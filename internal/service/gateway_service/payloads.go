package gateway_service

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
