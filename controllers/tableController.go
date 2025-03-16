package controllers

import (
	"golang-restaurant-management/database"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

var tableCollection *mongo.Collection = database.GetCollectionName("table")

func GetTables(w http.ResponseWriter, r *http.Request) {

}
func CreateTable(w http.ResponseWriter, r *http.Request) {

}
func UpdateTable(w http.ResponseWriter, r *http.Request) {

}
func GetTable(w http.ResponseWriter, r *http.Request) {

}
