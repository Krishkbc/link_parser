package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	nodes := LinkNodes(doc)

	var links []Link

	for _, node := range nodes {
		links = append(links, BuildLink(node))
		fmt.Println(node)
	}

	return links, nil
}

func LinkNodes(n *html.Node) []*html.Node {

	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, LinkNodes(c)...)
	}
	return ret
}

func BuildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.text = strings.Join(strings.Fields(Text(n)), " ")
	return ret
}

func Text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += Text(c) + " "
	}
	return ret
}

// func dfs(n *html.Node, padding string) {
// 	fmt.Println(padding, n.Data)
// 	msg := n.Data

// 	if n.Type == html.ElementNode {
// 		msg = "<" + msg + ">"
// 	}

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		dfs(c, padding+"  ")
// 	}
// }
