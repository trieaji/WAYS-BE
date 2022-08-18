package toppingsdto

//Declare CreateProductRequest struct here ...
type CreateToppingRequest struct {
	Name  string `json:"name" gorm:"type: varchar(255)" validate:"required"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	//CategoryID int    `json:"category_id" form:"category_id" gorm:"type: int" validate:"required"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)" validate:"required"`
}

//Declare UpdateProductRequest struct here ...
type UpdateToppingRequest struct {
	Name  string `json:"name" form:"name"`
	Image string `json:"image" form:"image"`
}
