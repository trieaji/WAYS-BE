package models

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id" form:"user_id"`
	Amount    int       `json:"amount" form:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
