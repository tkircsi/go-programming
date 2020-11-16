// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var p = flag.Int("p", 8000, "network port on clock server listens")
var n = flag.String("n", "localhost", "name of the service")
var it = flag.Int("it", 5, "iteration")

func main() {
	flag.Parse()
	// listener listens for incoming connection on a network port
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *p))
	if err != nil {
		log.Fatal(err)
	}
	for {
		// Accpet() blocks until incoming connection request
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, *n, *it)
	}
}

func handleConn(c net.Conn, name string, iter int) {
	log.Println("enter handleConn")
	defer c.Close()
	for i := 0; i < iter; i++ {
		_, err := io.WriteString(c, fmt.Sprintf("Place: %s, Time: %s", name, time.Now().Format("15:04:05\n")))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	log.Println("end handleConn")
}
