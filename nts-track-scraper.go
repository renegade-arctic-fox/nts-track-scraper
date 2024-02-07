package main

import (
	"regexp"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.URLFilters(
			regexp.MustCompile(
				`https://www.nts.live/shows/\S+/episodes/\S+-\d{1,2}\w{1,2}-\w+-\d{4}`,
			),
		),
	)

	c.
}