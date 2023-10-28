package config

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func StartDB() (*sql.DB, error) {

	dbUrl := os.Getenv("DB_URL")
	dbUrl = strings.TrimSpace(dbUrl)
	fmt.Println(dbUrl)
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
