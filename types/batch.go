package types

import (
	"time"
)

type Batch struct {
	ID          int       `gorm:"primaryKey;column:id" json:"id"`
	ItemID      int       `gorm:"column:item_id" json:"item_id"`
	ReceiptDate time.Time `gorm:"type:date;column:receipt_date" json:"receipt_date"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
