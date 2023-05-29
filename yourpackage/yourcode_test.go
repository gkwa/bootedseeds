package yourpackage_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/taylormonacelli/bootedseeds/yourpackage"
)

func TestYourFunction(t *testing.T) {
	// Decode the JSON data into the Result struct
	dataPath := filepath.Join("..", "testdata", "data.json")
	// dataPath := filepath.Join(testdata.Dir(), "data.json")

	// Open the file
	file, err := os.Open(dataPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the data
	var result yourpackage.Result
	err = json.Unmarshal(content, &result)
	if err != nil {
		fmt.Println("Error unmarshaling data:", err)
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
			fmt.Println(message.GoString())

			jsonBytes, _ := json.MarshalIndent(message, "", "  ")
			fmt.Println(string(jsonBytes))
			jsonBytes, _ = json.MarshalIndent(message.Attributes, "", "  ")
			fmt.Println("Attributes:", string(jsonBytes))
			fmt.Println("Body:", *message.Body)

			var notification yourpackage.NotificationMessage
			err := json.Unmarshal([]byte(*message.Body), &notification)
			if err != nil {
				fmt.Println("Error deserializing JSON:", err)
				return
			}

			jsonBytes, _ = json.MarshalIndent(notification, "", "  ")
			fmt.Println("Notification:", string(jsonBytes))

			var data map[string]json.RawMessage
			err = json.Unmarshal([]byte(notification.Message), &data)
			if err != nil {
				fmt.Println("Error unmarshaling JSON:", err)
				return
			}
			jsonBytes, _ = json.MarshalIndent(data, "", "  ")
			fmt.Println("Message:", string(jsonBytes))

			fmt.Println("MD5OfBody:", *message.MD5OfBody)

			if message.MD5OfMessageAttributes != nil {
				fmt.Println("MD5OfMessageAttributes:", *message.MD5OfMessageAttributes)
			}

			jsonBytes, _ = json.MarshalIndent(message.MessageAttributes, "", "  ")
			fmt.Println("MessageAttributes:", string(jsonBytes))

			fmt.Println("MessageId:", *message.MessageId)
			fmt.Println("ReceiptHandle:", *message.ReceiptHandle)

		}
	}
}
