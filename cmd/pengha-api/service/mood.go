package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"yonghochoi.com/depthon-2019/model/mood"
)

func CreateMood(m mood.Mood) (mood.Mood, error) {
	if err := mood.Insert(m); err != nil {
		fmt.Println(err.Error())
		return m, fmt.Errorf("database error.\n")
	}

	fmt.Printf("create mood sucess. %v\n", m)
	return m, nil
}

func GetMoods() ([]mood.Mood, error) {
	return mood.Find(bson.D{})
}
