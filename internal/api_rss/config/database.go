package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func StartDB() (*sql.DB, error) {

	dbUrl := os.Getenv("DB_URL")
	fmt.Println(dbUrl)
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
