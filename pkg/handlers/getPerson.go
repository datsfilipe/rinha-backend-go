package handlers

import (
	"encoding/json"
	"errors"

	"github.com/datsfilipe/rinha-backend-go/pkg/database"
	"github.com/datsfilipe/rinha-backend-go/pkg/models"
	"github.com/datsfilipe/rinha-backend-go/pkg/utils"
)

func GetPersonHandler(id string) ([]byte, int, error) {
	if len(id) == 0 {
		return nil, 400, errors.New("Invalid ID")
	}

	db, err := database.Open()
	if err != nil {
		return nil, 500, err
	}

	people, err := db.Query("SELECT * FROM people WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return nil, 404, err
	}

	var person models.Person
	for people.Next() {
		var id string
		var nick string
		var name string
		var birthDate string
		var stack []uint8
		var search string

		err = people.Scan(&id, &nick, &name, &birthDate, &stack, &search)
		if err != nil {
			return nil, 500, err
		}

		person = models.Person{
			ID:        id,
			Nick:      nick,
			Name:      name,
			BirthDate: birthDate,
			Stack:     utils.DeserializeStringArray(stack),
		}
	}

	if person.ID == "" {
		return nil, 404, errors.New("Person not found")
	}

	response, err := json.Marshal(person)
	if err != nil {
		return nil, 500, err
	}

	return response, 200, nil
}
