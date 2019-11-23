package mood

import (
	"github.com/google/uuid"
	"time"
)

type Mood struct {
	Id         string `json:"id"`
	EmoticonId string `json:"emoticon_id"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	ImageUrl   string `json:"image_url"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func New(emoticonId, title, desc, imageUrl string) *Mood {
	m := new(Mood)
	m.Id = uuid.New().String()
	m.EmoticonId = emoticonId
	m.Title = title
	m.Desc = desc
	m.ImageUrl = imageUrl
	m.CreateTime = time.Now().String()
	m.UpdateTime = time.Now().String()
	return m
}
