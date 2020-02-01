package costtypemgr

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ttcg/GoLangExercies/Api.CostDiary/models"
)

var (
	costTypes []*models.CostType
)

// SeedCostTypesData : data seeding for Cost Types
func SeedCostTypesData() {
	costTypes = append(costTypes, createCostType("e401b4e5-b3c6-4855-8846-9f83005391a7", "TW Food"))
	costTypes = append(costTypes, createCostType("7a9e65c0-2a68-49c9-9038-1c5e033908a8", "TW Others"))
	costTypes = append(costTypes, createCostType("7d7483a0-73d4-421c-821f-1ac8f68ad579", "AP"))
	costTypes = append(costTypes, createCostType("55337e4b-7392-489e-a633-17bc7cf8e9db", "Others"))
	costTypes = append(costTypes, createCostType("a10b83aa-795a-460a-b0a1-0b051871f46c", "Groceries"))
	costTypes = append(costTypes, createCostType("4c5cedd5-c144-4225-92ff-2e472ed6c274", "Diesel"))
	costTypes = append(costTypes, createCostType("30f4ecf3-fd76-4022-bc1b-598a7cdb6179", "NA"))
}

func createCostType(id string, name string) *models.CostType {
	return &models.CostType{
		ID:           uuid.MustParse(id),
		CostTypeName: name}
}

// GetCostTypes : to return all cost types
func GetCostTypes() []*models.CostType {
	return costTypes
}

// GetCostTypeByID : to return cost type by given ID
func GetCostTypeByID(id uuid.UUID) (models.CostType, error) {
	for _, v := range costTypes {
		if v.ID == id {
			return *v, nil
		}
	}

	return models.CostType{}, fmt.Errorf("The given ID %s not found", id)
}
