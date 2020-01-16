package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// 해당 노드가 style 혹은 script 태그 안에 정의 되어 있는 노드 인지 확인
func isThisNodeScriptOrStyle(n *html.Node) bool {
	if n == nil {
		return false
	}

	if n.Type != html.ElementNode {
		return isThisNodeScriptOrStyle(n.Parent)
	}

	return n.Data == "script" || n.Data == "style"
}

func visit(texts []string, n *html.Node) []string {

	if n.Type == html.TextNode && isThisNodeScriptOrStyle(n.Parent) == false {
		texts = append(texts, strings.TrimSpace(n.Data))
	}
	if n.FirstChild != nil {
		texts = visit(texts, n.FirstChild)
	}
	if n.NextSibling != nil {
		texts = visit(texts, n.NextSibling)
	}
	return texts
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ex 5.3: %v\n", err)
		os.Exit(1)
	}

	for _, text := range visit(nil, doc) {
		fmt.Println(text)
	}

}
