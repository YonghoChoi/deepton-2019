package main

import (
	"fmt"
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"time"
	"yonghochoi.com/depthon-2019/cmd/pengha-api/service"
	"yonghochoi.com/depthon-2019/model/emoticon"
	"yonghochoi.com/depthon-2019/model/mood"
	"yonghochoi.com/depthon-2019/model/packet"
	"yonghochoi.com/depthon-2019/model/user"
)

func InitEmoticons() {
	emoticons, _ := service.GetEmoticons()
	if len(emoticons) > 0 {
		return
	}

	emoticons = []emoticon.Emoticon{
		emoticon.New("best", "기분 최고!", "http://pengha.yonghochoi.com/icons/icon_01.png"),
		emoticon.New("angry", "화나!", "http://pengha.yonghochoi.com/icons/icon_02.png"),
		emoticon.New("sad", "슬퍼", "http://pengha.yonghochoi.com/icons/icon_03.png"),
		emoticon.New("embarrassed", "당황스러워", "http://pengha.yonghochoi.com/icons/icon_04.png"),
		emoticon.New("happy", "행복해", "http://pengha.yonghochoi.com/icons/icon_05.png"),
		emoticon.New("lonely", "외로워..", "http://pengha.yonghochoi.com/icons/icon_06.png"),
		emoticon.New("flex", "Flex", "http://pengha.yonghochoi.com/icons/icon_07.png"),
		emoticon.New("good", "좋아!", "http://pengha.yonghochoi.com/icons/icon_08.png"),
	}

	for _, e := range emoticons {
		if err := emoticon.Insert(e); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func InitMoods() {
	moods, _ := service.GetMoods()
	if len(moods) > 0 {
		return
	}
	moodTypes := []string{
		"good", "flex", "sad", "angry", "lonely", "happy", "embarrassed", "best",
	}
	titles := []string{
		"좋아!", "Flex!", "슬퍼", "화나!", "외로워..", "행복해", "당황스러워", "기분 최고!",
	}
	descs := []string{
		"디프만 좋아!", "플렉스!", "시간이 짧아서 너무 슬퍼", "다 못해서 화나!", "혼자라 외로워..", "겁나 행복해", "졸 당황스러워", "끝나서 기분 최고!",
	}

	moodTypeCount := len(moodTypes)
	size := 90
	ago := time.Now().AddDate(0, -3, 0)
	for i := 0; i < size; i++ {
		moodChoice := rand.Int() % moodTypeCount
		dt := ago.AddDate(0, 0, i)
		m := mood.New(moodTypes[moodChoice], titles[moodChoice], descs[moodChoice], "")
		m.CreateTime = dt
		m.UpdateTime = dt
		m.SetUpdateTimeStr()
		if err := mood.Insert(m); err != nil {
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
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	m, err := service.GetMoods()
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = m
	return nil
}

func CreateMood(c echo.Context) error {
	resp := packet.Resp{Code: "200"}
	defer func() {
		if err := c.JSON(http.StatusOK, resp); err != nil {
			fmt.Println(err.Error())
		}
	}()

	var m mood.Mood
	if err := c.Bind(&m); err != nil {
		resp.Code = "500"
		resp.Message = "invalid data"
		fmt.Println(err.Error())
		return nil
	}

	m, err := service.CreateMood(m)
	if err != nil {
		resp.Code = "500"
		resp.Message = err.Error()
		fmt.Println(err.Error())
		return nil
	}

	resp.Data = m
	return nil
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
