package main

//import package
import (
	//jangan lupa import untuk json
	"fmt"
	"net/http"
	"ways/database"
	"ways/pkg/mysql"
	"ways/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//Main function for Declare Route
func main() {

	// env
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	//path file
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads")))) // add this code. Kegunaannya untuk bisa mememasukkan atau membaca sebuah folder yang ada di dalam folder uploads

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
