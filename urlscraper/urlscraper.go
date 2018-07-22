package urlscraper

import (
	"../httpclient"
)

func ScrapeLinksFromPageUrl(url string) []string {

	page := httpclient.LoadPage(url)

	linksFoundInPage := extractLinks(page)

	filteredLinks := filterLinks(linksFoundInPage)

	return filteredLinks
}
