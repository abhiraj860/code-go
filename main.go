package main

import (
	"log"
	"runtime"
	"sync"
)

func main() {
	times := 0
	for {
		times++
		value := PackItems(0)
		if value != 2000 {
			log.Fatalf("if should be 2000 but but found %d on execution of counter %d", value, times)
			break
		}
	}
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
				runtime.Gosched()
				totalItem = itemsPacked
			}
		}(i)
	}

	wg.Wait()
	return totalItem
}
