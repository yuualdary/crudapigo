package structs

type Status struct {
	// gorm.Model
	//GORM defined a gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Status_id  int `gorm:"primary_key"`
	Statusname string
	// Orders     []Order `gorm:"foreignkey:Status_id"`
	// Order    []Orders `gorm:"foreignKey:o_id;association_foreignkey:Order_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
