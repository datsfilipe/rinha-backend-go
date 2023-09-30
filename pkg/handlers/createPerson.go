package handlers

import (
	"encoding/json"

	"github.com/datsfilipe/rinha-backend-go/pkg/database"
	"github.com/datsfilipe/rinha-backend-go/pkg/models"
	"github.com/datsfilipe/rinha-backend-go/pkg/utils"
	"github.com/google/uuid"
)

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

	var _, err2 = db.Exec(
		"INSERT INTO people (id, nick, name, birth_date, stack) VALUES ($1, $2, $3, $4, $5)",
		uuid.New().String(), person.Nick, person.Name, person.BirthDate, utils.SerializeStringArray(person.Stack),
	)

	if err2 != nil {
		return nil, err2
	}

	return []byte("Person created"), nil
}
