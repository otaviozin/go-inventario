package types

import (
	"time"
)

type Item struct {
	ID        int       `gorm:"primaryKey;column:id"         json:"id"`
	Name      string    `gorm:"column:name"              json:"name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
