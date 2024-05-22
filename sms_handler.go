// currently, sending message via Beem channel
package beem

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/iam-kevin/beem-client-go/sms"
)

// this is the defaulty known sender
// that's available to call the users
//
// See: https://docs.beem.africa/index.html
const defaultSenderId = "INFO"

// Perform action of sending a message defines by the options
func (s *SMSInstance) SendTextMessage(opts sms.TextMessage) ([]byte, error) {
	_, cancel := context.WithCancelCause(s.Context())
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

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/send", s.ApiUrl()), bytes.NewReader(output))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", s.GetRequestAuthToken()))
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		cancel(err)
		return nil, err
	}

	data, _ := io.ReadAll(resp.Body)

	return data, nil
}
