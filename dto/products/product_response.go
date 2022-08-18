package productsdto

//Declare ProductResponse struct here ...
type ProductResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}
