package service

import (
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"yonghochoi.com/depthon-2019/model/mood"
)

func CreateMood(m mood.Mood) (mood.Mood, error) {
	m.Id = uuid.New().String()
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()

	if err := mood.Insert(m); err != nil {
		fmt.Println(err.Error())
		return m, fmt.Errorf("database error.\n")
	}

	fmt.Printf("create mood sucess. %v\n", m)
	return m, nil
}

func GetMoods() ([]mood.Mood, error) {
	//now := time.Now()
	//ago := now.AddDate(0, -3, 0)
	//filter := bson.D{
	//	{"create_time", []interface{}{
	//		bson.D{{"$gt", ago}},
	//		bson.D{{"$lt", now}},
	//	}},
	//}
	return mood.Find(bson.D{})
}
