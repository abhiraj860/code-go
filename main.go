package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	clownChannel := make(chan int, 3)

	clowns := 5

	// sender and receiver logic here
	
	var wg sync.WaitGroup
	go func() {
		defer close(clownChannel)
		for clownId := range clownChannel {
			balloon := fmt.Sprintf("Ballon %d", clownId)
			fmt.Printf("Driver: Drove the car with %s inside \n", balloon)
			time.Sleep(50000 * time.Millisecond)
			fmt.Printf("Driver: Clown finished with %s, the car is ready for more \n", balloon)
		}
	}()
	
	for clown := 1; clown <= clowns; clown++ {
		wg.Add(1)
		go func(clownId int) {
			defer wg.Done()
			balloon := fmt.Sprintf("Balloon %d", clownId)
			fmt.Printf("Clown %d: Hoppend into car with %s \n", clownId, balloon)
			select {
			case clownChannel <- clownId:
				fmt.Printf("Clown %d: Finished with %s\n", clownId,balloon)
			default:
				fmt.Printf("clown %d: Oops the car is filled can fit %s!\n", clownId, balloon)
			}
		}(clown)
	}
	wg.Wait()
	
	fmt.Println("Circus ride is over!!")
}
