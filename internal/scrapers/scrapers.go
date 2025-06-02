package scrapers

import (
	"fmt"
	"log"

	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

var (
	c    *colly.Collector
	jobs [][]Job
	mu   sync.Mutex
	wg   sync.WaitGroup
)

type Job struct {
	Title  string `json:"title"`
	Link   string `json:"link"`
	Source string `json:"source"`
}

func newCollector() *colly.Collector {
	c = colly.NewCollector(
		colly.AllowedDomains(
			"jobs.dou.ua",
			"layboard.com",
			"www.work.ua",
			"djinni.co",
			"relocate.me",
		),
	)

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"
	return c
}

func parseDOU(role string) {
	var job []Job
	url := fmt.Sprintf("https://jobs.dou.ua/vacancies/?category=%s&exp=1-3", role)
	c := newCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "vt" {
			link := e.Request.AbsoluteURL(e.Attr("href"))
			title := strings.TrimSpace(e.Text)

			data := Job{
				Title:  title,
				Link:   link,
				Source: "DOU",
			}

			job = append(job, data)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL.String())
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Println("Failed to visit: ", err)
	}

	mu.Lock()
	jobs = append(jobs, job)
	mu.Unlock()
}

func parseLayboard(role string) {
	var job []Job
	url := fmt.Sprintf("https://layboard.com/vakansii/search?q=%s", role)
	c := newCollector()

	c.OnHTML("a.card__body.js-card", func(e *colly.HTMLElement) {
		title := e.ChildText("span.card__title")
		link := e.Request.AbsoluteURL(e.Attr("href"))

		data := Job{
			Title:  title,
			Link:   link,
			Source: "Layboard",
		}

		job = append(job, data)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL.String())
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Println("Failed to visit: ", err)
	}

	mu.Lock()
	jobs = append(jobs, job)
	mu.Unlock()
}

func parseWorkUa(role string) {
	var job []Job
	c := newCollector()
	url := fmt.Sprintf("https://www.work.ua/jobs-%s", role)

	c.OnHTML("h2.my-0 a", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		title := strings.TrimSpace(e.Text)

		data := Job{
			Title:  title,
			Link:   link,
			Source: "WorkUA",
		}

		job = append(job, data)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL.String())
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Println("Failed to visit: WORK", err)
	}

	mu.Lock()
	jobs = append(jobs, job)
	mu.Unlock()
}

func parseDjini(role string) {
	var job []Job
	var maxPage int
	url := fmt.Sprintf("https://djinni.co/jobs/?primary_keyword=%s&exp_level=no_exp&exp_level=1y&exp_level=2y", role)
	c := newCollector()

	c.OnHTML("a.job-item__title-link", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		title := strings.TrimSpace(e.Text)

		data := Job{
			Title:  title,
			Link:   link,
			Source: "Djini",
		}

		job = append(job, data)

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

	err := c.Visit(url)
	if err != nil {
		log.Printf("Failed to visit: %v", err)
	}

	for page := 2; page <= maxPage; page++ {
		url := fmt.Sprintf("%s&page=%d", url, page)
		err := c.Visit(url)
		if err != nil {
			log.Printf("Failed to visit %s: %v", url, err)
		}
	}

	mu.Lock()
	jobs = append(jobs, job)
	mu.Unlock()
}

func parseRelocateMe(role string) {
	var job []Job
	c := newCollector()
	url := fmt.Sprintf("https://relocate.me/international-jobs?query=%s", role)

	c.OnHTML("div.job__title a", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		title := strings.TrimSpace(e.Text)

		data := Job{
			Title:  title,
			Link:   link,
			Source: "RelocateMe",
		}

		job = append(job, data)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL.String())
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Println("Failed to visit: ", err)
	}

	mu.Lock()
	jobs = append(jobs, job)
	mu.Unlock()
}

func StartScrapers(role string) ([][]Job, error) {
	jobs = [][]Job{}

	allScrapers := []func(string){
		parseDOU,
		parseDjini,
		parseLayboard,
		parseRelocateMe,
		parseWorkUa,
	}

	wg.Add(len(allScrapers))

	for _, scraper := range allScrapers {
		go func(s func(string)) {
			defer wg.Done()
			s(role)
		}(scraper)
	}

	wg.Wait()

	return jobs, nil
}
