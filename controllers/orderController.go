package controllers

import (
	"context"
	"encoding/json"
	"golang-restaurant-management/database"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.GetCollectionName("order")

func GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	var orders []bson.M
	defer cancel()
	result, err := orderCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error in finding Orders", http.StatusNotFound)
		return
	}
	err = result.All(ctx, &orders)
	if err != nil {
		http.Error(w, "Error in Decoding Orders", http.StatusForbidden)
		return
	}
	res, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, "Marshalling error", http.StatusForbidden)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetOrder(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	params := mux.Vars(r)
	orderId := params["order_id"]
	if orderId == "" {
		http.Error(w, "Order Id Not found", http.StatusNotFound)
		return
	}
	var order bson.M
	err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
	if err != nil {
		http.Error(w, "Error in Finding order", http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(order)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {

}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {

}
