package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"yonghochoi.com/depthon-2019/model/user"
)

func GetUsers() ([]user.User, error) {
	return user.Find(bson.D{})
}
