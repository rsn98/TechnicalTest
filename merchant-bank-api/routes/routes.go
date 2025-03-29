package routes

import (
	"github.com/gorilla/mux"
	"github.com/rsn98/merchant-bank-api/controllers"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/payment", controllers.Payment).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("POST")
	return r
}
