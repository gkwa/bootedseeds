package yourpackage

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
