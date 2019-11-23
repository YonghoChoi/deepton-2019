package mood

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"yonghochoi.com/depthon-2019/db"
)

const CollectionName = "moods"

func Insert(m Mood) error {
	_, err := db.GetCollection(CollectionName).InsertOne(context.TODO(), m)
	return err
}

func Delete(m Mood) error {
	_, err := db.GetCollection(CollectionName).
		DeleteOne(context.TODO(), bson.M{"_id": m.Id})
	return err
}

func Update(m Mood) error {
	_, err := db.GetCollection(CollectionName).
		UpdateOne(
			context.TODO(),
			bson.M{"_id": m.Id},
			bson.D{{Key: "$set", Value: m}},
		)
	return err
}

func FindOne(filter bson.D) (Mood, error) {
	var m Mood
	result := db.GetCollection(CollectionName).
		FindOne(context.TODO(), filter)

	if result.Err() != nil {
		fmt.Println(result.Err().Error())
		if result.Err() == mongo.ErrNoDocuments || result.Err() == mongo.ErrNilDocument {
			return m, fmt.Errorf("not exist mood")
		}
	}

	if err := result.Decode(&m); err != nil {
		fmt.Println(err.Error())
		return m, fmt.Errorf("invalid user data type")
	}

	return m, nil
}

func Find(filter bson.D) ([]Mood, error) {
	cur, err := db.GetCollection(CollectionName).
		Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	var moods []Mood
	for cur.Next(context.TODO()) {
		var m Mood
		err := cur.Decode(&m)
		if err != nil {
			fmt.Println(err)
			continue
		}

		moods = append(moods, m)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	if err := cur.Close(context.TODO()); err != nil {
		return nil, err
	}
	return moods, nil
}
