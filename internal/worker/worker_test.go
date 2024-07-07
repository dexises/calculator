package worker

import (
	"github.com/dexises/calculator/internal/data"
	"sync"
	"testing"
)

func TestWorker(t *testing.T) {
	tests := []struct {
		name        string
		inputData   []data.Data
		expectedSum int
	}{
		{
			name: "PositiveAndNegativeValues",
			inputData: []data.Data{
				{A: 1, B: 3},
				{A: 5, B: -9},
				{A: -2, B: 4},
			},
			expectedSum: 1 + 3 + 5 - 9 - 2 + 4,
		},
		{
			name: "AllZeroValues",
			inputData: []data.Data{
				{A: 0, B: 0},
				{A: 0, B: 0},
			},
			expectedSum: 0,
		},
		{
			name:        "EmptyData",
			inputData:   []data.Data{},
			expectedSum: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultChan := make(chan int, 1)
			var wg sync.WaitGroup

			wg.Add(1)
			go Work(tt.inputData, resultChan, &wg)

			wg.Wait()
			close(resultChan)

			actualSum := 0
			for sum := range resultChan {
				actualSum += sum
			}

			if actualSum != tt.expectedSum {
				t.Fatalf("expected sum %d, got %d", tt.expectedSum, actualSum)
			}
		})
	}
}
