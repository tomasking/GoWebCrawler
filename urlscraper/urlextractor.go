package urlscraper

import (
	"regexp"
	"strings"
)

func extractLinks(page string) []string {

	r := "<a\\s+(?:[^>]*?\\s+)?href=\"([^\"]*)\""
	re := regexp.MustCompile(r)
	matches := re.FindAllString(page, -1)

	for index, match := range matches {

		i := strings.Index(match, "href")
		match = match[i:]

		i = strings.Index(match, "\"")
		match = match[i+1:]

		i = strings.Index(match, "\"")
		match = match[:i]

		matches[index] = match
	}

	return matches
}
