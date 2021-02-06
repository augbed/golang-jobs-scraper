package scrapers

import (
	"context"
	"fmt"
	"go-graphql-test/job_listings"

	"github.com/gocolly/colly"
)

type remoteokScraper struct {
	jobListingService job_listings.Service
}

func NewRemoteokScraper(jobListingService job_listings.Service) JobListingScraper {
	return &remoteokScraper{jobListingService}
}

func (s *remoteokScraper) Scrape() {
	const url = "https://remoteok.io/remote-golang-jobs"
	c := colly.NewCollector()

	c.OnHTML("tr.job", func(e *colly.HTMLElement) {

		company := e.Attr("data-company")
		title := e.Attr("data-search")
		listingLink := url + e.Attr("data-href")

		s.jobListingService.AddJobListing(context.TODO(), company, title, listingLink, url)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)

	c.Wait()
	fmt.Println("Done with", url)
}
