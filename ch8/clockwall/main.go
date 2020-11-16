package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	for _, s := range os.Args[1:] {
		go listenClock(s)
	}
	for {
		time.Sleep(1 * time.Minute)
	}
}

func listenClock(server string) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
