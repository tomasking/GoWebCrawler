package main

import (
	"./crawler"
)

func main() {

	complete := make(chan bool, 1)

	crawler.Crawl("https://monzo.com/", complete)

	<-complete
}
