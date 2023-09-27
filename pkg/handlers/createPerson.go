package handlers

// import (
// 	"encoding/json"
// 
// 	"github.com/datsfilipe/rinha-backend-go/pkg/database"
// 	"github.com/datsfilipe/rinha-backend-go/pkg/models"
// )

func CreatePersonHandler(request []byte) ([]byte, error) {
	if len(request) == 0 {
		return nil, nil
	} else {
		return request, nil
	}

	// var person models.Person
	// err := json.Unmarshal(request, &person)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// db, err := database.Open()
	//
	// db.Exec(
	// 	"INSERT INTO people (nick, name, birth_date, stack) VALUES ($1, $2, $3, $4)",
	// 	person.Nick, person.Name, person.BirthDate, person.Stack,
	// )
	//
	// return []byte("Person created"), nil
}
