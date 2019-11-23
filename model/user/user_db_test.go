package user

import (
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
	u, err := FindOne(id)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(u)
}
