package repositories

import (
	"ways/models"

	"gorm.io/gorm"
)

//Declare ProductRepository interface here ... biasa kita anggap sebagai kontrak, kalau kontrak tidak dipanggil akan error di routes nya
type ProductRepository interface { //untuk melakukan fetching data dari id. Untuk mengembalikan data tunggal
	FindProduct() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	//UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product) (models.Product, error)
}

//Declare repository struct here ... yang berisikan database


//Create RepositoryProduct function here ...
func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

//Create FindProduct method here ...
func (r *repository) FindProduct() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error

	return products, err
}

//Create GetProduct method here ...
func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, ID).Error

	return product, err
}

//Create CreateProduct method here ...
func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}

//Create UpdateProduct method here ...
// func (r *repository) UpdateProduct(product models.Product) (models.Product, error) {
// err := r.db.Debug().Save(&product)Error
// return product, err
// }

//Create DeleteProduct method here ...
func (r *repository) DeleteProduct(product models.Product) (models.Product, error) {
	err := r.db.Delete(&product).Error

	return product, err
}
