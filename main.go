package main

import (
	"sync"
	"fmt"
)

func main() {
	signalChannel := make(chan bool)
	wg := sync.WaitGroup{}
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1 is waiting for signal...")
		<-signalChannel
		fmt.Println("Gorouting 1 received the signal and is now doing something")
	} ()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2 is about to send a signal...")
		signalChannel <- true
		fmt.Println("Gourouting 2 send the signal.")
	} ()
	wg.Wait()
	fmt.Println("Both goroutines have finished")
}