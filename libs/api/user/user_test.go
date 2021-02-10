package user

import (
	"testing"
)

func TestUser(t *testing.T) {
	result := User("world")
	if result != "User world" {
		t.Error("Expected User to append 'world'")
	}
}
