package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// TAGMap 은 태그의 수를 새는 맵이다.
type TAGMap = map[string]int

func visit(counter *TAGMap, n *html.Node) {
	if n.Type == html.ElementNode {
		(*counter)[n.Data] = (*counter)[n.Data] + 1
	}

	if n.FirstChild != nil {
		visit(counter, n.FirstChild)
	}
	if n.NextSibling != nil {
		visit(counter, n.NextSibling)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ex 5.2: %v\n", err)
		os.Exit(1)
	}

	counter := make(TAGMap)
	visit(&counter, doc)
	for key, count := range counter {
		fmt.Println(key, count)
	}

}
