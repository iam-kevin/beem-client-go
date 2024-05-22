package beem

type HTTPRequestSend struct {
	SourceAddr    string      `json:"source_addr"`
	Message       string      `json:"message"`
	Encoding      string      `json:"encoding"`
	Schedule_time string      `json:"schedule_time"`
	Recipients    []recipient `json:"recipients"`
}

type recipient struct {
	Rid    string `json:"recipient_id"`
	Msisdn string `json:"dest_addr"`
}

type HTTPResponseSend struct {
	MessageId string `json:"request_id"`
	Code      string `json:"code"`
}
