package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Total Items Packet:", PackItems(0))
}

func PackItems(totalItem int) int {
	const workers = 2
	const itemsPerWorker = 1000

	wg := sync.WaitGroup{}
	itemsPacked := 0
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < itemsPerWorker; j++ {
				itemsPacked = totalItem
				itemsPacked++
				totalItem = itemsPacked
			}
		}(i)
	}
	wg.Wait()
	return totalItem
}
