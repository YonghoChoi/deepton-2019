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
