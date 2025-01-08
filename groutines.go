package main

import (
	"fmt"
	"sync"
)

func hello(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "Hello"
}

func world(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "World"
}

func nums(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- i
}

func groutines() {
	var wg sync.WaitGroup

	chInt := make(chan int, 5)
	chString := make(chan string, 5)
	chString2 := make(chan string, 5)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go hello(chString, &wg)
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go world(chString2, &wg)
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go nums(i, chInt, &wg)
	}

	wg.Wait()
	close(chInt)
	close(chString)
	close(chString2)

	for v := range chString {
		fmt.Println(v)
	}

	for v := range chString2 {
		fmt.Println(v)
	}

	for i := range chInt {
		fmt.Println(i)
	}
}
