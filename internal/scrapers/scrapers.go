package scrapers

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func ParseDOU() []string {
	var links []string
	startURL := "https://jobs.dou.ua/vacancies/?category=Golang&exp=1-3"

	c := colly.NewCollector(
		colly.AllowedDomains(
			"jobs.dou.ua",
		),
	)

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "vt" {
			links = append(links, e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL.String())
	})

	err := c.Visit(startURL)
	if err != nil {
		fmt.Println("Failed to visit: ", err)
	}

	return links
}

func ParseLayboard() []string {
	var links []string
	var finallLinks []string
	startURL := "https://layboard.com/vakansii/search?q=golang"

	c := colly.NewCollector(
		colly.AllowedDomains(
			"layboard.com",
		),
	)

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "card__body js-card" {
			links = append(links, e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL.String())
	})

	err := c.Visit(startURL)
	if err != nil {
		fmt.Println("Failed to visit: ", err)
	}

	for _, v := range links {
		fullLink := "https://layboard.com" + v
		finallLinks = append(finallLinks, fullLink)
	}

	return finallLinks
}

func ParseWorkUa() []string {
	var res []string
	var res2 []string
	c := colly.NewCollector(
		colly.AllowedDomains(
			"www.work.ua",
		),
	)

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"

	c.OnHTML("h2.my-0 a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		res = append(res, link)

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL.String())
	})

	err := c.Visit("https://www.work.ua/jobs-golang/")
	if err != nil {
		fmt.Println("Failed to visit: ", err)
	}

	for _, v := range res {
		fullLink := "https://work.ua" + v
		res2 = append(res2, fullLink)
	}
	return res2
}

func ParseDjini() []string {
	var res []string
	var res2 []string
	var maxPage int
	startURL := "https://djinni.co/jobs/?primary_keyword=iOS&exp_level=no_exp&exp_level=1y&exp_level=2y"

	c := colly.NewCollector(
		colly.AllowedDomains(
			"djinni.co",
		),
	)

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "job-item__title-link" {
			res = append(res, e.Attr("href"))
		}

	})

	c.OnHTML("li.page-item a.page-link", func(e *colly.HTMLElement) {
		pageText := strings.TrimSpace(e.Text)

		if pageNum, err := strconv.Atoi(pageText); err == nil {
			if pageNum > maxPage {
				maxPage = pageNum
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL.String())
	})

	err := c.Visit(startURL)
	if err != nil {
		log.Printf("Failed to visit: %v", err)
	}

	for page := 2; page <= maxPage; page++ {
		url := fmt.Sprintf("%s&page=%d", startURL, page)
		err := c.Visit(url)
		fmt.Println(page)
		if err != nil {
			log.Printf("Failed to visit %s: %v", url, err)
		}
	}

	for _, v := range res {
		fullLink := "https://djinni.co" + v
		res2 = append(res2, fullLink)
	}
	return res2
}

func ParseRelocateMe() []string {
	var res []string
	var res2 []string
	c := colly.NewCollector(
		colly.AllowedDomains(
			"relocate.me",
		),
	)

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"

	c.OnHTML("div.job__title a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		res = append(res, link)

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL.String())
	})

	err := c.Visit("https://relocate.me/international-jobs?query=golang")
	if err != nil {
		fmt.Println("Failed to visit: ", err)
	}

	for _, v := range res {
		fullLink := "https://relocate.me" + v
		res2 = append(res2, fullLink)
	}
	return res2
}
