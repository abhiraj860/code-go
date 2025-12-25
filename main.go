package main

import (
	"fmt"
	"sync"
)

func throwBall(color string, ch chan string) {
	fmt.Printf("Sending %s in the channel\n", color)
	ch <- color
}

func main() {
	balls := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		throwBall("red", balls)
	}()
	go func() {
		defer wg.Done()
		throwBall("Green", balls)
	}()
	go func() {
		wg.Wait()
		close(balls)
	} ()
	
	for color := range balls {
		fmt.Println(color, "received")
	}
}
