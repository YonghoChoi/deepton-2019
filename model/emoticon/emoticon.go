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

func New(name, title, url string) *Emoticon {
	m := new(Emoticon)
	m.Id = uuid.New().String()
	m.Name = name
	m.Title = title
	m.Url = url
	m.CreateTime = time.Now().String()
	m.UpdateTime = time.Now().String()
	return m
}
