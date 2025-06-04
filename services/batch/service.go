package batch

import (
	"github.com/otaviozin/go-inventory/database"
	"github.com/otaviozin/go-inventory/types"
	"gorm.io/gorm"
)

func Create(newBatch types.Batch) (*types.Batch, error) {
	result := database.DB.Create(&newBatch)
	return &newBatch, result.Error
}

func GetById(id int) (*types.Batch, error) {
	var batch types.Batch
	result := database.DB.First(&batch, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &batch, nil
}
