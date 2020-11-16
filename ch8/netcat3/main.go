package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("waiting for 3s...")
		time.Sleep(3 * time.Second)
		log.Println("done")
		done <- struct{}{}
	}()
	<-done
	log.Println("finished getting data")
	conn.Close()

}

// func mustCopy(dst io.Writer, src io.Reader) {
// 	if _, err := io.Copy(dst, src); err != nil {
// 		log.Fatal("mustCopy ", err)
// 	}
// 	log.Println("end mustCopy")
// }
