package main

import (
	"fmt"
	"sync"
)

type SafeSaver struct {
	mu      sync.Mutex
	visited map[string]bool
	wg      sync.WaitGroup
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func (s *SafeSaver) checkVisited(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.visited[url]

	if ok == false {
		s.visited[url] = true
		return false
	}

	return true
}

func Crawl(url string, depth int, fetcher Fetcher, V *SafeSaver) {
	defer V.wg.Done()

	if depth <= 0 {
		return
	}

	if V.checkVisited(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		V.wg.Add(1)
		go Crawl(u, depth-1, fetcher, V)
	}
	return
}

func main() {
	V := SafeSaver{visited: make(map[string]bool)}
	V.wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, &V)
	V.wg.Wait()
}

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
	"https://golang.org/": &fakeResult{
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
