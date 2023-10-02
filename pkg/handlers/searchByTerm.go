package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strings"

	"github.com/datsfilipe/rinha-backend-go/pkg/models"
	"github.com/datsfilipe/rinha-backend-go/pkg/utils"
)

func SearchByTermHandler(db *sql.DB, t string) ([]byte, int, error) {
	if !utils.ValidSearchTerm(t) {
		return nil, 400, errors.New("Invalid search term")
	}

	people, err := db.Query("SELECT people, word_similarity($1, search) AS sml FROM people WHERE $1 <% search", t)
	if err != nil {
		return nil, 500, err
	}

	var peopleList []models.Person
	for people.Next() {
		var rawPerson string
		var sml float64

		err = people.Scan(&rawPerson, &sml)
		if err != nil {
			return nil, 500, err
		}

		rawPerson = strings.ReplaceAll(rawPerson, "(", "")
		rawPerson = strings.ReplaceAll(rawPerson, ")", "")

		rawPersonArray := strings.Split(rawPerson, ",")

		for i := 0; i < len(rawPersonArray); i++ {
			rawPersonArray[i] = strings.ReplaceAll(rawPersonArray[i], "\"", "")
		}

		stack := utils.DeserializeStringArray([]byte(rawPersonArray[4]))

		person := models.Person{
			ID:        rawPersonArray[0],
			Nick:      rawPersonArray[1],
			Name:      rawPersonArray[2],
			BirthDate: rawPersonArray[3],
			Stack:     stack,
		}

		peopleList = append(peopleList, person)
	}

	if len(peopleList) == 0 {
		return nil, 200, errors.New("No people found")
	}

	response, err := json.Marshal(peopleList)
	if err != nil {
		return nil, 500, err
	}

	return response, 200, nil
}
