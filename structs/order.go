package structs

import (
	"time"
)

type Order struct {
	// gorm.Model
	Order_id      int       `gorm:" primary_key"`
	Customer_name string    `gorm:"type:varchar(100)"`
	Ordered_at    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"ordered_at"`
	Status_id     int
	Items         []Item `gorm:"foreignkey:Order_id"`
	//kalau mau automigrate di migrate dulu tanpa fk, baru setelah migrate masukkan fknya
}

// func (Orders) TableName() string {
//     return "orders"
// }
