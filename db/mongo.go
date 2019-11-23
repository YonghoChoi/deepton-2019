package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"yonghochoi.com/depthon-2019/cmd/pengha-api/conf"
)

const DBName = "pengha"

var gClient *mongo.Client

func GetClient() *mongo.Client {
	if gClient == nil {
		var err error
		clientOptions := options.Client().ApplyURI("mongodb://" + conf.GetConfig().Mongo)
		gClient, err = mongo.NewClient(clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		err = gClient.Connect(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		log.Println("connect success. mongo : " + conf.GetConfig().Mongo)
	}

	return gClient
}

func GetCollection(collection string) *mongo.Collection {
	return GetClient().Database(DBName).Collection(collection)
}
