package models

type Person struct {
	ID			int			`json:"id"`
	Nick		string		`json:"apelido"`
	Name		string		`json:"nome"`
	BirthDate	string		`json:"nascimento"`
	Stack		[]string	`json:"stack"`
}
