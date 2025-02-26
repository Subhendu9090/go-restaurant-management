package routes

import "github.com/gorilla/mux"

func OrderItemRoutes(incomingRoutes *mux.Router) {
	incomingRoutes.HandleFunc("/orderItems", controllers.GetOrderItems()).Methods("GET")

	incomingRoutes.HandleFunc("/orderItems/:orderItem_id", controllers.GetOrderItem()).Methods("GET")

	incomingRoutes.HandleFunc("/orderItems", controllers.CreateOrderItem()).Methods("POST")

	incomingRoutes.HandleFunc("/orderItems:/orderItem_id", controllers.UpdateOrderItems()).Methods("PATCH")

	incomingRoutes.HandleFunc("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())

}
