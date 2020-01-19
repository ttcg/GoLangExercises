package main

import (
	"fmt"
	"net/http"

	"github.com/ttcg/golangExercises/pluralsight/controllers"
)

func main() {

	port := "1983"

	fmt.Println("The webservice can be consumed at http://localhost:" + port)

	controllers.RegisterControllers()
	http.ListenAndServe(":"+port, nil)
}
