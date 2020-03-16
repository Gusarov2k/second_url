package db

import (
	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	defaultMaxConnections = 5
)

type Client struct {
	db             *sqlx.DB
	logger         log.Logger
	maxConnections int
}

// Open connection to PostgreSQL.
func (c *Client) Open(dataSourceName string) error {
	var err error

	c.logger.Log("level", "debug", "msg", "connecting to db")
	if c.db, err = sqlx.Open("postgres", dataSourceName); err != nil {
		c.logger.Log("level", "debug", "msg", "test")
		return err
	}
	if err = c.db.Ping(); err != nil {
		return err
	}
	c.db.SetMaxOpenConns(c.maxConnections)
	c.db.SetMaxIdleConns(c.maxConnections)
	c.logger.Log("level", "debug", "msg", "connected to db")

	return nil
}

// Close closes PostgreSQL connection.
func (c *Client) Close() error {
	return c.db.Close()
}

// InitSchema sets up the initial schema.
func (c *Client) InitSchema() error {
	_, err := c.db.Exec(Schema)
	return err
}

// Queries
const (
	urlInsert = `
INSERT INTO urls (code, original_url)
  VALUES (:code, :original_url)
RETURNING urls.id`
	urlByCode = `
SELECT id, code, original_url
FROM urls
WHERE code = :code`
	urlUpdate = `
UPDATE urls
SET code = :code, original_url = :original_url
WHERE id = :id`
	urlDelete = `
DELETE FROM urls
WHERE id = :id`
)
