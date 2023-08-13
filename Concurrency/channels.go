package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	fmt.Println("Started....")
	go buy(ch)
	go sell(ch)
	time.Sleep(2 * time.Second)
}

func sell(ch chan string) {
	fmt.Println("Sent data to sell the item")
	ch <- "Furniture"
}

func buy(ch chan string) {
	fmt.Println("Receiving .....")
	val := <-ch
	fmt.Println("Received data", val)
}
