package main

import "fmt"

func main() {
	input := `
type User('users') {
    id: ID!
    name: string
    ids: array<int>
    username: string
    someInt: int
    someFloat: float
    someBool: bool
    someDate: date
    someDateTime: datetime
}`

	tokens := lex(input)
	parser := NewParser(tokens)
	typeDef, err := parser.Parse()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Parsed type: %+v\n", typeDef)
}
