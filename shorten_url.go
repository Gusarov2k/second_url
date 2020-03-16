package shorten

import "context"

type URL struct {
	ID   int    `db:"id" json:"id"`
	URL  string `db:"original_url" json:"url"`
	Code string `db:"code" json:"code"`
}

// UrlRepository is a storage for urls.
type URLRepository interface {
	// Create creates a new url.
	Create(ctx context.Context, u *URL) error
	// Update creates a new url.
	Update(ctx context.Context, u *URL) (URL, error)
	// Delete the created url.
	Delete(ctx context.Context, u *URL) error
	// ByCode retrieves an url by its code.
	ByCode(ctx context.Context, u *URL) (URL, error)
}
