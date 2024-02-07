package main

import (
	"regexp"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.URLFilters(
			regexp.MustCompile(
				
			),
		),
	)

	c.
}