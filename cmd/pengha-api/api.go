package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"yonghochoi.com/depthon-2019/cmd/pengha-api/service"
	"yonghochoi.com/depthon-2019/model/emoticon"
	"yonghochoi.com/depthon-2019/model/packet"
	"yonghochoi.com/depthon-2019/model/user"
)

func InitSample() {
	emoticons, _ := service.GetEmoticons()
	if len(emoticons) > 0 {
		return
	}

	emoticons := []emoticon.Emoticon{
		emoticon.New("icon_01", "하트", "http://pengha.yonghochoi.com/icons/icon_01.png"),
		emoticon.New("icon_02", "흠", "http://pengha.yonghochoi.com/icons/icon_02.png"),
		emoticon.New("icon_03", "눈물", "http://pengha.yonghochoi.com/icons/icon_03.png"),
		emoticon.New("icon_04", "땀", "http://pengha.yonghochoi.com/icons/icon_04.png"),
		emoticon.New("icon_05", "부끄", "http://pengha.yonghochoi.com/icons/icon_05.png"),
		emoticon.New("icon_06", "띠용", "http://pengha.yonghochoi.com/icons/icon_06.png"),
		emoticon.New("icon_07", "멋짐", "http://pengha.yonghochoi.com/icons/icon_07.png"),
		emoticon.New("icon_08", "야호", "http://pengha.yonghochoi.com/icons/icon_08.png"),
	}

	for _, e := range emoticons {
		if err := emoticon.Insert(e); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func GetEmoticonDatas(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	emoticons, err := service.GetEmoticons()
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = emoticons
	return nil
}

func GetMoods(c echo.Context) error {
	return c.String(http.StatusOK, version)
}

func GetUsers(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	users, err := service.GetUsers()
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())

	}
	resp.Data = users
	return nil
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
