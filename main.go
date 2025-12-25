package main

import (
	"fmt"
	"sync"
)

func main() {
	value := PackItems(0)
	fmt.Println(value)
}

func PackItems(totalItem int) int {
	const workers = 2
	const itemsPerWorker = 1000

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	itemsPacked := 0
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < itemsPerWorker; j++ {
				mu.Lock()
				itemsPacked = totalItem
				itemsPacked++
				totalItem = itemsPacked
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()
	return totalItem
}
