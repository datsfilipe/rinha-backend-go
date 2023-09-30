package handlers

import (
	"encoding/json"

	"github.com/datsfilipe/rinha-backend-go/pkg/database"
	"github.com/datsfilipe/rinha-backend-go/pkg/utils"
	"github.com/datsfilipe/rinha-backend-go/pkg/models"
)

func GetPersonHandler(id string) ([]byte, error) {
	if len(id) == 0 {
		return nil, nil
	}

	db, err := database.Open()
	if err != nil {
		return nil, err
	}

	people, err := db.Query("SELECT * FROM people WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return []byte("Error getting person"), err
	}

	var person models.Person
	for people.Next() {
		var id string
		var nick string
		var name string
		var birthDate string
		var stack []uint8

		err = people.Scan(&id, &nick, &name, &birthDate, &stack)
		if err != nil {
			return []byte("Error getting person 1"), err
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
		return []byte("Person not found"), nil
	}

	response, err := json.Marshal(person)
	if err != nil {
		return []byte("Error getting person"), err
	}

	return response, nil
}
