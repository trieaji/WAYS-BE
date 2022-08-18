package routes

import (
	"ways/handlers"
	"ways/pkg/middleware"
	"ways/pkg/mysql"
	"ways/repositories"

	"github.com/gorilla/mux"
)

//Create ProductRoutes function here ...
func ToppingRoutes(r *mux.Router) {
	toppingRepository := repositories.RepositoryTopping(mysql.DB)
	h := handlers.HandlerTopping(toppingRepository) //jika kontrak yang ada di repositories tidak terpenuhi akan error

	r.HandleFunc("/toppings", h.FindTopping).Methods("GET")
	r.HandleFunc("/topping/{id}", middleware.Auth(h.GetTopping)).Methods("GET")
	//r.HandleFunc("/product", h.CreateProduct).Methods("POST")
	r.HandleFunc("/topping", middleware.Auth(middleware.UploadFile(h.CreateTopping))).Methods("POST")
	r.HandleFunc("/topping/{id}", middleware.Auth(h.DeleteTopping)).Methods("DELETE")
}
