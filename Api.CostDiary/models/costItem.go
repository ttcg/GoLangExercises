package models

import (
	"time"

	"github.com/google/uuid"
)

// CostItem : Cost detail item
type CostItem struct {
	ID         uuid.UUID
	ItemName   string
	CostTypeID uuid.UUID
	Amount     float32
	DateUsed   time.Time
}
