package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func ConnectToMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/"))

	defer cancel()

	if err != nil {
		log.Fatalf("connection error :%v", err)
		return nil
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("ping mongodb error :%v", err)
		return nil
	}

	return mongoClient.Database("pass-gw")
}

func GetServiceCollection(db *mongo.Database) *mongo.Collection {
	coll := db.Collection("service")
	return coll
}
