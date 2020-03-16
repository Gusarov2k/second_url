package db

import (
	"context"
	"github.com/Gusarov2k/second_url"
	_ "github.com/jmoiron/sqlx"
	"log"
)

// urlsRepo is a service for managing URLs.
type urlsRepo struct {
	client *Client
}

// NewURLRepository creates a new NewURLRepository instance backed by Postgres.
func NewURLRepository(c *Client) shorten.URLRepository {
	return &urlsRepo{client: c}
}

// Create URLs's information into repository.
func (r *urlsRepo) Create(ctx context.Context, u *shorten.URL) error {
	row, err := r.client.db.NamedQueryContext(ctx, urlInsert, u)

	for row.Next() {
		row.StructScan(&u)

		if err != nil {
			log.Fatal(err)
		}
	}

	return err
}

// ByCode returns an URL object by code.
func (r *urlsRepo) ByCode(ctx context.Context, u *shorten.URL) (shorten.URL, error) {
	row, err := r.client.db.NamedQueryContext(ctx, urlByCode, u)

	var event shorten.URL
	for row.Next() {
		row.StructScan(&event)

		if err != nil {
			log.Fatal(err)
		}
	}
	return event, err
}

// Update URLs's information into repository.
func (r *urlsRepo) Update(ctx context.Context, u *shorten.URL) (shorten.URL, error) {
	row, err := r.client.db.NamedQueryContext(ctx, urlUpdate, u)

	var event shorten.URL
	for row.Next() {
		row.StructScan(&event)

		if err != nil {
			log.Fatal(err)
		}
	}
	return event, err

}

// Update URLs's information into repository.
func (r *urlsRepo) Delete(ctx context.Context, u *shorten.URL) error {
	_, err := r.client.db.NamedQueryContext(ctx, urlDelete, u)

	return err
}
