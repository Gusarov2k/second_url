package db_test

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"testing"

	"github.com/Gusarov2k/second_url/internal/db"
)

var (
	PostgresHost           = getEnv("POSTGRES_HOST", "localhost")
	PostgresPort           = getEnv("POSTGRES_PORT", "5432")
	PostgresDB             = getEnv("POSTGRES_DB", "short_link_development")
	PostgresDBTest         = getEnv("POSTGRES_DB_TEST", "short_link_test")
	PostgresUser           = getEnv("POSTGRES_USER", "ivan")
	PostgresPassword       = getEnv("POSTGRES_PASSWORD", "1234")
	PostgresConnectTimeout = getEnv("POSTGRES_CONNECT_TIMEOUT", "3")

	PostgresSys = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s connect_timeout=%s sslmode=disable",
		PostgresUser, PostgresPassword, PostgresHost, PostgresPort, PostgresDB, PostgresConnectTimeout)

	PostgresTest = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s connect_timeout=%s sslmode=disable",
		PostgresUser, PostgresPassword, PostgresHost, PostgresPort, PostgresDBTest, PostgresConnectTimeout)
)

func setUp(t *testing.T) {
	t.Helper()

	clearSQLDb(t)
}

func clearSQLDb(t *testing.T) {
	t.Helper()
	var err error

	pool, err := sqlx.Open("postgres", PostgresSys)
	if err != nil {
		t.Fatal("can't connect to db")
	}
	defer func() { _ = pool.Close() }()

	_, err = pool.Exec("DROP DATABASE IF EXISTS " + PostgresDBTest)
	if err != nil {
		t.Fatal(err)
	}

	_, err = pool.Exec("CREATE DATABASE " + PostgresDBTest)
	if err != nil {
		t.Fatal(err)
	}

	// Create schema
	c := db.NewClient()
	if err = c.Open(PostgresTest); err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	if err = c.InitSchema(); err != nil {
		t.Fatal(err)
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}
