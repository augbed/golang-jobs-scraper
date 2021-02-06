package main

import (
	"go-graphql-test/db/mysql"
	"go-graphql-test/job_listings"
	"go-graphql-test/scrapers"
	"sync"
)

func main() {
	repo := job_listings.NewSQLRepository(mysql.DB)
	jobListingService := job_listings.NewJobListingService(repo)
	jobListingScrapers := make([]scrapers.JobListingScraper, 0)
	jobListingScrapers = append(jobListingScrapers, scrapers.NewRemoteokScraper(jobListingService))
	jobListingScrapers = append(jobListingScrapers, scrapers.NewWelovegolangScraper(jobListingService))
	jobListingScrapers = append(jobListingScrapers, scrapers.NewGolangcafeScraper(jobListingService))

	// Scrape job listings concurrently
	var wg sync.WaitGroup
	for _, scraper := range jobListingScrapers {
		wg.Add(1)
		go func(scraper scrapers.JobListingScraper) {
			defer wg.Done()
			scraper.Scrape()
		}(scraper)
	}
	wg.Wait()
}
