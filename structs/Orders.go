package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Orders struct {
	gorm.Model
	Order_id      int       `json:"order_id"`
	Customer_name string    `json:"costumer_name"`
	ordered_at    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"ordered_at"`
}
