package models

import "time"

type Cart struct {
	ID        int       `json:"id"`
	ProductId int       `json:"productId" form:"productId"`
	TransId   int       `json:"transId" form:"transId"`
	Qty       int       `json:"qty" form:"qty"`
	SubAmount int       `json:"subAmount" form:"subAmount"`
	UserID    int       `json:"user_id" form:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
// func convertResponseCart(u models.Cart) models.Cart {
// 	return models.Cart{
// 		ID:       u.ID,
// 		QTY:      u.QTY,
// 		SubTotal: u.SubTotal,
// 		Product:  u.Product,
// 		Topping:  u.Topping,
// 	}

