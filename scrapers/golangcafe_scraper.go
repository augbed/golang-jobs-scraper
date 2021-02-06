package scrapers

import (
	"context"
	"fmt"
	"go-graphql-test/job_listings"
	"time"

	"github.com/gocolly/colly"
)

type golangcafeScraper struct {
	jobListingService job_listings.Service
}

func NewGolangcafeScraper(jobListingService job_listings.Service) JobListingScraper {
	return &golangcafeScraper{jobListingService}
}

func (s *golangcafeScraper) Scrape() {
	const url = "https://golang.cafe/Remote-Golang-Jobs"
	c := colly.NewCollector(
		colly.AllowedDomains("golang.cafe"),
	)

	c.OnHTML("article.line-item", func(e *colly.HTMLElement) {

		title := e.DOM.Find("div a").First().Find("b").First().Text()
		company := e.DOM.Find("div a").First().Siblings().Filter("a").First().Text()

		if title != "" && company != "" {
			err := s.jobListingService.AddJobListing(context.TODO(), company, title, url, url)
			if err != nil {
				fmt.Println(err)
			}
		}
	})

	// Visit all pagination links
	c.OnHTML("section nav a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Limit(&colly.LimitRule{
		RandomDelay: 1 * time.Second,
	})

	c.Visit(url)
	c.Wait()
	fmt.Println("Done with", url)
}
