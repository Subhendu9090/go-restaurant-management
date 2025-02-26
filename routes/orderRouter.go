package routes

import "github.com/gorilla/mux"

func OrderRoutes(incomingRoutes *mux.Router) {
	incomingRoutes.HandleFunc("/orders", controllers.GetOrders()).Methods("GET")

	incomingRoutes.HandleFunc("/orders/:order_id", controllers.GetOrders()).Methods("GET")

	incomingRoutes.HandleFunc("/orders", controllers.CreateOrders()).Methods("POST")

	incomingRoutes.HandleFunc("/orders", controllers.UpdateOrders()).Methods("PATCH")

}
