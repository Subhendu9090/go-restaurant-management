package routes

import (
	"github.com/gorilla/mux"
)

func FoodRoutes(incomingRoutes *mux.Router) {
	incomingRoutes.HandleFunc("/foods", controllers.GetFoods()).Methods("GET")
	incomingRoutes.HandleFunc("/food/:food_id", controllers.GetFood()).Methods("GET")
	incomingRoutes.HandleFunc("/foods", conttrollers.CreateFood()).Methods("POST")
	incomingRoutes.HandleFunc("/foods/:food_id", controllers.UpdateFood()).Methods("PATCH")
}
