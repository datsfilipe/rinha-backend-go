package handlers

import (
	"encoding/json"
	"strings"

	"github.com/datsfilipe/rinha-backend-go/pkg/database"
	"github.com/datsfilipe/rinha-backend-go/pkg/models"
	"github.com/datsfilipe/rinha-backend-go/pkg/utils"
)

func SearchByTermHandler(t string) ([]byte, error) {
	if len(t) == 0 {
		return nil, nil
	}

	db, err := database.Open()
	if err != nil {
		return nil, err
	}

	people, err := db.Query("SELECT people, word_similarity($1, search) AS sml FROM people WHERE $1 <% search", t)
	if err != nil {
		return []byte("Error getting people"), err
	}

	var peopleList []models.Person
	for people.Next() {
		var rawPerson string
		var sml float64

		err = people.Scan(&rawPerson, &sml)
		if err != nil {
			return []byte("Error getting people 1"), err
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
		return []byte("[]"), nil // json.Marshal with an empty array returns "null" somehow
	}

	response, err := json.Marshal(peopleList)
	if err != nil {
		return []byte("Error getting people"), err
	}

	return response, nil
}
