package main

import (
	// "jobAggregator/cmd/server"
	"fmt"
	"jobAggregator/internal/scrapers"
)

func main() {
	for _, link := range scrapers.ParseDjini() {
		fmt.Println(link)
	}
}
