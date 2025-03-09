package controllers

import (
	"context"
	"encoding/json"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.GetCollectionName("food")
var validate = validator.New()

func GetFoods(w http.ResponseWriter, r *http.Request) {

}
func GetFood(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	params := mux.Vars(r)
	foodId := params["food_id"]
	var food models.Food

	err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
	defer cancel()

	if err != nil {
		http.Error(w, "Error in finding food ", http.StatusNotFound)
		return
	}
	w.Header().Set("Content/type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(food)
}

func CreateFood(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var menu models.Menu
	var food models.Food
	err := json.NewDecoder(r.Body).Decode(&food)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = validate.Struct(food)
	if err != nil {
		http.Error(w, " validation Error", http.StatusForbidden)
		return
	}
	err = menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
	defer cancel()
	if err != nil {
		http.Error(w, "Menu not Found", http.StatusNotFound)
		return
	}
	food.Created_at = time.Now()
	food.Updated_at = time.Now()

	food.Id = primitive.NewObjectID()
	food.Food_id = food.Id.Hex()
	// var num = toFixed(*food.Price, 2)
	// food.Price = &num

	result, err := foodCollection.InsertOne(ctx, food)
	if err != nil {
		http.Error(w, "error in creating food", 400)
	}
	defer cancel()
	w.Header().Set("Content/type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func UpdateFoods(w http.ResponseWriter, r *http.Request) {

}
