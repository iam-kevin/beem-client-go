package beem

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"
)

// Instance of the SMS client to handle
// any one request at a time
type SmsInstance struct {
	// handles perform certain http related actions
	Client *http.Client
	// beem instance
	b *App
}

type SmsOptions struct {
	client *http.Client
}

// default insecure client that's seem to work well
// with beem. Performing TLS check seems to fail the API call
// on live environments
func defaultInsecureClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

// Helper to instantiate options for the SMS
func NewSmsOptions() *SmsOptions {
	return &SmsOptions{}
}

// Set's the HTTPClient to the SMS client
func (op *SmsOptions) SetClient(client *http.Client) *SmsOptions {
	op.client = client
	return op
}

// Get the request authorization token expected to be attached
// to the request header as `Authorization: Basic <TOKEN>`
func (op *SmsInstance) GetRequestAuthToken() string {
	return op.b.GetRequestAuthToken()
}

// Get the api url where the request is expected to be sent
func (op *SmsInstance) ApiUrl() string {
	return op.b.ApiUrl()
}

// Create instance that helps manage SMS related actions / interactions
func (b *App) Sms(ctx context.Context, opts *SmsOptions) *SmsInstance {
	o := opts
	if o == nil {
		o = NewSmsOptions()
		o.SetClient(defaultInsecureClient())
	}

	return &SmsInstance{
		Client: o.client,
		b:      b,
	}
}
