package emoticon

import (
	"github.com/google/uuid"
	"time"
)

type Emoticon struct {
	Id         string    `json:"id" bson:"_id"`
	Type       string    `json:"type"`
	Title      string    `json:"title"`
	Url        string    `json:"url"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func New(name, title, url string) Emoticon {
	e := Emoticon{}
	e.Id = uuid.New().String()
	e.Type = name
	e.Title = title
	e.Url = url

	e.CreateTime = time.Now()
	e.UpdateTime = time.Now()
	return e
}
