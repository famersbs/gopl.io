package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "ex5.5: %v\n", err)
			continue
		}

		fmt.Println(url, " : words - ", words, " images - ", images)
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return

}

func visit(n *html.Node, cb func(*html.Node)) {
	if n == nil {
		return
	}

	cb(n)

	if n.FirstChild != nil {
		visit(n.FirstChild, cb)
	}
	if n.NextSibling != nil {
		visit(n.NextSibling, cb)
	}
}

func countWordsAndImages(_n *html.Node) (words, images int) {

	visit(_n, func(n *html.Node) {
		switch n.Type {
		case html.TextNode:
			words += wordCount(n.Data)
		case html.ElementNode:
			if n.Data == "img" {
				images++
			}
		}
	})

	return
}

func wordCount(s string) int {
	n := 0
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n++
	}
	return n
}
