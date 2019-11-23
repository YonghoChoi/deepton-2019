package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"yonghochoi.com/depthon-2019/model/user"
)

func Login(reqUser user.User) (user.User, error) {
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"_id", reqUser.Id}},
			bson.D{{"name", reqUser.Name}},
		}},
	}
	findUser, err := user.FindOne(filter)
	if err != nil {
		return findUser, err
	}

	if reqUser.Token != findUser.Token {
		return reqUser, fmt.Errorf("invalid token")
	}

	return findUser, nil
}
