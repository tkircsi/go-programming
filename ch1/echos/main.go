// echos prints its command-line arguments
// the first argument selects the echo function
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var sArgs string
	if len(os.Args) > 1 {
		v := os.Args[1]
		switch v {
		case "1":
			sArgs = echo1()
		case "1v2":
			sArgs = echo1_2()
		case "2":
			sArgs = echo2()
		case "3":
			sArgs = echo3()
		default:
			sArgs = echo1()
		}
	}
	fmt.Println(sArgs)
}

func echo1() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo1_2() string {
	var s, sep string
	l := len(os.Args)
	for i := 1; i < l; i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo2() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3() string {
	return strings.Join(os.Args[1:], " ")
}
