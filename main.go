package main

import (
	"golang-restaurant-management/middleware"
	"golang-restaurant-management/routes"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var foodCollection *mongoCollection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := mux.NewRouter()
	routes.UserRoutes(router)

	authRouter := router.PathPrefix("/").Subrouter()
	authRouter.Use(middleware.Authenticate())

	routes.FoodRoutes(authRouter)
	routes.MenuRoutes(authRouter)
	routes.TableRoutes(authRouter)
	routes.OrderRoutes(authRouter)
	routes.OrderItemRoutes(authRouter)
	routes.InvoiceRoutes(authRouter)

	http.ListenAndServe(":"+port, router)

}
