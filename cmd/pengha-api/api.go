package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
	"yonghochoi.com/depthon-2019/cmd/pengha-api/service"
	"yonghochoi.com/depthon-2019/model/emoticon"
	"yonghochoi.com/depthon-2019/model/packet"
	"yonghochoi.com/depthon-2019/model/user"
)

var emoticonDatas []emoticon.Emoticon
var userDatas []user.User

func InitSample() {
	emoticonDatas = []emoticon.Emoticon{
		emoticon.New("good", "기분 좋음", "http://localhost"),
		emoticon.New("angry", "화남", "http://localhost"),
		emoticon.New("boring", "따분함", "http://localhost"),
		emoticon.New("sad", "슬픔", "http://localhost"),
		emoticon.New("tired", "지침", "http://localhost"),
	}

	for _, name := range []string{"펭수1", "펭수2", "펭수3", "펭수4"} {
		userDatas = append(userDatas, user.New(name, uuid.New().String()))
	}
}

func GetEmoticonDatas(c echo.Context) error {
	return c.JSON(http.StatusOK, emoticonDatas)
}

func GetMoods(c echo.Context) error {
	return c.String(http.StatusOK, version)
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, userDatas)
}

func Join(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	var u user.User
	if err := c.Bind(&u); err != nil {
		resp.Code = "500"
		resp.Message = "invalid data"
		fmt.Println(err.Error())
		return nil
	}

	u, err := service.Join(u)
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = u
	return nil
}

func Login(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	var u user.User
	if err := c.Bind(&u); err != nil {
		resp.Code = "500"
		resp.Message = "invalid data"
		fmt.Println(err.Error())
		return nil
	}

	u, err := service.Login(u)
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = u
	return nil
}
