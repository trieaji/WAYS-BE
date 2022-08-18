package routes

// Import gorilla/mux package here ...
import "github.com/gorilla/mux"

// Create RouteInit function and Call TodoRoutes function here ...
func RouteInit(r *mux.Router) {
	ProductRoutes(r)
	ToppingRoutes(r)
	CartRoutes(r)
	TransactionRoutes(r)
}
