//File_name: webcrawler.go
//Author: Wenbin Xiao
//Description: http://tour.golang.org/#73

package main

import (
	"fmt"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	//map to record the fecthed-or-not state and a lock to synchronize
	fetched := make(map[string]bool)
	fetched_lock := make(chan int, 1)
	fetched_lock <- 1

	//Function to help fetch
	var fetch func(url string, depth int)
	fetch = func(url string, depth int) {
		//Get the lock
		<-fetched_lock
		//Release the lock when this function exits
		defer func() {
			fetched_lock <- 1
		}()

		//Check whether fetched
		if _, ok := fetched[url]; ok {
			return
		}
		fetched[url] = true

		//Check depth
		if depth <= 0 {
			return
		}
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		//Found
		fmt.Printf("found: %s %q\n", url, body)
		//Crawl to others
		for _, u := range urls {
			go fetch(u, depth-1)
		}
		return
	}
	//Start the crawling
	go fetch(url, depth)
	time.Sleep(1 * time.Second)
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
