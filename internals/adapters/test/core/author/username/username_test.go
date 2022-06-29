package username_test

import (
	"testing"

	"github.com/jdavicp/blog/internals/adapters/core/author/username"
)

func TestValidUsername(t *testing.T) {
	username := username.New("jdavicp")
	want := "jdavicp"
	got := username.Username()

	if got != want {
		t.Errorf("Want: %s, got: %s", want, got)
	}
}

func TestInvalidUsername(t *testing.T) {
	username := username.New("jdavicp45")
	want := ""
	got := username.Username()

	if got != want {
		t.Errorf("Want: %s, got: %s", want, got)
	}
}
