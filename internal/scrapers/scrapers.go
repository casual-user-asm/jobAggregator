package scrapers

import (
	"fmt"

	"github.com/gocolly/colly"
)

var c = colly.NewCollector()

func ParseData() {
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "vt" {
			fmt.Printf("Link is %v\n", e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL.String())
	})

	err := c.Visit("https://jobs.dou.ua/vacancies/?category=Golang")
	if err != nil {
		fmt.Println("Failed to visit: ", err)
	}
}
