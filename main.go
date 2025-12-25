package main

import "fmt"

func main() {
	balls := make(chan string)
	go throwBall("red", balls)
	go throwBall("green", balls)
	fmt.Println(<-balls, "received")
	fmt.Println(<-balls, "received")
}

func throwBall(color string, balls chan string) {
	fmt.Printf("Throwing string, ball %s ball \n", color)
	balls <- color
}