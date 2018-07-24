package main

import (
	"./crawler"
)

func main() {

	complete := make(chan bool, 1)

	crawler.Crawl("https://mycompany.com", complete) //TODO: pass in URL here

	<-complete
}
