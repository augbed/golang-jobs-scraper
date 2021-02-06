package job_listings

import (
	"context"
	"database/sql"
	"log"
)

type Repository interface {
	GetListings(ctx context.Context) ([]*JobListing, error)
	AddNewListing(ctx context.Context, jobListing *JobListing) error
	FindListing(ctx context.Context, title, company, listingLink, website string) (*JobListing, error)
}

func NewSQLRepository(db *sql.DB) Repository {
	return &sqlRepository{db}
}

type sqlRepository struct {
	db *sql.DB
}

func (r *sqlRepository) GetListings(ctx context.Context) ([]*JobListing, error) {
	stmt, err := r.db.Prepare("select id, title, company, listing_link, website from job_listings")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var jobListings []*JobListing
	for rows.Next() {
		var jobListing JobListing
		err := rows.Scan(&jobListing.ID, &jobListing.Title, &jobListing.Company, &jobListing.ListingLink, &jobListing.Website)
		if err != nil {
			log.Fatal(err)
		}
		jobListings = append(jobListings, &jobListing)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return jobListings, nil
}
func (r *sqlRepository) FindListing(ctx context.Context, title, company, listingLink, website string) (*JobListing, error) {
	sqlQuery := `SELECT id, title, company, listing_link, website FROM job_listings WHERE title=? AND company=? AND listing_link=? AND website=?;`
	listing := JobListing{}
	row := r.db.QueryRowContext(ctx, sqlQuery, title, company, listingLink, website)
	err := row.Scan(&listing.ID, &listing.Title, &listing.Company, &listing.ListingLink, &listing.Website)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &listing, nil
	default:
		log.Println(err)
		return nil, err
	}
}

func (r *sqlRepository) AddNewListing(ctx context.Context, jobListing *JobListing) error {
	stmt, err := r.db.Prepare("INSERT INTO job_listings(title, company, website, listing_link) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(jobListing.Title, jobListing.Company, jobListing.Website, jobListing.ListingLink)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("New job listing persisted")
	return nil
}
