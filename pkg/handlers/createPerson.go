package handlers

import (
	"encoding/json"

	"github.com/datsfilipe/rinha-backend-go/pkg/database"
	"github.com/datsfilipe/rinha-backend-go/pkg/models"
	"github.com/google/uuid"
)

func serializeStringArray(array []string) any {
	if len(array) == 0 {
		return nil
	}

	var serialized string
	for i, item := range array {
		if i == 0 {
			serialized = "{" + item
		} else {
			serialized = serialized + "," + item
		}

		if i == len(array)-1 {
			serialized = serialized + "}"
		}
	}
	return serialized
}

func CreatePersonHandler(request []byte) ([]byte, error) {
	if len(request) == 0 {
		return nil, nil
	}

	var person models.Person
	err := json.Unmarshal(request, &person)

	if err != nil {
		return nil, err
	}

	db, err := database.Open()

	if err != nil {
		return nil, err
	}

	repeatedNick, err := db.Query("SELECT nick FROM people WHERE nick = $1", person.Nick)

	if err != nil {
		return nil, err
	}

	if repeatedNick.Next() {
		return []byte("Nick already in use"), nil
	}

	people, err := db.Query("SELECT * FROM people")
	if err == nil {
		for people.Next() {
			var person models.Person
			err := people.Scan(&person.Nick, &person.Name, &person.BirthDate, &person.Stack)
			if err != nil {
				return nil, err
			}
		}
	}

	var _, err2 = db.Exec(
		"INSERT INTO people (id, nick, name, birth_date, stack) VALUES ($1, $2, $3, $4, $5)",
		uuid.New().String(), person.Nick, person.Name, person.BirthDate, serializeStringArray(person.Stack),
	)

	if err2 != nil {
		return nil, err2
	}

	return []byte("Person created"), nil
}
