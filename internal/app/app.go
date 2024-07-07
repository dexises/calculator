package app

import (
	"github.com/dexises/calculator/internal/data"
	"github.com/dexises/calculator/internal/worker"
	"sync"
)

func Run(numGoroutines int, fileName string) (int, error) {
	resultChan := make(chan int, numGoroutines)
	var wg sync.WaitGroup

	dataItems, err := data.ReadFile(fileName)
	if err != nil {
		return 0, err
	}

	dataSize := len(dataItems)
	batchSize := dataSize / numGoroutines
	remaining := dataSize % numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * batchSize
		end := start + batchSize
		if i == numGoroutines-1 {
			end += remaining
		}

		wg.Add(1)
		go worker.Work(dataItems[start:end], resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)

	totalSum := 0
	for sum := range resultChan {
		totalSum += sum
	}

	return totalSum, nil
}
