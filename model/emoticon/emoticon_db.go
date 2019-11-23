package emoticon

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"yonghochoi.com/depthon-2019/db"
)

const CollectionName = "emoticons"

func Insert(u Emoticon) error {
	_, err := db.GetCollection(CollectionName).InsertOne(context.TODO(), u)
	return err
}

func Find(filter bson.D) ([]Emoticon, error) {
	cur, err := db.GetCollection(CollectionName).
		Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	var emoticons []Emoticon
	for cur.Next(context.TODO()) {
		var e Emoticon
		err := cur.Decode(&e)
		if err != nil {
			fmt.Println(err)
			continue
		}

		emoticons = append(emoticons, e)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	if err := cur.Close(context.TODO()); err != nil {
		return nil, err
	}
	return emoticons, nil
}
