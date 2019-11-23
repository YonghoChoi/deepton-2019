package mood

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	moodTypes := []string{
		"good", "flex", "sad", "angry", "lonely", "happy", "embarrassed",
	}

	moodTypeCount := len(moodTypes)
	size := 60
	ago := time.Now().AddDate(0, -3, 0)
	for i := 0; i < size; i++ {
		moodChoice := rand.Int() % moodTypeCount
		dt := ago.AddDate(0, 0, i)
		m := New(moodTypes[moodChoice], moodTypes[moodChoice], moodTypes[moodChoice], "")
		m.CreateTime = dt
		m.UpdateTime = dt
		if err := Insert(m); err != nil {
			fmt.Println(err.Error())
		}
	}

	m := New("angry", "angry", "angry", "")
	if err := Insert(m); err != nil {
		t.Fatal(err.Error())
	}
}

func TestUpdate(t *testing.T) {
	m := New("angry", "angry", "angry", "")
	if err := Insert(m); err != nil {
		t.Fatal(err.Error())
	}

	m.EmoticonType = "happy"
	if err := Update(m); err != nil {
		t.Fatal(err.Error())
	}
}

func TestDelete(t *testing.T) {
	u := New("angry", "angry", "angry", "")
	if err := Insert(u); err != nil {
		t.Fatal(err.Error())
	}

	if err := Delete(u); err != nil {
		t.Fatal(err.Error())
	}
}

func TestFindOne(t *testing.T) {
	emoticonType := "angry"
	filter := bson.D{
		{"type", emoticonType},
	}
	u, err := FindOne(filter)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(u)
}

func TestFind(t *testing.T) {
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"emoticon_type", "good"}},
		}},
	}
	users, err := Find(filter)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(users)
}
