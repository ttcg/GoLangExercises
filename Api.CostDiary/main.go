package main

import (
	"fmt"
	"net/http"

	"github.com/ttcg/GoLangExercies/Api.CostDiary/controllers"
	"github.com/ttcg/GoLangExercies/Api.CostDiary/managers/costitemmgr"
	"github.com/ttcg/GoLangExercies/Api.CostDiary/managers/costtypemgr"
)

func main() {
	setUpDataSeeding()

	setUpWebServer()
}

func setUpDataSeeding() {
	costtypemgr.SeedCostTypesData()
	costitemmgr.SeedCostItemsData()
}

func setUpWebServer() {
	port := "3000"

	fmt.Println("The webservice can be consumed at http://localhost:" + port)

	controllers.RegisterControllers()
	http.ListenAndServe(":"+port, nil)
}
