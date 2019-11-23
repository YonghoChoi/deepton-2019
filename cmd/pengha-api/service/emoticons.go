package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"yonghochoi.com/depthon-2019/model/emoticon"
)

func GetEmoticons() ([]emoticon.Emoticon, error) {
	return emoticon.Find(bson.D{})
}
