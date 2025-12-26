package main

import (
	"fmt"
	"time"
)

func dowork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("DOING_WORK")
		}
	}
}

func main() {
	done := make(chan bool)
	go dowork(done)
	time.Sleep(3 * time.Second)
}
