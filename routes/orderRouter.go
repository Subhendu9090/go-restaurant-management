package routes

import (
	"golang-restaurant-management/controllers"

	"github.com/gorilla/mux"
)

func OrderRoutes(incomingRoutes *mux.Router) {
	incomingRoutes.HandleFunc("/orders", controllers.GetOrders).Methods("GET")

	incomingRoutes.HandleFunc("/orders/:order_id", controllers.GetOrder).Methods("GET")

	incomingRoutes.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")

	incomingRoutes.HandleFunc("/orders", controllers.UpdateOrder).Methods("PATCH")

}
