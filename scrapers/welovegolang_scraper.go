package scrapers

import (
	"context"
	"fmt"
	"go-graphql-test/job_listings"

	"github.com/gocolly/colly"
)

type welovegolangScraper struct {
	jobListingService job_listings.Service
}

func NewWelovegolangScraper(jobListingService job_listings.Service) JobListingScraper {
	return &welovegolangScraper{jobListingService}
}

func (s *welovegolangScraper) Scrape() {
	const url = "https://www.welovegolang.com/jobs"
	c := colly.NewCollector()

	c.OnHTML("div.stream-job", func(e *colly.HTMLElement) {

		company := e.DOM.Find("div.company span").First().Text()
		title := e.ChildText("p.summary")
		listingLink := url + e.DOM.Find("div.media-body a").First().AttrOr("href", "")

		s.jobListingService.AddJobListing(context.TODO(), company, title, listingLink, url)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)
	c.Wait()
	fmt.Println("Done with", url)
}
