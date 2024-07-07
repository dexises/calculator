package data

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	tests := []struct {
		name           string
		fileName       string
		content        string
		expectedLength int
		expectedSum    int
		shouldFail     bool
	}{
		{
			name:           "ValidFile",
			fileName:       "test_data1.json",
			content:        `[{"a":1,"b":3},{"a":5,"b":-9},{"a":-2,"b":4}]`,
			expectedLength: 3,
			expectedSum:    1 + 3 + 5 - 9 - 2 + 4,
			shouldFail:     false,
		},
		{
			name:           "EmptyFile",
			fileName:       "test_data2.json",
			content:        `[]`,
			expectedLength: 0,
			expectedSum:    0,
			shouldFail:     false,
		},
		{
			name:           "MalformedJSON",
			fileName:       "test_data3.json",
			content:        `{"a":1,"b":3}`,
			expectedLength: 0,
			expectedSum:    0,
			shouldFail:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := os.WriteFile(tt.fileName, []byte(tt.content), 0644); err != nil {
				t.Fatalf("failed to create test file: %v", err)
			}
			defer os.Remove(tt.fileName)

			data, err := ReadFile(tt.fileName)
			if tt.shouldFail {
				if err == nil {
					t.Fatalf("expected an error, but got none")
				}
				return
			} else {
				if err != nil {
					t.Fatalf("failed to read file: %v", err)
				}
			}

			if len(data) != tt.expectedLength {
				t.Fatalf("expected %d elements, got %d", tt.expectedLength, len(data))
			}

			actualSum := 0
			for _, d := range data {
				actualSum += d.A + d.B
			}

			if actualSum != tt.expectedSum {
				t.Fatalf("expected sum %d, got %d", tt.expectedSum, actualSum)
			}
		})
	}
}

func TestReadFileNotExist(t *testing.T) {
	_, err := ReadFile("non_existent_file.json")
	if err == nil {
		t.Fatalf("expected an error, but got none")
	}
}
