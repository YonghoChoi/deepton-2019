package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const DBName = "pengha"

var gClient *mongo.Client

func GetClient() *mongo.Client {
	if gClient == nil {
		var err error
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		gClient, err = mongo.NewClient(clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		err = gClient.Connect(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

	return gClient
}

func GetCollection(collection string) *mongo.Collection {
	return GetClient().Database(DBName).Collection(collection)
}
