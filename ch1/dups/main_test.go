package main

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkDup1(b *testing.B) {
	result := make(map[string]int)
	f, _ := os.Open("sample.txt")

	for n := 0; n < b.N; n++ {
		dup1(f, result)
	}

	for line, n := range result {
		fmt.Printf("%d\t%q\n", n, line)
	}
}
