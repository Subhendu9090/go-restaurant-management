package controllers

import (
	"context"
	"encoding/json"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var foodCollection *mongo.Collection = database.GetCollectionName("food")
var validate = validator.New()

func GetFoods(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	params := mux.Vars(r)
	recordPerPage, err := strconv.Atoi(params["recordPerPage"])
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, err := strconv.Atoi(params["page"])
	if err != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordPerPage

	totalCount, err := foodCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error counting foods", http.StatusInternalServerError)
		return
	}

	opts := options.Find().SetSkip(int64(startIndex)).SetLimit(int64(recordPerPage))
	crusor, err := foodCollection.Find(ctx, bson.M{}, opts)
	if err != nil {
		http.Error(w, "Error fetching foods", http.StatusInternalServerError)
		return
	}
	var foods models.Food
	err = crusor.All(ctx, &foods)
	if err != nil {
		http.Error(w, "Error fetching foods", http.StatusInternalServerError)
		return
	}
	// Create response structure
	response := map[string]interface{}{
		"totalItems": totalCount,
		"page":       page,
		"perPage":    recordPerPage,
		"data":       foods,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

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
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var food models.Food

	params := mux.Vars(r)
	foodId := params["food_id"]
	err := json.NewDecoder(r.Body).Decode(&food)

	if err != nil {
		http.Error(w, "Error counting foods", http.StatusInternalServerError)
		return
	}
	// Dynamically create update fields
	updateFields := bson.M{}
	if food.Name != nil && *food.Name != "" {
		updateFields["name"] = food.Name
	}
	if *food.Price > 0 {
		updateFields["price"] = food.Price
	}
	if food.Food_image != nil && *food.Food_image != "" {
		updateFields["food_image"] = food.Food_image
	}

	// If no valid fields were provided, return an error
	if len(updateFields) == 0 {
		http.Error(w, `{"error": "No valid fields provided for update"}`, http.StatusBadRequest)
		return
	}

	// Always update the timestamp
	updateFields["updated_at"] = primitive.NewDateTimeFromTime(time.Now())

	update := bson.M{"$set": updateFields}

	// update := bson.M{
	// 	"$set": bson.M{
	// 		"name":       food.Name,
	// 		"price":      food.Price,
	// 		"food_image": food.Food_image,
	// 		"updated_at": time.Now(),
	// 	},
	// }
	var updatedFood models.Food
	err = foodCollection.FindOneAndUpdate(ctx, bson.M{"food_id": foodId}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedFood)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Failed to update menu"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedFood)
}
