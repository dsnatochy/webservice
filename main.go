package main

import (
	"fmt"

	"github.com/dsnatochy/webservice/models"
)

func print(v interface{}) {
	fmt.Println(v)
}
func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	print(u)
}
