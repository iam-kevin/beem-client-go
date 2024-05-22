// currently, sending message via Beem channel
package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/iam-kevin/beem-client/v1"
	"github.com/iam-kevin/beem-client/v1/textutils"
)

type TextMessage struct {
	SenderId  string
	Message   string
	Recipient *textutils.PhoneNumber
}

// Sets the Sender Id to the message to send
func (t *TextMessage) SetSenderId(senderId string) *TextMessage {
	t.SenderId = senderId
	return t
}

// this is the defaulty known sender
// that's available to call the users
//
// See: https://docs.beem.africa/index.html
const defaultSenderId = "INFO"

// Perform action of sending a message defines by the options
func SendTextMessage(s *beem.SmsInstance, opts TextMessage) ([]byte, error) {
	if opts.SenderId == "" {
		opts.SetSenderId(defaultSenderId)
	}

	rc := recipient{
		Rid:    "1",
		Msisdn: opts.Recipient.Value(),
	}

	recipients := make([]recipient, 0, 1)
	recipients = append(recipients, rc)

	// set in friendly sender format
	msg := &HTTPRequestSend{
		SourceAddr:    opts.SenderId,
		Encoding:      "0",
		Message:       opts.Message,
		Recipients:    recipients,
		Schedule_time: "",
	}

	output, _ := json.Marshal(msg)
	jsonBody := bytes.NewReader(output)

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/send", s.ApiUrl()), jsonBody)
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", s.GetRequestAuthToken()))
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}

	data, _ := io.ReadAll(resp.Body)

	return data, nil
}
