package handlers

import (
	"fmt"

	"github.com/datsfilipe/rinha-backend-go/pkg/database"
)

func CountPeopleHandler() ([]byte, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}

	people, err := db.Query("SELECT * FROM people")
	if err != nil {
		return []byte("Error getting people"), err
	}

	var count int
	for people.Next() {
		count++
	}

	return []byte(fmt.Sprint(count)), nil
}
