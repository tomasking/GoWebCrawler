package crawler

import (
	"fmt"

	"../urlscraper"
)

func Crawl(seedUrl string, complete chan bool) {

	urlsToProcessQueue := make(chan string, 1000)
	urlsToProcessQueue <- seedUrl

	go processQueue(urlsToProcessQueue, complete)

}

func processQueue(queue chan string, complete chan bool) {

	visited := []string{}

	for url := range queue {

		if !contains(visited, url) {

			fmt.Println("Processing: " + url)
			visited = append(visited, url)

			scrapeURLAndAddToQueue(url, queue, visited)
		}

		if len(queue) == 0 {
			break
		}
	}

	complete <- true
}

func scrapeURLAndAddToQueue(url string, queue chan string, visited []string) {

	filteredLinks := urlscraper.ScrapeLinksFromPageUrl(url)
	for _, childURL := range filteredLinks {
		if contains(visited, childURL) {
			continue
		}
		queue <- childURL
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
