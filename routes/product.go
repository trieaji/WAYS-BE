package routes

import (
	"ways/handlers"
	"ways/pkg/middleware"
	"ways/pkg/mysql"
	"ways/repositories"

	"github.com/gorilla/mux"
)

//Create ProductRoutes function here ...
func ProductRoutes(r *mux.Router) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository) //jika kontrak yang ada di repositories tidak terpenuhi akan error

	r.HandleFunc("/products", h.FindProduct).Methods("GET")
	r.HandleFunc("/product/{id}", middleware.Auth(h.GetProduct)).Methods("GET")
	//r.HandleFunc("/product", h.CreateProduct).Methods("POST")
	r.HandleFunc("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct))).Methods("POST")
	r.HandleFunc("/product/{id}", middleware.Auth(h.DeleteProduct)).Methods("DELETE")
}
