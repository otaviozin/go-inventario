package types

import (
	"time"
)

type Box struct {
	ID          int       `gorm:"primaryKey;column:id"         json:"id"`
	BatchID     int       `gorm:"column:batch_id"              json:"batch_id"`
	BoxBatchID  int       `gorm:"column:box_batch_id"          json:"box_batch_id"`
	ItemsAmount int       `gorm:"column:items_amount"          json:"items_amount"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
