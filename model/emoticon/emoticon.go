package emoticon

import (
	"github.com/google/uuid"
	"time"
)

type Emoticon struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	Url        string `json:"url"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func New(name, title, url string) Emoticon {
	e := Emoticon{}
	e.Id = uuid.New().String()
	e.Name = name
	e.Title = title
	e.Url = url
	e.CreateTime = time.Now().String()
	e.UpdateTime = time.Now().String()
	return e
}
