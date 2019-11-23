package mood

import (
	"github.com/google/uuid"
	"time"
)

type Mood struct {
	Id           string    `json:"id" bson:"_id"`
	EmoticonType string    `json:"emoticon_type" bson:"emoticon_type"`
	Title        string    `json:"title" bson:"title"`
	Desc         string    `json:"desc" bson:"desc"`
	ImageUrl     string    `json:"image_url" bson:"image_url"`
	CreateTime   time.Time `json:"create_time" bson:"create_time"`
	UpdateTime   time.Time `json:"update_time" bson:"update_time"`
}

func New(emoticonType, title, desc, imageUrl string) Mood {
	m := Mood{}
	m.Id = uuid.New().String()
	m.EmoticonType = emoticonType
	m.Title = title
	m.Desc = desc
	m.ImageUrl = imageUrl
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return m
}
