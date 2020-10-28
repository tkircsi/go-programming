package main

import (
	"fmt"
	"os"
	"strconv"
	"tkircsi/popcount/popcount"
)

func main() {
	// popcount.PrintPC()
	n, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	fmt.Printf("%d popcount is: %d\n", n, popcount.PopCount(n))
	fmt.Printf("%d popcount2 is: %d\n", n, popcount.PopCount2(n))
	fmt.Printf("%d popcount3 is: %d\n", n, popcount.PopCount3(n))
	fmt.Printf("%d popcount4 is: %d\n", n, popcount.PopCount4(n))
}
