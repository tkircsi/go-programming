package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int

type WordCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		(*c)++
	}
	if err := scanner.Err(); err != nil {
		return int(*c), err
	}
	return int(*c), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Bogi"
	fmt.Fprintf(&c, "hello %s!", name)
	fmt.Println(c)

	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	var w WordCounter
	fmt.Fprint(&w, input)
	fmt.Println(w)
}
