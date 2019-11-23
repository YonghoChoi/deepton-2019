package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Token      string `json:"token"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func (o *User) Join() {
	o.Id = uuid.New().String()
	o.CreateTime = time.Now().String()
	o.UpdateTime = time.Now().String()
}

func New(name, token string) *User {
	u := new(User)
	u.Id = uuid.New().String()
	u.Name = name
	u.Token = token
	u.CreateTime = time.Now().String()
	u.UpdateTime = time.Now().String()
	return u
}
