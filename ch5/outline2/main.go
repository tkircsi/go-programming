// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func main() {
	resp, err := http.Get("http://gopl.io")
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %s\n", err)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	forEachNode(doc, startElement, endElement)
}

// forEachNode traverse over the html document and call pre and post
// functions for each node
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {

	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, getAttrStr(n.Attr))
		// fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}

	if n.Type == html.TextNode {
		s := strings.TrimSpace(n.Data)
		if len(s) > 0 {
			fmt.Printf("%*s %s\n", depth*2, "", n.Data)
		}
	}
}

func getAttrStr(attrs []html.Attribute) string {
	var attrStr string
	for _, attr := range attrs {
		attrStr += fmt.Sprintf(" %s=\"%s\"", attr.Key, attr.Val)
	}
	return attrStr
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
