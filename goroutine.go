package main

import (
	"fmt"
	"sync"
	"time"
)

func run() {
	time.Sleep(3 * time.Second)
	fmt.Println("run")
}

func run2() {
	time.Sleep(2 * time.Second)
	fmt.Println("run2")
}

func say(message string) {
	for _, m := range message {
		fmt.Print(string(m))
		time.Sleep(1 * time.Second)
	}
}

func say2(message string) {
	for _, m := range message {
		fmt.Print(string(m))
		time.Sleep(3 * time.Second)
	}
}

func addNums(x, y int, ch chan int) {
	ch <- x + y
}

func channel(message string, ch chan string) {
	ch <- message
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("The worker %d starting\n", id)
	fmt.Printf("The worker %d is done\n", id)
}

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func reducer(ch chan int) {
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
}

func goroutine() {

	// WaitGroup
	ch := make(chan int)

	go producer(ch)

	reducer(ch)

	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go worker(12, &wg)
	// }

	//wg.Wait()
	//close(ch)

	//**************************************

	// go say2("Hello")
	// say("Hello world")

	//time
	// go run()
	// go run2()
	// time.Sleep(4 * time.Second)
	// fmt.Println("End")

	//Channel
	// ch := make(chan int)
	// go addNums(2, 3, ch)
	// res := <-ch
	// fmt.Println(res)

	// chann := make(chan string)
	// go channel("Hello world", chann)

	// go channel("Hello everyone", chann)
	// go channel("Hello you!", chann)

	// message := <-chann
	// message2 := <-chann
	// message3 := <-chann
	// fmt.Println(message)
	// fmt.Println(message2)
	// fmt.Println(message3)

	// chann2 := make(chan int, 9)
	// chann2 <- 2

	// chann2 <- 209
	// fmt.Println(<-chann2)
	// fmt.Println(<-chann2)
	// fmt.Printf("%T\n", chann2)

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case msg1 := <-chann:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-chann2:
	// 		fmt.Println(msg2)
	// 	}
	//}

}
