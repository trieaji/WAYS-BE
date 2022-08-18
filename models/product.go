package models

import "time"

//Product model struct
type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	Desc      string    `json:"desc" gorm:"type:text" form:"desc"`
	Price     int       `json:"price" form:"price" gorm:"type: int"`
	Image     string    `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty       int       `json:"qty" form:"qty"`
	UserID    int       `json:"user_id" form:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
