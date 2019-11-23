package emoticon

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestInsert(t *testing.T) {
	emoticons := []Emoticon{
		New("icon_01", "하트", "http://pengha.yonghochoi.com/icons/icon_01.png"),
		New("icon_02", "흠", "http://pengha.yonghochoi.com/icons/icon_02.png"),
		New("icon_03", "눈물", "http://pengha.yonghochoi.com/icons/icon_03.png"),
		New("icon_04", "땀", "http://pengha.yonghochoi.com/icons/icon_04.png"),
		New("icon_05", "부끄", "http://pengha.yonghochoi.com/icons/icon_05.png"),
		New("icon_06", "띠용", "http://pengha.yonghochoi.com/icons/icon_06.png"),
		New("icon_07", "멋짐", "http://pengha.yonghochoi.com/icons/icon_07.png"),
		New("icon_08", "야호", "http://pengha.yonghochoi.com/icons/icon_08.png"),
	}

	for _, emoticon := range emoticons {
		if err := Insert(emoticon); err != nil {
			t.Error(err.Error())
		}
	}
}

func TestFind(t *testing.T) {
	filter := bson.D{}
	emos, err := Find(filter)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(emos)
}
