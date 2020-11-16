package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var r = flag.Int("r", 1, "readN function")

func main() {
	flag.Parse()
	switch *r {
	case 1:
		read1()
	case 2:
		read2()
	case 3:
		read3()
	case 4:
		read4()
	case 5:
		read5()
	default:
		read1()
	}
}

// bufio.NewReader.ReadLine
func read1() {
	fmt.Println("bufio.NewReader.ReadLine")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text or \"exit\" to exit: ")
		line, isPrefix, err := reader.ReadLine()
		_ = isPrefix
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(line)
		fmt.Println(string(line))

		if string(line) == "exit" {
			break
		}
	}
	fmt.Println("Viszlát!")
}

// bufio.NewReader.ReadBytes
func read2() {
	fmt.Println("bufio.NewReader.ReadBytes")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text or \"exit\" to exit: ")
		// read until the delimiter
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(line)
		// line contains the delimiter too
		line = line[:len(line)-1]
		fmt.Println(string(line))
		if string(line) == "exit" {
			break
		}
	}
	fmt.Println("Viszlát!")
}

// bufio.NewReader.ReadString
func read3() {
	fmt.Println("bufio.NewReader.ReadString")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text or \"exit\" to exit: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = line[:len(line)-1]
		fmt.Println(line)
		if string(line) == "exit" {
			break
		}
	}
	fmt.Println("Viszlát!")
}

// bufio.NewScanner.Scan/Text
func read4() {
	fmt.Println("bufio.NewScanner.Scan/Text")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter text or \"exit\" to exit: ")
		scanner.Scan()
		text := scanner.Text()
		fmt.Println(text)
		if text == "exit" {
			break
		}
	}
	fmt.Println("Viszlát!")
}

// fmt.Scanln
func read5() {
	fmt.Println("fmt.Scanln")
	for {
		var text1, text2 string
		fmt.Print("Enter text or \"exit\" to exit: ")
		fmt.Scanln(&text1, &text2)
		fmt.Println(text1, text2)
		if text1 == "exit" {
			break
		}
	}
	fmt.Println("Viszlát!")
}
