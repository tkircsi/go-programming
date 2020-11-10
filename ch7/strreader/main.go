package main

import (
	"fmt"
	"io"
)

type StrReader struct {
	s string
	i int
}

type LimitReader struct {
	R io.Reader // underlying reader
	N int       // max bytes remaining
}

func NewLimitReader(r io.Reader, n int) *LimitReader {
	return &LimitReader{r, n}
}

func (l *LimitReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int(n)
	return
}

func (r *StrReader) Read(b []byte) (n int, err error) {
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += n
	return n, nil
}

func NewStrReader(s string) *StrReader {
	return &StrReader{s: s}
}

const myHTML = `<html>
	<head>
		<title>Go Programming Language</title>
	</head>
	<body>
		<h1>Go Lang</h1>
	</body>
</html>`

func main() {
	r := NewStrReader(myHTML)

	b := make([]byte, 5)
	var res []byte
	for n, err := r.Read(b); err != io.EOF; n, err = r.Read(b) {
		// fmt.Printf("n: %d err:%v b:%s\n", n, err, b[:n])
		res = append(res, b[:n]...)
	}
	fmt.Println(string(res))

	r = NewStrReader(myHTML)
	lr := NewLimitReader(r, 20)
	res = nil
	b = make([]byte, 7)
	for n, err := lr.Read(b); err != io.EOF; n, err = lr.Read(b) {
		// fmt.Printf("n: %d, err:%v b:%s\n", n, err, b[:n])
		res = append(res, b[:n]...)
	}
	fmt.Println(string(res))
	// b := make([]byte, len(myHTML))
	// n, _ := r.Read(b)
	// fmt.Println(string(b[:n]))
	// doc, _ := html.Parse(r)
	// forEachNode(doc)
}

// func forEachNode(n *html.Node) {
// 	if n.Type == html.ElementNode {
// 		fmt.Println(n.Data)
// 	}
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		forEachNode(c)
// 	}
// 	if n.Type == html.ElementNode {
// 		fmt.Println(n.Data)
// 	}
// }
