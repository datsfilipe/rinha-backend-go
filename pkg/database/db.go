package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Open() (*sql.DB, error) {
	connsStr := "postgres://rinha:rinha@db/rinha?sslmode=disable"
	db, err := sql.Open("postgres", connsStr)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(450)
	db.SetMaxIdleConns(50)

	return db, err
}
