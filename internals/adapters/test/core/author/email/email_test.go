package email_test

import (
	"testing"

	e "github.com/jdavicp/blog/internals/adapters/core/author/email"
)

func TestValidEmail(t *testing.T) {
	email := e.New("jdavi.fac@gmail.com")
	want := "jdavi.fac@gmail.com"
	got := email.Address()

	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestInvalidEmail(t *testing.T) {
	email := e.New("jdavi")
	want := ""
	got := email.Address()
	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}
}
