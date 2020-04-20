package db

import (
	"github.com/go-kit/kit/log"
)

// NewClient returns a new Client backed by Postgres.
func NewClient(options ...ConfigOption) *Client {
	c := Client{
		logger:         log.NewNopLogger(),
		maxConnections: defaultMaxConnections,
	}
	for _, opt := range options {
		opt(&c)
	}

	return &c
}

// ConfigOption configures the client.
type ConfigOption func(*Client)

// WithLogger configures a logger to debug interactions with Postgres.
func WithLogger(l log.Logger) ConfigOption {
	return func(c *Client) {
		c.logger = l
	}
}

// WithMaxConnections configures maximum number of opened and idle connections.
func WithMaxConnections(n int) ConfigOption {
	return func(c *Client) {
		c.maxConnections = n
	}
}
