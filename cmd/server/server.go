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

	jobs, err := scrapers.StartScrapers(role)
	if err != nil {
		fmt.Println("Error take data from scrapers")
	}
	fmt.Println(jobs)

	// data, err := os.ReadFile("data/jobsData.json")
	// if err != nil {
	// 	fmt.Println("Failed to read file")
	// }

	// err = json.Unmarshal(data, &jobs)
	// if err != nil {
	// 	fmt.Println("Failed to unmarshal json")
	// }

	c.HTML(http.StatusOK, "search.html", gin.H{
		"role": role,
		"jobs": jobs,
	})
}

func WelcomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Job Aggregator",
	})
}
