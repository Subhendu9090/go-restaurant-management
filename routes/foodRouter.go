package routes

import (
	"golang-restaurant-management/controllers"

	"github.com/gorilla/mux"
)

func FoodRoutes(incomingRoutes *mux.Router) {
	incomingRoutes.HandleFunc("/foods", controllers.GetFoods).Methods("GET")
	incomingRoutes.HandleFunc("/food/:food_id", controllers.GetFood).Methods("GET")
	incomingRoutes.HandleFunc("/foods", controllers.CreateFood).Methods("POST")
	incomingRoutes.HandleFunc("/foods/:food_id", controllers.UpdateFoods).Methods("PATCH")
}
