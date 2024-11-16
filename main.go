package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/krishkbc/link_parser/link"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com/", "the url you want to get")
	depthFlag := flag.Int("depth", 3, "the depth of the url you want to get")
	flag.Parse()

	fmt.Println("url:", *urlFlag)
	pages := get(*urlFlag)

	for _, page := range pages {
		fmt.Println(page)
	}

}

func get(UrlString string) []string {
	resp, err := http.Get(UrlString)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}

	base := baseURL.String()
	fmt.Println("reqURL:", reqURL)
	fmt.Println("base:", base)

	return filter(href(base, resp.Body), withPrefix(base))

}

func href(base string, r io.Reader) []string {
	var ret []string
	links, err := link.Parse(r)
	_ = err

	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)

		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)

		}
	}
	return ret
}

func filter(links []string, keepfun func(string) bool) []string {
	var ret []string
	for _, l := range links {
		if keepfun(l) {
			ret = append(ret, l)
		}

	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})

	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}

	for i := 0; i <= maxDepth; i++ {
		q, nq := nq, make(map[string]struct{})
		for url, _ := range q {
			if _, ok := seen[url]; ok {
				continue
			}

		}
	}
}

// return the url string and tell the depthof the url string
// func bfs(urlStr string, maxDepth int) []string {
// 	seen := make(map[string]struct{})
