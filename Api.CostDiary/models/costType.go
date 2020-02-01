package models

import (
	"github.com/google/uuid"
)

// CostType : class
type CostType struct {
	ID           uuid.UUID
	CostTypeName string
}
