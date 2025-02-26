package routes

import (
	"golang-restaurant-management/controllers"

	"github.com/gorilla/mux"
)

func UserRoutes(inComingRoutes *mux.Router) {
	inComingRoutes.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	inComingRoutes.HandleFunc("/user/{user_id}", controllers.GetUser).Methods("GET")
	inComingRoutes.HandleFunc("/users/signUp", controllers.SignUp()).Methods("POST")
	inComingRoutes.HandleFunc("/users/login", controllers.Login()).Methods("POST")

}
