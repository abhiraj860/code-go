package main

import (
	"fmt"
	"sync"
)

func say(s string, wg * sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	go say("Hello", &wg)
	go say("World", &wg)

	wg.Wait()
	fmt.Println("Finish")
}