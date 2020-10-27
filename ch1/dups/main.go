// Dups prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	result := make(map[string]int)

	if len(os.Args) < 2 {
		log.Fatal("Missing command line arguments.")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "d1":
		dup1(os.Stdin, result)
	case "d2":
		dup2(result)
	case "d3":
		dup3(result)
	default:
		log.Fatal("Invalid command line argument.")
		os.Exit(0)
	}

	for line, n := range result {
		fmt.Printf("%d\t%q\n", n, line)
	}
}

func dup1(r io.Reader, counts map[string]int) {
	countLines(r, counts)
	filterLines(counts)
}

func dup2(counts map[string]int) {
	files := os.Args[2:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	filterLines(counts)
}

func dup3(counts map[string]int) {
	for _, f := range os.Args[2:] {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	filterLines(counts)
}

func countLines(r io.Reader, counts map[string]int) {
	input := bufio.NewScanner(r)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func filterLines(counts map[string]int) {
	for line, n := range counts {
		if n <= 1 {
			delete(counts, line)
		}
	}
}
