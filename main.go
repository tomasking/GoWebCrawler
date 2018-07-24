package main

import (
	"fmt"
	"time"

	"./crawler"
)

func main() {

	start := time.Now()
	crawler.Crawl("https://mycompany.com/", 10)
	elapsed := time.Since(start)

	fmt.Println("Completed in ", elapsed)
}
