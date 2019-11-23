package mood

import (
	"github.com/google/uuid"
	"time"
)

type Mood struct {
	Id           string    `json:"id" bson:"_id"`
	EmoticonType string    `json:"emoticon_type"`
	Title        string    `json:"title"`
	Desc         string    `json:"desc"`
	ImageUrl     string    `json:"image_url"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
}

func New(emoticonId, title, desc, imageUrl string) *Mood {
	m := new(Mood)
	m.Id = uuid.New().String()
	m.EmoticonType = emoticonId
	m.Title = title
	m.Desc = desc
	m.ImageUrl = imageUrl
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return m
}
