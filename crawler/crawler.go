package crawler

import (
	"fmt"
	"sync"

	"../urlscraper"
)

var visited = make(map[string]bool)
var mutex = &sync.Mutex{}
var urlsToProcessQueue = make(chan string, 100)

// Crawl - performs traverse of site map
func Crawl(seedURL string, threads int) {

	urlsToProcessQueue <- seedURL

	complete := make(chan bool, threads)

	for threadNumber := 0; threadNumber < threads; threadNumber++ {
		go processQueue(complete, threadNumber)
	}

	for threadNumber := 0; threadNumber < threads; threadNumber++ {
		<-complete
	}
}

func processQueue(complete chan bool, thread int) {

	for url := range urlsToProcessQueue {
		if !haveVisited(url) {
			fmt.Println("Processing: ", url, thread)
			addToVisited(url)
			scrapeURLAndAddToQueue(url)
		}

		if len(urlsToProcessQueue) == 0 {
			break
		}
	}

	complete <- true
}

func addToVisited(url string) {
	mutex.Lock()
	visited[url] = true
	mutex.Unlock()
}

func haveVisited(url string) bool {
	mutex.Lock()
	hasVisited := visited[url]
	mutex.Unlock()
	return hasVisited
}

func scrapeURLAndAddToQueue(url string) {

	filteredLinks := urlscraper.ScrapeLinksFromPageUrl(url)
	for _, childURL := range filteredLinks {
		if haveVisited(childURL) {
			continue
		}
		urlsToProcessQueue <- childURL
	}
}
