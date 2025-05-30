package item

import (
	"github.com/otaviozin/go-inventory/database"
	"github.com/otaviozin/go-inventory/types"
)

func Create(newItem types.Item) (*types.Item, error) {
	result := database.DB.Create(&newItem)
	return &newItem, result.Error
}
func GetById(id int) (*types.Item, error) {
	var item types.Item
	result := database.DB.First(&item, id)
	return &item, result.Error
}
