package main

import (
    "fmt"
    "sync"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

// SafeCounter is safe to use concurrently.
type SafeUrlMap struct {
    v   map[string]int
    mux sync.Mutex
}

func (um SafeUrlMap) hasUrl(k string) bool {
    defer um.mux.Unlock()
    um.mux.Lock()
    _, ok := um.v[k]
    return ok
}

func (um SafeUrlMap) registerUrl(k string) {
    defer um.mux.Unlock()
    um.mux.Lock()
    um.v[k] = 1
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
    var wg sync.WaitGroup
    var crawl func(string, int)
    
    urlMap := SafeUrlMap{v: make(map[string]int)}
    ch := make(chan string)
    crawl = func(url string, depth int) {
        defer wg.Done()

        if urlMap.hasUrl(url) {
            return
        } else {
            urlMap.registerUrl(url)
        }
        if depth <= 0 {
            return
        }
        body, urls, err := fetcher.Fetch(url)
        if err != nil {
            ch <- fmt.Sprintf("%v",err)
            return
        }
        ch <- fmt.Sprintf("found: %s %q", url, body)
        for _, u := range urls {
            wg.Add(1)
            go crawl(u, depth-1)
        }
        return
    }
    wg.Add(1)
    go crawl(url, depth)

    var collect_wg sync.WaitGroup
    collect_wg.Add(1)
    go func() {
        defer collect_wg.Done()
        for s := range ch {
            fmt.Println(s)    
        }
    }()

    wg.Wait()
    close(ch)
    collect_wg.Wait()
}

func main() {
    Crawl("http://golang.org/", 4, fetcher)
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
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}

