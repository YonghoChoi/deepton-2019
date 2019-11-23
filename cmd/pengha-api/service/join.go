package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"yonghochoi.com/depthon-2019/model/user"
)

func Join(u user.User) (user.User, error) {
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"_id", u.Id}},
			bson.D{{"name", u.Name}},
		}},
	}

	users, err := user.Find(filter)
	if err != nil {
		fmt.Println(err.Error())
		return u, fmt.Errorf("database error.\n")
	}

	if len(users) > 0 {
		return u, fmt.Errorf("already exist user")
	}

	u.Join()
	if err := user.Insert(u); err != nil {
		fmt.Println(err.Error())
		return u, fmt.Errorf("database error.\n")
	}

	fmt.Printf("join sucess. %v\n", u)
	return u, nil
}

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
