package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2 := <-ch2:
		fmt.Println(val2)
	}
}

func goOne(s chan string) {
	s <- "Tushar"
}

func goTwo(s chan string) {
	s <- "Rikku"
}
