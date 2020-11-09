package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	bigSlowOperation()

	i := double(4)
	fmt.Println(i)

	i = triple(5)
	fmt.Println(i)
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(3 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter: %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d)  = %d\n", x, result) }()
	return 2 * x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}
