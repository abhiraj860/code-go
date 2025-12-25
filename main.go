package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	userId := 10
	respch := make(chan string, 3)

	wg := &sync.WaitGroup{}

	wg.Add(3)
	go fetchUserData(userId, respch, wg)
	go fetchUserRecommendation(userId, respch, wg)
	go fetchUserLike(userId, respch, wg)

	wg.Wait()
	close(respch)

	for resp := range respch {
		fmt.Println(resp)
	}

	fmt.Println(time.Since(now))
}

func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(80 * time.Millisecond)
	respch <- "user data"
}

func fetchUserRecommendation(userID int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(120 * time.Millisecond)
	respch <- "user recommendation"
}

func fetchUserLike(userID int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(50 * time.Millisecond)
	respch <- "user likes"
}
