package scrapers

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ParseData(class, link string) []string {
	var res []string

	c := colly.NewCollector(
		colly.AllowedDomains(
			"jobs.dou.ua",
			"layboard.com",
			"work.ua",
			"djinni.co",
			"toptal.com",
			"relocate.me",
			"robota.ua",
			"happymonday.ua",
			"jooble.org",
		),
	)

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Attr("class") == class {
			res = append(res, e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL.String())
	})

	err := c.Visit(link)
	if err != nil {
		fmt.Println("Failed to visit: ", err)
	}
	return res
}

func ParseDOU() []string {
	res := ParseData("vt", "https://jobs.dou.ua/vacancies/?category=Golang")
	return res
}

func ParseLayboard() []string {
	res := ParseData("card__body js-card", "https://layboard.com/vakansii/search?q=golang")
	res2 := []string{}
	for _, v := range res {
		fullLink := "https://layboard.com" + v
		res2 = append(res2, fullLink)
	}
	return res2
}

func ParseWorkUa() []string {
	var res []string
	c := colly.NewCollector(
		colly.AllowedDomains(
			"work.ua",
		),
	)

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Attr("tabindex") == "-1" {
			res = append(res, e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL.String())
	})

	err := c.Visit("https://www.work.ua/jobs-golang/")
	if err != nil {
		fmt.Println("Failed to visit: ", err)
	}
	return res
}
