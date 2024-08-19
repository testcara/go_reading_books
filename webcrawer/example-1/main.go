// does not work well
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type Visitor struct {
	visited map[string]bool
	mux     sync.Mutex
}

func Crawl(url string, depth int, fetcher Fetcher) {
	quit := make(chan bool)
	var v Visitor
	v.visited = make(map[string]bool)
	go MyCrawl(url, depth, fetcher, quit, &v)
	<-quit
}

func MyCrawl(url string, depth int, fetcher Fetcher, quit chan bool, v *Visitor) {
	if depth <= 0 {
		quit <- true
		return
	}
	v.mux.Lock()
	if !v.visited[url] {
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			v.visited[url] = true
			v.mux.Unlock()
			quit <- true
			return
		}
		v.visited[url] = true
		fmt.Println("found: %s %q\n", url, body)
		v.mux.Unlock()
		for _, u := range urls {
			childquit := make(chan bool)
			go MyCrawl(u, depth-1, fetcher, childquit, v)
			<-childquit
		}
	} else {
		v.mux.Unlock()
	}
	quit <- true
}

func main() {
	Crawl("https://golang.org", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://go.dev/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
