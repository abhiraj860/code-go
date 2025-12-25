package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	value := PackItems(0)
	fmt.Println(value)
}

func PackItems(totalItem int32) int32 {
	const workers = 2
	const itemsPerWorker = 1000

	wg := sync.WaitGroup{}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < itemsPerWorker; j++ {
				atomic.AddInt32(&totalItem, 1)
			}
		}(i)
	}

	wg.Wait()
	return totalItem
}
