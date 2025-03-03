package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDb := "mongodb://localhost:27107"
	fmt.Println(MongoDb)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal("Error in Mongo db connection ", err)
	}
	fmt.Println("Connected To MongoDb")
	return mongoClient
}

func GetCollectionName(collectionName string) *mongo.Collection {
	dbName := os.Getenv("DB_NAME")

	if dbName == "" {
		log.Fatal("Db name missing from env")
	}
	Client := DBinstance()
	return Client.Database(dbName).Collection(collectionName)
}
