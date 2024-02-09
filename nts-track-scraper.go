package main

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly"
)

type Track struct {
	Title  string
	Artist string
}

func main() {
	fmt.Println("Running...")

	c := colly.NewCollector(
		colly.URLFilters(
			regexp.MustCompile(
				`https:\/\/www\.nts\.live\/shows\/\S+\/episodes\/\S+-\d{1,2}\w{1,2}-\w+-\d{4}`,
			),
			regexp.MustCompile(
				`https:\/\/www\.nts\.live\/explore\?genres\[\]=housetechno-techno&genres\[\]=housetechno-acid`,
			),
		),
	)

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")

		fmt.Printf("Link found: %q -> %s\n", h.Text, link)

		c.Visit(h.Request.AbsoluteURL(link))
	})

	c.OnHTML("div.track__detail", func(h *colly.HTMLElement) {
		title := h.ChildText("span.track__title")
		artist := h.ChildText("span.track__artist--mobile")

		fmt.Printf("Track found: %s by %s\n", title, artist)

		// track := Track{
		// 	Title:  title,
		// 	Artist: artist,
		// }
	})

	c.Visit("https://www.nts.live/explore?genres[]=housetechno-techno&genres[]=housetechno-acid")

	fmt.Println("Complete.")
}
