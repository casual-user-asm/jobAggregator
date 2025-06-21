
# Job Aggregator

A web application that aggregates job listings from multiple sources based on search queries.




## 🚀 Features

 - Search for jobs across multiple job platforms simultaneously
 - Results grouped by source for easy comparison
 - Clean, responsive user interface
 - Fast, asynchronous job scraping
 - Docker support for easy deployment

## 📋 Supported Job Sources

 - Work.ua
 - Djinni
 - DOU
 - Layboard
 - RelocateMe
 - (More sources can be easily added)

## 🛠️ Tech Stack

 - Backend: Go with Gin web framework
 - Web Scraping: Colly scraper
 - Frontend: HTML/CSS with Go templates
 - Deployment: Docker
 - Hosted on Fly.io



## 🏁 Running with Docker

Clone the repository

```bash
  git clone https://github.com/casual-user-asm/jobAggregator.git
  cd jobAggregator
```

Build the Docker image

```bash
  docker build -t job-aggregator .
```

Run the container

```bash
  docker run -p 8080:8080 job-aggregator
```
