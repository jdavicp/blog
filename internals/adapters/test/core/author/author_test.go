package author_test

import (
	"testing"

	"github.com/jdavicp/blog/internals/adapters/core/author"
	"github.com/jdavicp/blog/internals/adapters/core/author/email"
	"github.com/jdavicp/blog/internals/adapters/core/author/username"
)

func TestCreateValidAuthor(t *testing.T) {
	e := email.New("jdavicp.fac@gmail.com")
	u := username.New("jdavicp")
	user := author.New("José Davi", e, u, "sorvete123")
	want := "José Davi"
	got := user.Name()

	if got != want {
		t.Errorf("Want: %s, got: %s", want, got)
	}
}

func TestCantCreateAuthorWithInvalidEmail(t *testing.T) {
	e := email.New("")
	u := username.New("jdavicp")
	user := author.New("José Davi", e, u, "sorvete123")
	want := ""
	got := user.Name()

	if got != want {
		t.Errorf("Want: %s, got: %s", want, got)
	}
}

func TestCantCreateAuthorWithInvalidUsername(t *testing.T) {
	e := email.New("jdavicp.fac@gmail.com")
	u := username.New("jd")
	user := author.New("José Davi", e, u, "sorvete123")
	want := ""
	got := user.Name()

	if got != want {
		t.Errorf("Want: %s, got: %s", want, got)
	}
}
