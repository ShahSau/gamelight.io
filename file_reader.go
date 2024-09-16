package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ReadJSONFile reads the content of a JSON file and returns the byte data or an error.
func ReadJSONFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	// Read the file's content
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	return byteValue, nil
}
