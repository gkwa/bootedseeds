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
			fmt.Println(*message.MD5OfBody)
		}
	}
}
