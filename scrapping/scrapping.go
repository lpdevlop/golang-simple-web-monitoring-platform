package scrapping

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func StartScraping(url string) (string, string) {

	fmt.Println("call Start Scraping")

	c := colly.NewCollector()

	var title string
	var content string

	c.OnHTML("body", func(e *colly.HTMLElement) {
		content = e.Text
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.Visit(url)

	return title, content
}
