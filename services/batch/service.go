package batch

import (
	"github.com/otaviozin/go-inventory/database"
	"github.com/otaviozin/go-inventory/types"
)

func Create(newBatch types.Batch) (*types.Batch, error) {
	result := database.DB.Create(&newBatch)
	return &newBatch, result.Error
}
