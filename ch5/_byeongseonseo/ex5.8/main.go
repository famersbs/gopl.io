package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func elementByID(n *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}

		}
		return true
	}
	return forEachElement(n, pre, nil)
}

func forEachElement(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	u := make([]*html.Node, 0) // unvisited
	u = append(u, n)
	for len(u) > 0 {
		n = u[0]
		u = u[1:]
		if pre != nil {
			if !pre(n) {
				return n
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}
		if post != nil {
			if !post(n) {
				return n
			}
		}
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: ex5.8 HTML_FILE ID")
	}

	url := os.Args[1]
	id := os.Args[2]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	node := elementByID(doc, id)
	fmt.Printf("Type %d, Tag : %s, Attr : %v\n", node.Type, node.Data, node.Attr)
}
