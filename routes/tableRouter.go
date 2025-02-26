package routes

import "github.com/gorilla/mux"

func TableRoutes(incomingRoutes *mux.Router) {
	incomingRoutes.HandleFunc("/tables", controllers.GetTables()).Methods("GET")

	incomingRoutes.HandleFunc("/tables/:table_id", controllers.GetTables()).Methods("GET")

	incomingRoutes.HandleFunc("/tables", controllers.CreateTables()).Methods("POST")

	incomingRoutes.HandleFunc("/tables", controllers.UpdateTables()).Methods("PATCH")

}
