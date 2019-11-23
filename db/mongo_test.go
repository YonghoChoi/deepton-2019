package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"testing"
)

func TestGetClient(t *testing.T) {
	c := GetClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
}
