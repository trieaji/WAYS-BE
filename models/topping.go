package models

type Topping struct {
	ID     int    `json:"id"`
	Name   string `json:"name" gorm:"type: varchar(255)"`
	Price  int    `json:"price" form:"price" gorm:"type: int"`
	Image  string `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID int    `json:"user_id" form:"user_id"`
}
