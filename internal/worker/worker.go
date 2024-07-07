package worker

import (
	"github.com/dexises/calculator/internal/data"
	"sync"
)

func Work(dataItems []data.Data, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, d := range dataItems {
		sum += d.A + d.B
	}
	resultChan <- sum
}
