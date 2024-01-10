package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func createDownloadDir() string {
	downloadDir := "download"
	if err := os.MkdirAll(downloadDir, 0755); err != nil {
		log.Fatal("Error creating download directory:", err)
	}
	return downloadDir
}

func downloadFile(url string, downloadDir string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	filename := strings.TrimPrefix(url, "http://localhost:8080/")
	filepath := filepath.Join(downloadDir, filename)
	err = ioutil.WriteFile(filepath, body, 0644)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func extractLinks(body []byte, currentURL string, toVisit chan string, visited *sync.Map, wg *sync.WaitGroup) {
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		log.Println("Error parsing HTML:", err)
		return
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := url.Parse(a.Val)
					if err != nil {
						log.Println("Error parsing URL:", err)
						continue
					}
					base, err := url.Parse(currentURL)
					if err != nil {
						log.Println("Error parsing base URL:", err)
						continue
					}
					absoluteURL := base.ResolveReference(link).String()
					if _, exists := visited.LoadOrStore(absoluteURL, true); !exists {
						wg.Add(1) // Increment the counter before enqueuing the URL
						toVisit <- absoluteURL
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func crawl(toVisit chan string, downloadDir string, visited *sync.Map, wg *sync.WaitGroup) {
	for url := range toVisit {
		fmt.Println("Visiting:", url)
		body, err := downloadFile(url, downloadDir)
		if err != nil {
			log.Println("Error:", err)
			defer wg.Done()
			continue
		}

		func() {
			defer wg.Done() // Decrement the counter after processing the URL
			extractLinks(body, url, toVisit, visited, wg)
		}()
	}
}

func main() {
	fmt.Println("# 42Snippets")
	fmt.Println("## Golang Crawler")

	/*
		Concurrency in this crawler snippet is achieved through the use of goroutines, channels, and synchronization primitives like WaitGroups and sync maps.

		It launches multiple worker goroutines to process URLs concurrently. Each worker goroutine reads URLs from the "to visit" channel and processes them independently of the others. This allows the program to download and parse multiple HTML files simultaneously, taking advantage of multi-core processors.

		The "to visit" channel is used to distribute URLs among the worker goroutines. By sending URLs to the channel and having each worker goroutine read from it, the program ensures that URLs are processed in a coordinated manner. The channel also acts as a queue, allowing the program to handle a large number of URLs without overwhelming the system.

		The WaitGroup is used to synchronize the main goroutine with the worker goroutines. The main goroutine waits for all worker goroutines to finish processing the URLs before printing the "Crawling complete!" message. This ensures that the program doesn't exit prematurely, leaving some URLs unprocessed.

		The sync map is used to keep track of visited URLs. By using a sync map instead of a regular map, the program ensures that multiple goroutines can safely read and write to the map without causing data races. This allows the program to efficiently avoid revisiting the same URL multiple times.

		The combination of channels, WaitGroups, and sync maps ensures that the URLs are processed concurrently but without duplication or data races. The use of defer wg.Done() ensures that the WaitGroup counter is correctly decremented, even if an error occurs or a return statement is executed early.

		This is a slightly modified version of the original 42Snippets golangcrawler snippet "Golang Crawler".
	*/

	downloadDir := createDownloadDir()
	startURL := "http://localhost:8080/file1.html"
	toVisit := make(chan string, 100)
	visited := &sync.Map{}
	var wg sync.WaitGroup
	var workerWg sync.WaitGroup

	visited.Store(startURL, true)
	wg.Add(1)
	toVisit <- startURL

	workers := 7
	for i := 0; i < workers; i++ {
		workerWg.Add(1)
		go func() {
			defer workerWg.Done()
			crawl(toVisit, downloadDir, visited, &wg)
		}()
	}

	wg.Wait()
	close(toVisit)
	workerWg.Wait()

	fmt.Println("Crawling complete!")
}
