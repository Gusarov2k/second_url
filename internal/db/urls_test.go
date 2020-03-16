package db_test

import (
	"context"
	"github.com/Gusarov2k/second_url"
	"github.com/Gusarov2k/second_url/internal/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestURLsRepo_Create(t *testing.T) {
	setUp(t)

	c := db.NewClient()
	if err := c.Open(PostgresTest); err != nil {
		t.Fatal(err)
	}
	defer func() { c.Close() }()

	u := shorten.URL{
		Code: "some_code",
		URL:  "http://example.org",
	}

	r := db.NewURLRepository(c)
	if err := r.Create(context.Background(), &u); err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 1, u.ID, "bad url id, expected 1, but got: ", u.ID)
}

func TestURLsRepo_ByCode(t *testing.T) {
	setUp(t)

	c := db.NewClient()
	if err := c.Open(PostgresTest); err != nil {
		t.Fatal(err)
	}
	defer func() { c.Close() }()

	u := shorten.URL{
		Code: "some_code",
		URL:  "http://example.org",
	}

	r := db.NewURLRepository(c)
	if err := r.Create(context.Background(), &u); err != nil {
		t.Fatal(err)
	}

	url, err := r.ByCode(context.Background(), &u)

	assert.Nil(t, err, "Err not nil")
	assert.EqualValues(t, url.Code, u.Code, "urls ids are not equal")
	assert.EqualValues(t, url.URL, u.URL, "urls are not equal")

}

func TestURLsRepo_Update(t *testing.T) {
	setUp(t)

	c := db.NewClient()
	if err := c.Open(PostgresTest); err != nil {
		t.Fatal(err)
	}
	defer func() { c.Close() }()

	u := shorten.URL{
		Code: "some_code",
		URL:  "http://example.org",
	}

	uSecond := shorten.URL{
		ID:   1,
		Code: "second",
		URL:  "http://example_second.ru",
	}

	r := db.NewURLRepository(c)

	if err := r.Create(context.Background(), &u); err != nil {
		t.Fatal(err)
	}

	url, err := r.Update(context.Background(), &uSecond)

	assert.Nil(t, err, "Err not nil")
	assert.NotEqual(t, url.URL, u.URL, "url equal updated url")
	assert.NotEqual(t, url.Code, u.Code, "code equal updated code")

}

func TestURLsRepo_Delete(t *testing.T) {
	setUp(t)

	c := db.NewClient()
	if err := c.Open(PostgresTest); err != nil {
		t.Fatal(err)
	}
	defer func() { c.Close() }()

	u := shorten.URL{
		ID: 1,
	}

	r := db.NewURLRepository(c)

	err := r.Delete(context.Background(), &u)

	assert.Nil(t, err, "Err not nil")

}
