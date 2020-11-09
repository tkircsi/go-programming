// Fetch downloads the URL and returns the
// name and length of the local file.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	filename, n, err := fetch(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("%s save as %s (%d btyes)\n", os.Args[1], filename, n)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
