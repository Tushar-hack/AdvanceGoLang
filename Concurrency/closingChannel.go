package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	val, ok := <-ch
	close(ch)
	if ok == true {
		fmt.Println(val)
	} else {
		fmt.Println("Channel is Closed!!")
	}
	val, ok = <-ch
	if ok == true {
		fmt.Println(val)
	} else {
		fmt.Println("Channel is Closed!!")
	}
	val, ok = <-ch
	if ok == true {
		fmt.Println(val)
	} else {
		fmt.Println("Channel is Closed!!")
	}
}
