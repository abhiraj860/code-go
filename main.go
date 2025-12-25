package main

import (
	"fmt"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	mych := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		mych <- "data"
	} ()

	go func() {
		anotherChannel <- "cow"
	} ()	

	select {
	case msg:= <-mych:
		fmt.Println(msg)
	case anotherChannelmsg := <-anotherChannel:
		fmt.Println(anotherChannelmsg)
	}

}

