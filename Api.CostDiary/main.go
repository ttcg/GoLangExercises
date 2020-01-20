package main

import (
	"fmt"

	"github.com/ttcg/GoLangExercies/Api.CostDiary/models"
)

func main() {
	fmt.Println("hello world")

	newCostType := models.CostType{123, "thet"}

	fmt.Println(newCostType)
}
