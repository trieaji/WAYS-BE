package productsdto

//Declare CreateProductRequest struct here ...
type CreateProductRequest struct {
	Name  string `json:"name" gorm:"type: varchar(255)" validate:"required"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	Qty   int    `json:"qty" form:"qty" gorm:"type: int" validate:"required"`
	//CategoryID int    `json:"category_id" form:"category_id" gorm:"type: int" validate:"required"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)" validate:"required"`
}

//Declare UpdateProductRequest struct here ...
type UpdateProductRequest struct {
	Name  string `json:"name" form:"name"`
	Image string `json:"image" form:"image"`
}
