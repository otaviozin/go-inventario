package types

import (
	"time"
)

type Box struct {
	ID          int       `gorm:"primaryKey;column:id"         json:"id"`
	BatchID     int       `gorm:"column:batch_id"              json:"batch_id"`
	BoxBatchNum int       `gorm:"column:box_batch_num"          json:"box_batch_num"`
	ItemsAmount int       `gorm:"column:items_amount"          json:"items_amount"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
