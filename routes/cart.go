package routes

import (
	"ways/handlers"
	"ways/pkg/middleware"
	"ways/pkg/mysql"
	"ways/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/carts", h.FindCart).Methods("GET")
	r.HandleFunc("/cart/{id}", h.GetCart).Methods("GET")
	//r.HandleFunc("/carts-id", h.FindCartsByID).Methods("GET")
	r.HandleFunc("/cart", middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.UpdateCart)).Methods("PATCH")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.DeleteCart)).Methods("DELETE")
}
