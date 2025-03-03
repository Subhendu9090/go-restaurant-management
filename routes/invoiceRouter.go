package routes

import (
	"golang-restaurant-management/controllers"

	"github.com/gorilla/mux"
)

func InvoiceRoutes(incomingRoutes *mux.Router) {
	incomingRoutes.HandleFunc("/invoice", controllers.CreateInvoice).Methods("POST")

	incomingRoutes.HandleFunc("/invoice", controllers.GetInvoices).Methods("GET")

	incomingRoutes.HandleFunc("/invoice/:invoice_id", controllers.GetInvoice).Methods("GET")

	incomingRoutes.HandleFunc("/invoice", controllers.UpdateInvoice).Methods("PATCH")
}
