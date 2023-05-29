package yourpackage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/sqs"
)

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

func main() {
	// Open the JSON file
	filePath := "data.json"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a JSON decoder
	decoder := json.NewDecoder(file)

	// Decode the JSON data into the Result struct
	var result Result
	err = decoder.Decode(&result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Loop over the extended SQS receive message outputs and print the fields
	for _, output := range result.ExtendedSQSReceiveMessageOutputs {
		fmt.Println("ID:", output.ID)

		var rmo sqs.ReceiveMessageOutput
		err := json.Unmarshal([]byte(output.JSONDef), &rmo)
		if err != nil {
			fmt.Println("Error deserializing message:", err)
			return
		}
		for _, message := range rmo.Messages {
			fmt.Println(message.MD5OfBody)
		}
	}
}
