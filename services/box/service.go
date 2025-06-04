package box

import (
	"github.com/otaviozin/go-inventory/database"
	"github.com/otaviozin/go-inventory/types"
)

func Create(boxes []types.Box) error {
	return database.DB.Create(&boxes).Error
}
