package repositories

import (
	"ways/models"

	"gorm.io/gorm"
)

//Declare ToppingRepository interface here ... biasa kita anggap sebagai kontrak, kalau kontrak tidak dipanggil akan error di routes nya
type ToppingRepository interface { //untuk melakukan fetching data dari id. Untuk mengembalikan data tunggal
	FindTopping() ([]models.Topping, error)
	GetTopping(ID int) (models.Topping, error)
	CreateTopping(topping models.Topping) (models.Topping, error)
	// //UpdateTopping(topping models.Topping) (models.Topping, error)
	DeleteTopping(topping models.Topping) (models.Topping, error)
}

//Declare repository struct here ... yang berisikan database

//Create RepositoryProduct function here ...
func RepositoryTopping(db *gorm.DB) *repository {
	return &repository{db}
}

//Create FindTopping method here ...
func (r *repository) FindTopping() ([]models.Topping, error) {
	var toppings []models.Topping
	err := r.db.Find(&toppings).Error

	return toppings, err
}

//Create GetProduct method here ...
func (r *repository) GetTopping(ID int) (models.Topping, error) {
	var topping models.Topping
	err := r.db.First(&topping, ID).Error

	return topping, err
}

//Create CreateProduct method here ...
func (r *repository) CreateTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Create(&topping).Error

	return topping, err
}

//Create UpdateProduct method here ...
// func (r *repository) UpdateProduct(product models.Product) (models.Product, error) {
// err := r.db.Debug().Save(&product)Error
// return product, err
// }

//Create DeleteProduct method here ...
func (r *repository) DeleteTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Delete(&topping).Error

	return topping, err
}
