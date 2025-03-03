package routes

import (
	"golang-restaurant-management/controllers"

	"github.com/gorilla/mux"
)

func TableRoutes(incomingRoutes *mux.Router) {
	incomingRoutes.HandleFunc("/tables", controllers.GetTables).Methods("GET")

	incomingRoutes.HandleFunc("/tables/:table_id", controllers.GetTables).Methods("GET")

	incomingRoutes.HandleFunc("/tables", controllers.CreateTable).Methods("POST")

	incomingRoutes.HandleFunc("/tables", controllers.UpdateTable).Methods("PATCH")

}
