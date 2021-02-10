package feature

import (
	"testing"
)

func TestPostFeature(t *testing.T) {
	result := PostFeature("world")
	if result != "PostFeature world" {
		t.Error("Expected PostFeature to append 'world'")
	}
}
