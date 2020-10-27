package db

import (
	"log"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


const (
	dbName	= "Todo-App"
	collectionName = "todo-list"
)

var Collection *mongo.Collection


func Init() {
	URI := "mongodb://localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {

		log.Fatal(err)

	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	
	err = client.Connect(ctx)
	if err != nil {

		log.Fatal(err)

	}

	Collection = client.Database(dbName).Collection(collectionName)

}