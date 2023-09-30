package handlers

import (
	"encoding/json"
	"errors"

	"github.com/datsfilipe/rinha-backend-go/pkg/database"
	"github.com/datsfilipe/rinha-backend-go/pkg/models"
	"github.com/datsfilipe/rinha-backend-go/pkg/utils"
	"github.com/google/uuid"
)

func CreatePersonHandler(request []byte) ([]byte, int, error) {
	if len(request) == 0 {
		return nil, 400, errors.New("Invalid request")
	}

	var person models.Person
	err := json.Unmarshal(request, &person)
	if err != nil {
		return nil, 422, err
	}

	db, err := database.Open()
	if err != nil {
		return nil, 500, err
	}

	repeatedNick, err := db.Query("SELECT nick FROM people WHERE nick = $1", person.Nick)
	if err != nil {
		return nil, 500, err
	}

	if repeatedNick.Next() {
		return nil, 422, errors.New("Nick already in use")
	}

	var _, err2 = db.Exec(
		"INSERT INTO people (id, nick, name, birth_date, stack) VALUES ($1, $2, $3, $4, $5)",
		uuid.New().String(), person.Nick, person.Name, person.BirthDate, utils.SerializeStringArray(person.Stack),
	)
	if err2 != nil {
		return nil, 500, err2
	}

	return []byte("Person created"), 200, nil
}
