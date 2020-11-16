package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go sqarer(naturals, squares)
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 5; x++ {
		out <- x

	}
	close(out)
}

func sqarer(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Print(v, " ")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println()
}
