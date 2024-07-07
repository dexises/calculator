package app

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name          string
		fileName      string
		content       string
		numGoroutines int
		expectedSum   int
		shouldFail    bool
	}{
		{
			name:          "ValidFileMultipleGoroutines",
			fileName:      "test_data1.json",
			content:       `[{"a":1,"b":3},{"a":5,"b":-9},{"a":-2,"b":4}]`,
			numGoroutines: 2,
			expectedSum:   1 + 3 + 5 - 9 - 2 + 4,
			shouldFail:    false,
		},
		{
			name:          "EmptyFile",
			fileName:      "test_data2.json",
			content:       `[]`,
			numGoroutines: 2,
			expectedSum:   0,
			shouldFail:    false,
		},
		{
			name:          "MalformedJSON",
			fileName:      "test_data3.json",
			content:       `{"a":1,"b":3}`,
			numGoroutines: 2,
			expectedSum:   0,
			shouldFail:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем временный JSON файл для тестирования
			if err := os.WriteFile(tt.fileName, []byte(tt.content), 0644); err != nil {
				t.Fatalf("failed to create test file: %v", err)
			}
			defer os.Remove(tt.fileName) // Удаляем файл после теста

			result, err := Run(tt.numGoroutines, tt.fileName)
			if tt.shouldFail {
				if err == nil {
					t.Fatalf("expected an error, but got none")
				}
				return
			} else {
				if err != nil {
					t.Fatalf("Run() returned an error: %v", err)
				}
			}

			if result != tt.expectedSum {
				t.Fatalf("expected sum %d, got %d", tt.expectedSum, result)
			}
		})
	}
}
