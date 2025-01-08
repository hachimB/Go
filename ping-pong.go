package main

import (
	"fmt"
	"sync"
)

func ping(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- "ping"
		fmt.Println("Message sent: ping")
		res := <-ch
		fmt.Println("Message received: ", res)
	}
}

func pong(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		res := <-ch
		fmt.Println("Message received: ", res)
		ch <- "pong"
		fmt.Println("Message send: pong")
	}
}

func pingPong() {
	var wg sync.WaitGroup

	ch := make(chan string)

	wg.Add(2)
	go ping(ch, &wg)
	go pong(ch, &wg)

	wg.Wait()
	close(ch)

}
