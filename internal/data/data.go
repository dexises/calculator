package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadFile(fileName string) ([]Data, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	var data []Data
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return data, nil
}
