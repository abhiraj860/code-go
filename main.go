package main

import (
	"fmt"
)

func sliceToChannel(inp []int) chan int {
	ch := make(chan int)
	go func() {
		for _, v := range inp {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func sq(data chan int) chan int {
	p := make(chan int)
	go func() {
		for itr := range data {
			p <- (itr * itr)
		}
		close(p)
	} ()
	return p
}

func main() {
	inp := []int{2, 3, 4, 7, 1}

	dataChannel := sliceToChannel(inp)
	finalChannel := sq(dataChannel)

	for itr := range finalChannel {
		fmt.Println(itr)
	}
}
