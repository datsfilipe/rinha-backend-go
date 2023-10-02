package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/datsfilipe/rinha-backend-go/pkg/models"
	"github.com/datsfilipe/rinha-backend-go/pkg/utils"
	"github.com/google/uuid"
)

func CreatePersonHandler(db *sql.DB, request []byte) ([]byte, int, error, string) {
	if len(request) == 0 {
		return nil, 400, errors.New("Invalid request"), ""
	}

	var person models.Person
	err := json.Unmarshal(request, &person)
	if err != nil {
		return nil, 422, err, ""
	}

	if !utils.IsValidDate(person.BirthDate) {
		return nil, 422, errors.New("Invalid date"), ""
	}

	if !utils.VerifyLength(person.Nick, 1, 32) {
		return nil, 422, errors.New("Invalid nick"), ""
	}

	if !utils.VerifyLength(person.Name, 1, 100) {
		return nil, 422, errors.New("Invalid name"), ""
	}

	if len(person.Stack) > 0 {
		for _, item := range person.Stack {
			if !utils.VerifyLength(item, 1, 32) {
				return nil, 422, errors.New("Invalid stack"), ""
			}
		}
	}

	repeatedNick, err := db.Query("SELECT nick FROM people WHERE nick = $1", person.Nick)
	if err != nil {
		return nil, 500, err, ""
	}

	if repeatedNick.Next() {
		return nil, 422, errors.New("Nick already in use"), ""
	}

	var _, err2 = db.Exec(
		"INSERT INTO people (id, nick, name, birth_date, stack) VALUES ($1, $2, $3, $4, $5)",
		uuid.New().String(), person.Nick, person.Name, person.BirthDate, utils.SerializeStringArray(person.Stack),
	)
	if err2 != nil {
		return nil, 500, err2, ""
	}

	var personID string
	err3 := db.QueryRow("SELECT id FROM people WHERE nick = $1", person.Nick).Scan(&personID)
	if err3 != nil {
		return nil, 500, err3, ""
	}

	var location = fmt.Sprintf("/pessoas/%s", personID)

	return []byte("Person created"), 201, nil, location
}
