package main

import (
	"fmt"
	"time"
)

func say(s string ) {
	for i := 1; i <= 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	say("Hell0")
	say("golang")
}
