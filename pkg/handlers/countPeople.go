package handlers

import (
	"fmt"

	"github.com/datsfilipe/rinha-backend-go/pkg/database"
)

func CountPeopleHandler() ([]byte, int, error) {
	db, err := database.Open()
	if err != nil {
		return nil, 500, err
	}

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
