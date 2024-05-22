package beem

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"
)

// Instance of the SMS client to handle
// any one request at a time
type SMSInstance struct {
	// handles perform certain http related actions
	Client *http.Client
	// beem instance
	b *App

	// context applying throughout the SMS instance
	ctx context.Context
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
func (op *SMSInstance) GetRequestAuthToken() string {
	return op.b.GetRequestAuthToken()
}

// Get the context managing the lifetime of the SMS instance
func (op *SMSInstance) Context() context.Context {
	return op.ctx
}

// Get the api url where the request is expected to be sent
func (op *SMSInstance) ApiUrl() string {
	return op.b.ApiUrl()
}

// Create instance that helps manage SMS related actions / interactions
func (b *App) SMS(ctx context.Context, opts *SmsOptions) *SMSInstance {
	o := opts
	if o == nil {
		o = NewSmsOptions()
		o.SetClient(defaultInsecureClient())
	}

	return &SMSInstance{
		Client: o.client,
		b:      b,
		ctx:    ctx,
	}
}
