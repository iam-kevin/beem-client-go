package beem_test

import (
	"context"
	"testing"

	"iam-kevin/beem-client"
	"iam-kevin/beem-client/sms"
	"iam-kevin/beem-client/textutils"
)

func TestSendingMessage(t *testing.T) {
	app, _ := beem.New(beem.Options{})

	smsclient := app.Sms(context.TODO(), nil)

	phone, _ := textutils.NewTanzania("0754 311 611")

	sms.SendTextMessage(smsclient, sms.TextMessage{
		SenderId:  "AUTHINFO",
		Message:   "Hey there",
		Recipient: phone,
	})
}
