package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"yonghochoi.com/depthon-2019/db/mongo"
)

const CollectionName = "users"

func Insert(u User) error {
	_, err := mongo.GetCollection(CollectionName).InsertOne(context.TODO(), u)
	return err
}

func Delete(u User) error {
	_, err := mongo.GetCollection(CollectionName).
		DeleteOne(context.TODO(), bson.M{"_id": u.Id})
	return err
}

func Update(u User) error {
	_, err := mongo.GetCollection(CollectionName).
		UpdateOne(
			context.TODO(),
			bson.M{"_id": u.Id},
			bson.D{{Key: "$set", Value: u}},
		)
	return err
}

func FindOne(id string) (User, error) {
	result := mongo.GetCollection(CollectionName).
		FindOne(context.TODO(), bson.M{"_id": id})

	var u User
	if err := result.Decode(&u); err != nil {
		return u, err
	}

	return u, nil
}
