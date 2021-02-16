package structs

import (
	"github.com/jinzhu/gorm"
)

type Items struct {
	gorm.Model
	Item_id     int    `json:item_id`
	Item_code   int    `json:item_code`
	Description string `json:description`
	Quantity    int    `json:quantity`
	// Orders      Orders `gorm:"references:Order_id`
	Order_id int `json:order_id`
}
