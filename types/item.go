package types

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ID   int    `gorm:"primaryKey;column:id"         json:"id"`
	Name string `gorm:"unique;column:name"              json:"name"`
}
