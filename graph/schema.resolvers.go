package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-graphql-test/graph/generated"
	"go-graphql-test/graph/model"
)

func (r *queryResolver) JobListings(ctx context.Context) ([]*model.JobListing, error) {
	var listings []*model.JobListing
	dbListings, err := r.JobListingService.ListJobListings(ctx)
	if err != nil {
		return nil, err
	}
	for _, listing := range dbListings {
		listings = append(listings, &model.JobListing{
			ID:          listing.ID,
			Company:     listing.Company,
			Title:       listing.Title,
			Website:     listing.Website,
			ListingLink: listing.ListingLink,
		})
	}
	return listings, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
