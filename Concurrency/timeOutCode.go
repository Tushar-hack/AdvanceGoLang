package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go sendValue(ch1)
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Code Time out")
	}
}

func sendValue(ch chan int) {
	// ch <- 10
}
