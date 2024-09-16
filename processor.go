package main

import (
	"encoding/json"
	"fmt"
)

// Data represents the structure of the JSON data.
// In this example, we assume there's a "name" field we want to extract.
type Data struct {
	Name string `json:"name"`
}

// processFile handles the file reading and JSON processing, then sends the result or error.
func processFile(filePath string, resultChan chan<- string, errChan chan<- error) {
	// Read the file
	dataBytes, err := ReadJSONFile(filePath)
	if err != nil {
		errChan <- err
		return
	}

	// Parse the JSON
	var data Data
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		errChan <- fmt.Errorf("failed to parse JSON from file %s: %w", filePath, err)
		return
	}

	// Send the processed data (name field) to the result channel
	resultChan <- fmt.Sprintf("File: %s, Name: %s", filePath, data.Name)
}
