package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 3)
	var wg sync.WaitGroup
	wg.Add(2)
	go sell(ch, &wg)
	wg.Wait()
}

func sell(ch chan int, wg *sync.WaitGroup) {
	ch <- 10
	ch <- 11
	ch <- 12
	fmt.Println("Data Sent from seller...")
	go buy(ch, wg)
	wg.Done()
}

func buy(ch chan int, wg *sync.WaitGroup) {
	val := <-ch
	fmt.Println("Received Data by buyer: ", val)
	wg.Done()
}
