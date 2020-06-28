package main

import (
	"fmt"
	"net/http"

	"github.com/dsnatochy/webservice/controllers"
)

func print(v interface{}) {
	fmt.Println(v)
}
func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
