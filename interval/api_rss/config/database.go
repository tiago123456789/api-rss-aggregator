package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func StartDB() (*sql.DB, error) {

	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
