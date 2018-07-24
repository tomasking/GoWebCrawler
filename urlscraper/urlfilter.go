package urlscraper

import (
	"strings"
)

func filterLinks(inputLinks []string) []string {

	filteredLinks := []string{}

	for _, link := range inputLinks {
		link = sanitise(link)
		if strings.Contains(link, "mycompany.com") {
			filteredLinks = append(filteredLinks, link)
		}
	}

	return filteredLinks
}

func sanitise(input string) string {
	input = strings.ToLower(input)

	if strings.HasPrefix(input, "/") {
		input = "https://mycompany.com" + input
	}

	return input
}
