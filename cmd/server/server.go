package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Job struct {
	Title    string
	Company  string
	Location string
}

func RunServerStuff() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	router.GET("/", WelcomePage)
	router.GET("/search", SearchJob)
	router.Run(":8080")
}

func SearchJob(c *gin.Context) {
	role := c.Query("role")
	jobs := []Job{
		{Title: "Backend Developer", Company: "TechCorp", Location: "Remote"},
		{Title: "Go Developer", Company: "StartupX", Location: "Berlin"},
	}

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
