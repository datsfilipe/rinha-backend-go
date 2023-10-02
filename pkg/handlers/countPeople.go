package handlers

import (
	"database/sql"
	"fmt"
)

func CountPeopleHandler(db *sql.DB) ([]byte, int, error) {
	people, err := db.Query("SELECT * FROM people")
	if err != nil {
		return nil, 500, err
	}

	var count int
	for people.Next() {
		count++
	}

	return []byte(fmt.Sprint(count)), 200, nil
}
