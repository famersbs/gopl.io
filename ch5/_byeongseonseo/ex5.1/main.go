package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func visitSibling(links []string, n *html.Node) []string {

	if n == nil {
		return links
	}

	links = visit(links, n)

	if n.NextSibling == nil {
		return links
	}

	return visitSibling(links, n.NextSibling)

}

// 내가 푼것
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	return visitSibling(links, n.FirstChild)
}

// 인터넷으로 찾은 것
func visitFindFromInternet(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.FirstChild != nil {
		links = visitFindFromInternet(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visitFindFromInternet(links, n.NextSibling)
	}

	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ex 5.1: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("----------- mine")
	for _, link := range visitSibling(nil, doc) {
		fmt.Println(link)
	}

	fmt.Println("----------- answer from internet")
	for _, link := range visitFindFromInternet(nil, doc) {
		fmt.Println(link)
	}

}
