package controllers

import (
	"context"
	"encoding/json"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var table models.Table
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = validate.Struct(order)
	if err != nil {
		http.Error(w, "Give the required Field", http.StatusNotFound)
		return
	}
	err = tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
	if err != nil {
		http.Error(w, "Table not found", http.StatusNotFound)
		return
	}
	order.Id = primitive.NewObjectID()
	order.Created_at = time.Now()
	order.Updated_at = time.Now()
	order.Order_id = order.Id.Hex()

	result, err := orderCollection.InsertOne(ctx, order)
	if err != nil {
		http.Error(w, "error in creating order", 400)
		return
	}
	defer cancel()
	w.Header().Set("Content/type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	params := mux.Vars(r)
	orderId := params["order_id"]
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Error"+err.Error(), http.StatusInternalServerError)
		return
	}
	updatedFields := bson.M{}

	if order.Table_id != nil {
		updatedFields["table_id"] = order.Table_id
	}
	var updatedOrder = bson.M{}
	updatedFields["updated_at"] = time.Now()
	update := bson.M{"$set": updatedFields}
	err = orderCollection.FindOneAndUpdate(ctx, bson.M{"order_id": orderId}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedOrder)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Failed to update menu"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedOrder)

}
