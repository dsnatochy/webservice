package models

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1 // compiler will infer that this is an integer
)

func main() {
	fmt.Println("from models")
}
