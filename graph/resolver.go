package graph

import "go-graphql-test/job_listings"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	JobListingService job_listings.Service
}
