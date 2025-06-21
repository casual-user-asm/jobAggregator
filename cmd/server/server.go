package server

import (
	"fmt"
	"net/http"

	"jobAggregator/internal/scrapers"

	"github.com/gin-gonic/gin"
)

type Job struct {
	Title  string `json:"title"`
	Source string `json:"source"`
	Link   string `json:"link"`
}

func RunServerStuff() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/assets", "./web/assets")
	router.Static("/images", "./web/images")
	router.StaticFile("/favicon.ico", "./web/assets/favicon.ico")

	router.GET("/", WelcomePage)
	router.GET("/search", SearchJob)
	router.Run(":8080")
}

func SearchJob(c *gin.Context) {
	role := c.Query("role")

	scraperJobs, err := scrapers.StartScrapers(role)
	if err != nil {
		fmt.Println("Error take data from scrapers")
	}

	jobsBySource := make(map[string][]Job)
	hasJobs := false

	if len(scraperJobs) > 0 {
		hasJobs = true

		for _, jobList := range scraperJobs {
			for _, scraperJob := range jobList {
				job := Job{
					Title:  scraperJob.Title,
					Source: scraperJob.Source,
					Link:   scraperJob.Link,
				}
				source := job.Source
				jobsBySource[source] = append(jobsBySource[source], job)
			}
		}
	}
	c.HTML(http.StatusOK, "search.html", gin.H{
		"role":         role,
		"jobsBySource": jobsBySource,
		"hasJobs":      hasJobs,
	})
}

func WelcomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Job Aggregator",
	})
}
