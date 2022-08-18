package toppingsdto

//Declare ProductResponse struct here ...
type ToppingResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}
