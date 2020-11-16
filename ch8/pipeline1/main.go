package main

import (
	"flag"
	"fmt"
	"time"
)

var n = flag.Int("n", 10, "max number")

func main() {
	flag.Parse()
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < *n; x++ {
			naturals <- x
			time.Sleep(500 * time.Millisecond)
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		// for {

		// 	x, ok := <-naturals
		// 	if !ok {
		// 		break
		// 	}
		// 	squares <- x * x
		// }
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main gorputine)
	// for {
	// 	x, ok := <-squares
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Print(x, " ")
	// }
	for x := range squares {
		fmt.Print(x, " ")
	}
	fmt.Println("Done")
}
