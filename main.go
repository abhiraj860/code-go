package main

import (
	"fmt"
	"time"
)

type Message struct {
	from     string
	payloads string
}

type Server struct {
	msgch chan Message
}

func (s *Server) StartAndListen() {
	for {
		select {
		case msg := <-s.msgch:
			// block here until someone is sending a message to a channel
			fmt.Printf("received message for %s payloads %s\n", msg.from, msg.payloads)
		default:

		}
	}
}

func sendMessageToServer(msgch chan Message, payloadInp string) {
	msg := Message{
		from:     "YoBuyDm",
		payloads: payloadInp,
	}
	msgch <- msg
}

func main() {
	s := Server{
		msgch: make(chan Message),
	}
	go s.StartAndListen()
	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(s.msgch, "Hello Sailor")
	}()
	select {}
}

// func main() {
// 	now := time.Now()
// 	userId := 10
// 	respch := make(chan string, 3)

// 	wg := &sync.WaitGroup{}

// 	wg.Add(3)
// 	go fetchUserData(userId, respch, wg)
// 	go fetchUserRecommendation(userId, respch, wg)
// 	go fetchUserLike(userId, respch, wg)

// 	wg.Wait()
// 	close(respch)

// 	for resp := range respch {
// 		fmt.Println(resp)
// 	}

// 	fmt.Println(time.Since(now))
// }

// func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	time.Sleep(80 * time.Millisecond)
// 	respch <- "user data"
// }

// func fetchUserRecommendation(userID int, respch chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	time.Sleep(120 * time.Millisecond)
// 	respch <- "user recommendation"
// }

// func fetchUserLike(userID int, respch chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	time.Sleep(50 * time.Millisecond)
// 	respch <- "user likes"
// }
