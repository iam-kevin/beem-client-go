// currently, sending message via Beem channel
package beem

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/iam-kevin/beem-client-go/sms"
)

// this is the defaulty known sender
// that's available to call the users
//
// See: https://docs.beem.africa/index.html
const defaultSenderId = "INFO"

type SMSResponse struct {
	res *SMSHTTPResponse
}

func (s *SMSResponse) HttpResponse() *http.Response {
	return s.res.Raw
}

type SMSHTTPResponse struct {
	Raw *http.Response
}

func (res *SMSHTTPResponse) Body() ([]byte, error) {
	return io.ReadAll(res.Raw.Body)
}

// Perform action of sending a message defines by the options
func (s *SMSInstance) SendTextMessage(opts sms.TextMessage) (*SMSResponse, error) {
	ctx := s.Context()
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

	sendUrl := fmt.Sprintf("%s/v1/send", s.ApiUrl())
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, sendUrl, bytes.NewReader(output))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", s.GetRequestAuthToken()))
	req.Header.Set("Content-Type", "application/json")

	slog.Debug("making a call to url to send message", "request", req.URL, "headers", req.Header.Clone())

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return &SMSResponse{
		res: &SMSHTTPResponse{
			Raw: resp,
		},
	}, nil
}
