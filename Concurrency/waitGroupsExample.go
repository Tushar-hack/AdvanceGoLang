package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Started ....")
	wg.Add(2)
	go sell()
	go buy()
	wg.Wait()
}

func sell() {
	fmt.Println("Inside sell function....")
	wg.Done()
}

func buy() {
	fmt.Println("Inside buy function....")
	wg.Done()
}
