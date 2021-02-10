package post_data_access

import (
	"testing"
)

func TestPostDataAccess(t *testing.T) {
	result := PostDataAccess("world")
	if result != "PostDataAccess world" {
		t.Error("Expected PostDataAccess to append 'world'")
	}
}
