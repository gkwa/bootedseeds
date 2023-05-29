package yourpackage

import "gorm.io/gorm"

type ExtendedSQSReceiveMessageOutput struct {
	ID        int     `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
	JSONDef   string  `json:"json_def"`
}

type Result struct {
	ExtendedSQSReceiveMessageOutputs []ExtendedSQSReceiveMessageOutput `json:"extended_sqs_receive_message_outputs"`
}

type ExtendedSqsReceiveMessageOutput struct {
	gorm.Model
	JsonDef string
}

type MessageDetail struct {
	Version    string   `json:"version"`
	ID         string   `json:"id"`
	DetailType string   `json:"detail-type"`
	Source     string   `json:"source"`
	Account    string   `json:"account"`
	Time       string   `json:"time"`
	Region     string   `json:"region"`
	Resources  []string `json:"resources"`
	Detail     Detail   `json:"detail"`
}

type NotificationMessage struct {
	Type             string        `json:"Type"`
	MessageID        string        `json:"MessageId"`
	TopicArn         string        `json:"TopicArn"`
	Message          string        `json:"Message"`
	Timestamp        string        `json:"Timestamp"`
	SignatureVersion string        `json:"SignatureVersion"`
	Signature        string        `json:"Signature"`
	SigningCertURL   string        `json:"SigningCertURL"`
	UnsubscribeURL   string        `json:"UnsubscribeURL"`
	MessageDetail    MessageDetail `json:"-"`
}

type Detail struct {
	InstanceID string `json:"instance-id"`
	State      string `json:"state"`
}
