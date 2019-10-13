package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client : Database variable
var Client *mongo.Client

// DbConnect : Connecting MongoDB
func DbConnect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	Client, _ = mongo.Connect(ctx, clientOptions)
	if Client != nil {
		log.Println("DB Connected")
	}
}
