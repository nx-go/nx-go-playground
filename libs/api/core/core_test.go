package core

import (
	"testing"
)

func TestCore(t *testing.T) {
	result := Core("world")
	if result != "Core world" {
		t.Error("Expected Core to append 'world'")
	}
}
