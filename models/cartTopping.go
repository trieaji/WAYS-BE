package models

import "time"

type CartTopping struct {
	ID        int       `json:"id"`
	CartID    int       `json:"cartId" form:"cartId"`
	ToppingID int       `json:"toppingId" form:"toppingId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
