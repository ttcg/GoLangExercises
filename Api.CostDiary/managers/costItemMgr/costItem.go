package costitemmgr

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ttcg/GoLangExercies/Api.CostDiary/models"
)

var (
	costItems []*models.CostItem
)

// SeedCostItemsData : data seeding for Cost Items
func SeedCostItemsData() {
	const TwFood = "e401b4e5-b3c6-4855-8846-9f83005391a7"
	const TwOthers = "7a9e65c0-2a68-49c9-9038-1c5e033908a8"
	const AP = "7d7483a0-73d4-421c-821f-1ac8f68ad579"
	const Others = "55337e4b-7392-489e-a633-17bc7cf8e9db"
	const Groceries = "a10b83aa-795a-460a-b0a1-0b051871f46c"
	const Diesel = "4c5cedd5-c144-4225-92ff-2e472ed6c274"
	const NA = "30f4ecf3-fd76-4022-bc1b-598a7cdb6179"

	costItems = append(costItems, createCostItem("Sainsbury", Groceries, createDate(2020, time.January, 2), 19.24))
	costItems = append(costItems, createCostItem("Boots", Groceries, createDate(2020, time.January, 5), 1.99))
	costItems = append(costItems, createCostItem("Sainsbury", Groceries, createDate(2020, time.January, 5), 68.7))
	costItems = append(costItems, createCostItem("Asda", Groceries, createDate(2020, time.January, 5), 85.76))
	costItems = append(costItems, createCostItem("London Tickets", NA, createDate(2020, time.January, 2), 11.4))
	costItems = append(costItems, createCostItem("Diesel", Diesel, createDate(2020, time.January, 7), 20))
	costItems = append(costItems, createCostItem("AP Cash", AP, createDate(2020, time.January, 30), 50))
	costItems = append(costItems, createCostItem("Samsonite Luggage", Others, createDate(2020, time.January, 1), 100.75))
}

func createCostItem(itemName string, costTypeID string, dateUsed time.Time, amount float32) *models.CostItem {
	return &models.CostItem{
		ID:         uuid.New(),
		ItemName:   itemName,
		CostTypeID: uuid.MustParse(costTypeID),
		DateUsed:   dateUsed,
		Amount:     amount}
}

func createDate(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

// GetCostItems : to return all cost items
func GetCostItems() []*models.CostItem {
	return costItems
}

// GetCostItemByID : to return cost item by given ID
func GetCostItemByID(id uuid.UUID) (models.CostItem, error) {
	for _, v := range costItems {
		if v.ID == id {
			return *v, nil
		}
	}

	return models.CostItem{}, fmt.Errorf("The given ID %s not found", id)
}

// AddCostItem : to add Cost Item
func AddCostItem(costItem models.CostItem) (models.CostItem, error) {
	if costItem.ID == uuid.Nil {
		return models.CostItem{}, errors.New("New Cost Item must have an ID.  Please provide")
	}

	costItems = append(costItems, &costItem)
	return costItem, nil
}
