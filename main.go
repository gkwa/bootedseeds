package main

import (
	"fmt"
	"os"
)

func main() {
	// Open the JSON file
	filePath := "testdata/data.json"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
}
