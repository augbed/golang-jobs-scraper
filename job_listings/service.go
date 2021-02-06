package job_listings

import (
	"context"
	"errors"
)

type Service interface {
	AddJobListing(ctx context.Context, company string, title string, listingLink string, website string) error
	ListJobListings(ctx context.Context) ([]*JobListing, error)
}

func NewJobListingService(repository Repository) Service {
	return &jobListingService{repository}
}

type jobListingService struct {
	repository Repository
}

func (s *jobListingService) AddJobListing(ctx context.Context, company string, title string, listingLink string, website string) error {
	listing, err := s.repository.FindListing(ctx, title, company, listingLink, website)
	if err != nil {
		return errors.New("Error while trying to get existing job listing")
	}
	if listing != nil {
		return nil
	}

	newJobListing := &JobListing{
		Title:       title,
		Company:     company,
		Website:     website,
		ListingLink: listingLink,
	}
	err = s.repository.AddNewListing(ctx, newJobListing)
	if err != nil {
		return errors.New("Error while trying to persist new job listing")
	}
	return nil
}

func (s *jobListingService) ListJobListings(ctx context.Context) ([]*JobListing, error) {
	list, err := s.repository.GetListings(ctx)
	if err != nil {
		return nil, errors.New("Error while trying to get job listings from database")
	}
	return list, nil
}
