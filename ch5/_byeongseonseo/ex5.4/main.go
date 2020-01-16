package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 인터넷으로 찾은 것
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			switch a.Key {
			case "href":
			case "src":
				links = append(links, a.Val)
			}
		}
	}

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ex 5.1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

}
