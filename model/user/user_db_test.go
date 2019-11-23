package user

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestInsert(t *testing.T) {
	u := New("test user", "test token")
	if err := Insert(u); err != nil {
		t.Fatal(err.Error())
	}
}

func TestUpdate(t *testing.T) {
	u := New("test user", "test token")
	if err := Insert(u); err != nil {
		t.Fatal(err.Error())
	}

	u.Name = "modify user"
	if err := Update(u); err != nil {
		t.Fatal(err.Error())
	}
}

func TestDelete(t *testing.T) {
	u := New("test user", "test token")
	if err := Insert(u); err != nil {
		t.Fatal(err.Error())
	}

	if err := Delete(u); err != nil {
		t.Fatal(err.Error())
	}
}

func TestFindOne(t *testing.T) {
	id := "d81180ec-cf4a-4542-b807-e1d4532f2f67"
	filter := bson.D{
		{"_id", id},
	}
	u, err := FindOne(filter)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(u)
}

func TestFind(t *testing.T) {
	id := "86ac7a4e-7544-46a1-88a7-703c84dfd00f"
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"_id", id}},
			bson.D{{"name", "test user"}},
		}},
	}
	users, err := Find(filter)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(users)
}
