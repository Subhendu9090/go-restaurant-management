package routes

import "github.com/gorilla/mux"

func MenuRoutes(incomingRoutes *mux.Router) {
	incomingRoutes.HandleFunc("/menu", controllers.GetMenus()).Methods("GET")

	incomingRoutes.HandleFunc("/menu/:menu_id", controllers.GetMenu()).Methods("GET")

	incomingRoutes.HandleFunc("/menu", controllers.CreateMenu()).Methods("POST")

	incomingRoutes.HandleFunc("/menu", controllers.UpdateMenu()).Methods("PATCH")

}
