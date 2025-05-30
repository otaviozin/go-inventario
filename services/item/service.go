package item

import (
	"github.com/otaviozin/go-inventory/database"
	"github.com/otaviozin/go-inventory/types"
)

func Create(newItem types.Item) (*types.Item, error) {
    result := database.DB.Create(&newItem)
    return &newItem, result.Error
}