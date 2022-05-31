package mcore

import (
	"testing"
)

func TestMcore(t *testing.T) {
	result := Mcore("works")
	if result != "Mcore works" {
		t.Error("Expected Mcore to append 'works'")
	}
}
