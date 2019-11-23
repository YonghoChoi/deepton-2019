package mood

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Mood struct {
	Id            string    `json:"id" bson:"_id"`
	EmoticonType  string    `json:"emoticon_type" bson:"emoticon_type"`
	Title         string    `json:"title" bson:"title"`
	Desc          string    `json:"desc" bson:"desc"`
	ImageUrl      string    `json:"image_url" bson:"image_url"`
	CreateTime    time.Time `json:"create_time" bson:"create_time"`
	UpdateTime    time.Time `json:"update_time" bson:"update_time"`
	UpdateTimeStr string    `json:"update_time_str" bson:"update_time_str"`
}

func (m *Mood) SetUpdateTimeStr() {
	m.UpdateTimeStr = fmt.Sprintf("%d년 %d월 %d일", m.UpdateTime.Year(), m.UpdateTime.Month(), m.UpdateTime.Day())
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
	m.SetUpdateTimeStr()
	return m
}
