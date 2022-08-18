package cartsdto

type createCartRequest struct {
	ID        int   `json:"id"`
	UserID    int   `json:"user_id"`
	ProductID int   `json:"product_id"`
	ToppingID []int `json:"topping_id"`
	QTY       int   `json:"qty"`
	SubTotal  int   `json:"subtotal"`
}

type updateCartRequest struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	ToppingID int `json:"topping_id"`
}
