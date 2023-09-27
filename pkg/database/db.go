package database

import (
	"database/sql"
	"log"
)

func Open() (*sql.DB, error) {
	connsStr := "postgres://rinha:rinha@db/rinha?sslmode=disable"
	db, err := sql.Open("postgres", connsStr)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
