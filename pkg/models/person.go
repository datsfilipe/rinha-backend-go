package models

type Person struct {
	ID			string		`json:"id"`
	Nick		string		`json:"apelido"`
	Name		string		`json:"nome"`
	BirthDate	string		`json:"nascimento"`
	Stack		[]string	`json:"stack"`
}
