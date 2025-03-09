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

var menuCollection *mongo.Collection = database.GetCollectionName("Menu")

func GetMenus(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	result, err := menuCollection.Find(ctx, bson.M{})

	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	var allMenu []bson.M
	err = result.All(ctx, &allMenu)

	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	w.Header().Set("Content/type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allMenu)
}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	params := mux.Vars(r)
	menuId := params["menu_id"]
	var menu models.Menu
	err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menu)
}

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var menu models.Menu
	err := json.NewDecoder(r.Body).Decode(&menu)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	validateErr := validate.Struct(menu)
	if validateErr != nil {
		http.Error(w, "Validation Error", http.StatusBadRequest)
		return
	}
	menu.Created_at = time.Now()
	menu.Updated_at = time.Now()
	menu.Id = primitive.NewObjectID()
	menu.Menu_id = menu.Id.Hex()
	result, inserterr := menuCollection.InsertOne(ctx, &menu)
	if inserterr != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Inserting Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var menu models.Menu
	if err := json.NewDecoder(r.Body).Decode(&menu); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	menuId := params["menu_id"]

	update := bson.M{
		"$set": bson.M{
			"name":       menu.Name,
			"category":   menu.Category,
			"start_date": menu.Start_date,
			"end_date":   menu.End_date,
			"updated_at": time.Now(),
		},
	}

	var updatedMenu models.Menu
	err := menuCollection.FindOneAndUpdate(
		ctx,
		bson.M{"menu_id": menuId},
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedMenu)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		if err == mongo.ErrNoDocuments {
			http.Error(w, `{"error": "Menu not found"}`, http.StatusNotFound)
		} else {
			http.Error(w, `{"error": "Failed to update menu"}`, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedMenu)
}
