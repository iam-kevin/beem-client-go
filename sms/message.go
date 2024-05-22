package sms

import "github.com/iam-kevin/beem-client-go/textutils"

type TextMessage struct {
	SenderId  string
	Message   string
	Recipient *textutils.PhoneNumber
}

func New(senderId string) *TextMessage {
	return &TextMessage{}
}

// Sets the Sender Id to the message to send
func (t *TextMessage) SetSenderId(senderId string) *TextMessage {
	t.SenderId = senderId
	return t
}

// Sets the message
func (t *TextMessage) SetMessage(message string) *TextMessage {
	t.Message = message
	return t
}
