package crawler

import (
	"fmt"

	"../urlscraper"
)

func Crawl(seedUrl string, complete chan bool) {

	urlsToProcessQueue := make(chan string, 100)
	urlsToProcessQueue <- seedUrl

	go processQueue(urlsToProcessQueue, complete)
	go processQueue(urlsToProcessQueue, complete)
	go processQueue(urlsToProcessQueue, complete)
	go processQueue(urlsToProcessQueue, complete)
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



var collection []string

func add(url string){
	// Enter Mutex
	// Logic to add
	// Exit Mutex
}

//TODO use map


func contains(s []string, e string) bool {
	var m map[string]struct{}

	if _, ok := m[""]; ok {

	}

	// Enter Mutex
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
	// Exit Mutex
}

func init(){
	for {
		select {
		case val1 := <-chanA:
			collection = append(collection, val1)
		case string, chanB := <-chanC;
			chanB <- true / false
		}
	}
}
